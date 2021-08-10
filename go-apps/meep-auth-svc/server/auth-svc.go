/*
 * Copyright (c) 2020  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the \"License\");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an \"AS IS\" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * AdvantEDGE Platform Controller REST API
 *
 * This API is the main Platform Controller API for scenario configuration & sandbox management <p>**Micro-service**<br>[meep-pfm-ctrl](https://github.com/InterDigitalInc/AdvantEDGE/tree/master/go-apps/meep-platform-ctrl) <p>**Type & Usage**<br>Platform main interface used by controller software to configure scenarios and manage sandboxes in the AdvantEDGE platform <p>**Details**<br>API details available at _your-AdvantEDGE-ip-address/api_
 *
 * API version: 1.0.0
 * Contact: AdvantEDGE@InterDigital.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	dataModel "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model"
	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
	met "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-metrics"
	mq "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mq"
	pcc "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-platform-ctrl-client"
	sm "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-sessions"
	sam "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-swagger-api-mgr"
	users "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-users"
	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/roymx/viper"
	"github.com/xanzy/go-gitlab"
	"golang.org/x/oauth2"
)

const OAUTH_PROVIDER_GITHUB = "github"
const OAUTH_PROVIDER_GITLAB = "gitlab"
const OAUTH_PROVIDER_LOCAL = "local"

const serviceName = "Auth Service"
const moduleName = "meep-auth-svc"
const moduleNamespace = "default"
const postgisUser = "postgres"
const postgisPwd = "pwd"
const pfmCtrlBasepath = "http://meep-platform-ctrl/platform-ctrl/v1"
const providerModeSecure = "secure"
const mepPrefix = "mep--"

// Permission Configuration types
type Permission struct {
	Mode  string            `yaml:"mode"`
	Roles map[string]string `yaml:"roles"`
}
type Fileserver struct {
	Name  string            `yaml:"name"`
	Path  string            `yaml:"path"`
	Sbox  bool              `yaml:"sbox"`
	Mode  string            `yaml:"mode"`
	Roles map[string]string `yaml:"roles"`
}
type Endpoint struct {
	Name   string            `yaml:"name"`
	Path   string            `yaml:"path"`
	Method string            `yaml:"method"`
	Sbox   bool              `yaml:"sbox"`
	Mode   string            `yaml:"mode"`
	Roles  map[string]string `yaml:"roles"`
}
type Service struct {
	Name        string       `yaml:"name"`
	Api         string       `yaml:"api"`
	Path        string       `yaml:"path"`
	Sbox        bool         `yaml:"sbox"`
	Default     Permission   `yaml:"default"`
	Endpoints   []Endpoint   `yaml:"endpoints"`
	Fileservers []Fileserver `yaml:"fileservers"`
}
type PermissionsConfig struct {
	Default     Permission   `yaml:"default"`
	Fileservers []Fileserver `yaml:"fileservers"`
	Services    []Service    `yaml:"services"`
}

// Auth Service types
type AuthRoute struct {
	Name    string
	Method  string
	Pattern string
	Prefix  bool
}

type LoginRequest struct {
	provider      string
	createSandbox string
	timer         *time.Timer
}

type PermissionsCache struct {
	Default     *Permission
	Fileservers map[string]*Permission
	Services    map[string]map[string]*Permission
}

type AuthSvc struct {
	sessionMgr    *sm.SessionMgr
	userStore     *users.Connector
	metricStore   *met.MetricStore
	mqGlobal      *mq.MsgQueue
	apiMgr        *sam.SwaggerApiMgr
	pfmCtrlClient *pcc.APIClient
	maxSessions   int
	uri           string
	oauthConfigs  map[string]*oauth2.Config
	loginRequests map[string]*LoginRequest
	router        *mux.Router
	cache         PermissionsCache
}

var mutex sync.Mutex
var gitlabApiUrl = ""

// Declare as variables to enable overwrite in test
var redisDBAddr = "meep-redis-master:6379"
var influxDBAddr string = "http://meep-influxdb.default.svc.cluster.local:8086"

// Auth Service
var authSvc *AuthSvc

// Metrics
var (
	metricSessionLogin = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "auth_svc_session_login_total",
		Help: "The total number of session login attempts",
	}, []string{"type"})
	metricSessionLogout = promauto.NewCounter(prometheus.CounterOpts{
		Name: "auth_svc_session_logout_total",
		Help: "The total number of session logout attempts",
	})
	metricSessionSuccess = promauto.NewCounter(prometheus.CounterOpts{
		Name: "auth_svc_session_success_total",
		Help: "The total number of successful sessions",
	})
	metricSessionFail = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "auth_svc_session_fail_total",
		Help: "The total number of failed session login attempts",
	}, []string{"type"})
	metricSessionTimeout = promauto.NewCounter(prometheus.CounterOpts{
		Name: "auth_svc_session_timeout_total",
		Help: "The total number of timed out sessions",
	})
	metricSessionActive = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "auth_svc_session_active",
		Help: "The number of active sessions",
	})
	metricSessionDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name:    "auth_svc_session_duration",
		Help:    "A histogram of session durations",
		Buckets: prometheus.LinearBuckets(20, 20, 6),
	})
)

func Init() (err error) {

	// Create new Platform Controller
	authSvc = new(AuthSvc)

	// Create message queue
	authSvc.mqGlobal, err = mq.NewMsgQueue(mq.GetGlobalName(), moduleName, moduleNamespace, redisDBAddr)
	if err != nil {
		log.Error("Failed to create Message Queue with error: ", err)
		return err
	}
	log.Info("Message Queue created")

	// Create Swagger API Manager
	authSvc.apiMgr, err = sam.NewSwaggerApiMgr(moduleName, "", "", authSvc.mqGlobal)
	if err != nil {
		log.Error("Failed to create Swagger API Manager. Error: ", err)
		return err
	}
	log.Info("Swagger API Manager created")

	// Create Platform Controller REST API client
	pfmCtrlClientCfg := pcc.NewConfiguration()
	pfmCtrlClientCfg.BasePath = pfmCtrlBasepath
	authSvc.pfmCtrlClient = pcc.NewAPIClient(pfmCtrlClientCfg)
	if authSvc.pfmCtrlClient == nil {
		err := errors.New("Failed to create Platform Ctrl REST API client")
		return err
	}
	log.Info("Platform Ctrl REST API client created")

	// Connect to Session Manager
	authSvc.sessionMgr, err = sm.NewSessionMgr(moduleName, "", redisDBAddr, redisDBAddr)
	if err != nil {
		log.Error("Failed connection to Session Manager: ", err.Error())
		return err
	}
	log.Info("Connected to Session Manager")

	// Connect to User Store
	authSvc.userStore, err = users.NewConnector(moduleName, postgisUser, postgisPwd, "", "")
	if err != nil {
		log.Error("Failed connection to User Store: ", err.Error())
		return err
	}
	_ = authSvc.userStore.CreateTables()
	log.Info("Connected to User Store")

	// Retrieve & cache endpoint authorization permissions
	cachePermissions()

	// Connect to Metric Store
	authSvc.metricStore, err = met.NewMetricStore("session-metrics", "global", influxDBAddr, met.MetricsDbDisabled)
	if err != nil {
		log.Error("Failed connection to Metric Store: ", err)
		return err
	}

	// Retrieve maximum session count from environment variable
	if maxSessions, err := strconv.ParseInt(os.Getenv("MEEP_MAX_SESSIONS"), 10, 0); err == nil {
		authSvc.maxSessions = int(maxSessions)
	}
	log.Info("MEEP_MAX_SESSIONS: ", authSvc.maxSessions)

	// Get default platform URI
	authSvc.uri = strings.TrimSpace(os.Getenv("MEEP_HOST_URL"))

	// Initialize OAuth
	authSvc.oauthConfigs = make(map[string]*oauth2.Config)
	authSvc.loginRequests = make(map[string]*LoginRequest)

	// Initialize Github config
	githubEnabledStr := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITHUB_ENABLED"))
	githubEnabled, err := strconv.ParseBool(githubEnabledStr)
	if err == nil && githubEnabled {
		clientId := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITHUB_CLIENT_ID"))
		secret := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITHUB_SECRET"))
		redirectUri := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITHUB_REDIRECT_URI"))
		authUrl := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITHUB_AUTH_URL"))
		tokenUrl := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITHUB_TOKEN_URL"))
		if clientId != "" && secret != "" && redirectUri != "" && authUrl != "" && tokenUrl != "" {
			oauthConfig := &oauth2.Config{
				ClientID:     clientId,
				ClientSecret: secret,
				RedirectURL:  redirectUri,
				Scopes:       []string{},
				Endpoint: oauth2.Endpoint{
					AuthURL:  authUrl,
					TokenURL: tokenUrl,
				},
			}
			authSvc.oauthConfigs[OAUTH_PROVIDER_GITHUB] = oauthConfig
			log.Info("GitHub OAuth provider enabled")
		}
	}

	// Initialize GitLab config
	gitlabEnabledStr := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITLAB_ENABLED"))
	gitlabEnabled, err := strconv.ParseBool(gitlabEnabledStr)
	if err == nil && gitlabEnabled {
		gitlabApiUrl = strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITLAB_API_URL"))
		clientId := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITLAB_CLIENT_ID"))
		secret := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITLAB_SECRET"))
		redirectUri := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITLAB_REDIRECT_URI"))
		authUrl := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITLAB_AUTH_URL"))
		tokenUrl := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITLAB_TOKEN_URL"))
		if clientId != "" && secret != "" && redirectUri != "" && authUrl != "" && tokenUrl != "" {
			oauthConfig := &oauth2.Config{
				ClientID:     clientId,
				ClientSecret: secret,
				RedirectURL:  redirectUri,
				Scopes:       []string{"read_user"},
				Endpoint: oauth2.Endpoint{
					AuthURL:  authUrl,
					TokenURL: tokenUrl,
				},
			}
			authSvc.oauthConfigs[OAUTH_PROVIDER_GITLAB] = oauthConfig
			log.Info("GitLab OAuth provider enabled")
		}
	}

	return nil
}

func Run() (err error) {
	// Start Swagger API Manager (provider)
	err = authSvc.apiMgr.Start(true, false)
	if err != nil {
		log.Error("Failed to start Swagger API Manager with error: ", err.Error())
		return err
	}
	log.Info("Swagger API Manager started")

	// Add module Swagger APIs
	err = authSvc.apiMgr.AddApis()
	if err != nil {
		log.Error("Failed to add Swagger APIs with error: ", err.Error())
		return err
	}
	log.Info("Swagger APIs successfully added")

	// Start Session Watchdog
	err = authSvc.sessionMgr.StartSessionWatchdog(sessionTimeoutCb)
	if err != nil {
		log.Error("Failed start Session Watchdog: ", err.Error())
		return err
	}
	return nil
}

func getPermissionsConfig() (config *PermissionsConfig, err error) {
	// Read & apply API permissions from file
	permissionsFile := "/permissions.yaml"
	permissions := viper.New()
	permissions.SetConfigFile(permissionsFile)
	err = permissions.ReadInConfig()
	if err != nil {
		return nil, err
	}

	// Unmarshal config into Permission Configuration structure
	config = new(PermissionsConfig)
	err = permissions.Unmarshal(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func cachePermissions() {
	// Create new Auth Router
	authSvc.router = mux.NewRouter().StrictSlash(true)

	// Get permissions from configuration file
	config, err := getPermissionsConfig()
	if err != nil || config == nil {
		log.Warn("Failed to retrieve permissions from file with err: ", err.Error())
		log.Warn("Granting full API access for all roles by default")

		// Cache default permission
		authSvc.cache.Default = &Permission{Mode: sm.ModeAllow}
		return
	}
	fmt.Printf("%+v\n", config)

	// Parse & cache permissions from config file
	// IMPORTANT NOTE: Order is important to prevent prefix matches from running first
	cacheDefaultPermission(config)
	cacheServicePermissions(config)
	cacheFileserverPermissions(config)
}

func cacheDefaultPermission(cfg *PermissionsConfig) {
	authSvc.cache.Default = &cfg.Default
	if authSvc.cache.Default == nil {
		log.Warn("Failed to retrieve default permission")
		log.Warn("Granting full API access for all roles by default")
		permission := new(Permission)
		permission.Mode = sm.ModeAllow
		authSvc.cache.Default = permission
	}
}

func cacheServicePermissions(cfg *PermissionsConfig) {
	var routes []*AuthRoute

	// Initialize Service permissions cache
	authSvc.cache.Services = make(map[string]map[string]*Permission)

	for _, svc := range cfg.Services {
		// Get/Create service + add it to service cache
		svcMap, found := authSvc.cache.Services[svc.Name]
		if !found {
			svcMap = make(map[string]*Permission)
			authSvc.cache.Services[svc.Name] = svcMap
		}

		// Get API-specific prefix if present
		apiPrefix := ""
		if svc.Api != "" {
			apiPrefix = svc.Api + "--"
		}

		// Service Endpoints
		for _, ep := range svc.Endpoints {
			// Create service endpoint permissions
			permission := new(Permission)
			permission.Mode = ep.Mode
			permission.Roles = make(map[string]string)
			for role, access := range ep.Roles {
				permission.Roles[role] = access
			}

			// Add auth service routes + cache service endpoint permissions
			if svc.Sbox {
				// Mep-specific sandbox service endpoint
				route := new(AuthRoute)
				route.Prefix = false
				route.Method = ep.Method
				route.Name = mepPrefix + apiPrefix + ep.Name
				route.Pattern = "/{sbox}/{mep}" + svc.Path + ep.Path
				routes = append(routes, route)
				svcMap[route.Name] = permission

				// Sandbox service endpoint
				route = new(AuthRoute)
				route.Prefix = false
				route.Method = ep.Method
				route.Name = apiPrefix + ep.Name
				route.Pattern = "/{sbox}" + svc.Path + ep.Path
				routes = append(routes, route)
				svcMap[route.Name] = permission
			} else {
				// Global service endpoint
				route := new(AuthRoute)
				route.Prefix = false
				route.Method = ep.Method
				route.Name = apiPrefix + ep.Name
				route.Pattern = svc.Path + ep.Path
				routes = append(routes, route)
				svcMap[route.Name] = permission
			}
		}

		// Service Fileserver Endpoints
		for _, fs := range svc.Fileservers {
			// Create service fileserver permissions
			permission := new(Permission)
			permission.Mode = fs.Mode
			permission.Roles = make(map[string]string)
			for role, access := range fs.Roles {
				permission.Roles[role] = access
			}

			// Add auth service routes + cache filserver permissions
			if svc.Sbox {
				// Mep-specific sandbox service fileservers
				route := new(AuthRoute)
				route.Prefix = true
				route.Name = mepPrefix + apiPrefix + fs.Name
				route.Pattern = "/{sbox}/{mep}" + svc.Path + fs.Path
				routes = append(routes, route)
				svcMap[route.Name] = permission

				// Sandbox service fileserver
				route = new(AuthRoute)
				route.Prefix = true
				route.Name = apiPrefix + fs.Name
				route.Pattern = "/{sbox}" + svc.Path + fs.Path
				routes = append(routes, route)
				svcMap[route.Name] = permission
			} else {
				// Global service fileserver
				route := new(AuthRoute)
				route.Prefix = true
				route.Name = apiPrefix + fs.Name
				route.Pattern = svc.Path + fs.Path
				routes = append(routes, route)
				svcMap[route.Name] = permission
			}
		}

		// Default service permissions
		// IMPORTANT NOTE: This prefix route must be added after the service endpoint routes
		var permission *Permission
		if svc.Default.Mode != "" {
			permission = new(Permission)
			permission.Roles = make(map[string]string)
			permission.Mode = svc.Default.Mode
			for role, access := range svc.Default.Roles {
				permission.Roles[role] = access
			}
		} else {
			// Use cache default permission if service-specific default is not found
			permission = authSvc.cache.Default
		}

		// Add auth service routes + cache service permissions
		if svc.Sbox {
			// Mep-specific sandbox service
			route := new(AuthRoute)
			route.Prefix = true
			route.Name = mepPrefix + apiPrefix + svc.Name
			route.Pattern = "/{sbox}/{mep}" + svc.Path
			routes = append(routes, route)
			svcMap[route.Name] = permission

			// Sandbox service
			route = new(AuthRoute)
			route.Prefix = true
			route.Name = apiPrefix + svc.Name
			route.Pattern = "/{sbox}" + svc.Path
			routes = append(routes, route)
			svcMap[route.Name] = permission
		} else {
			// Global service
			route := new(AuthRoute)
			route.Prefix = true
			route.Name = apiPrefix + svc.Name
			route.Pattern = svc.Path
			routes = append(routes, route)
			svcMap[route.Name] = permission
		}
	}

	// Add routes to router
	addRoutes(routes)
}

func cacheFileserverPermissions(cfg *PermissionsConfig) {
	var routes []*AuthRoute

	// Initialize Fileserver permissions cache
	authSvc.cache.Fileservers = make(map[string]*Permission)

	for _, fs := range cfg.Fileservers {
		// Create fileserver permissions
		permission := new(Permission)
		permission.Mode = fs.Mode
		permission.Roles = make(map[string]string)
		for role, access := range fs.Roles {
			permission.Roles[role] = access
		}

		// Add auth service routes + cache filserver permissions
		if fs.Sbox {
			// Mep-specific sandbox fileservers
			route := new(AuthRoute)
			route.Prefix = true
			route.Name = mepPrefix + fs.Name
			route.Pattern = "/{sbox}/{mep}" + fs.Path
			routes = append(routes, route)
			authSvc.cache.Fileservers[route.Name] = permission

			// Sandbox fileserver
			route = new(AuthRoute)
			route.Prefix = true
			route.Name = fs.Name
			route.Pattern = "/{sbox}" + fs.Path
			routes = append(routes, route)
			authSvc.cache.Fileservers[route.Name] = permission
		} else {
			// Global fileserver
			route := new(AuthRoute)
			route.Prefix = true
			route.Name = fs.Name
			route.Pattern = fs.Path
			routes = append(routes, route)
			authSvc.cache.Fileservers[route.Name] = permission
		}
	}

	// Add routes to router
	addRoutes(routes)
}

func addRoutes(routes []*AuthRoute) {
	for _, route := range routes {
		fmt.Printf("%+v\n", route)
		if route.Prefix {
			authSvc.router.
				Name(route.Name).
				PathPrefix(route.Pattern)
		} else {
			authSvc.router.
				Name(route.Name).
				Methods(route.Method).
				Path(route.Pattern)
		}
	}
}

func sessionTimeoutCb(session *sm.Session) {
	log.Info("Session timed out. ID[", session.ID, "] Username[", session.Username, "]")
	var metric met.SessionMetric
	metric.Provider = session.Provider
	metric.User = session.Username
	metric.Sandbox = session.Sandbox
	_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeTimeout, metric)

	metricSessionTimeout.Inc()

	// Destroy session sandbox
	if session.Sandbox != "" {
		_, err := authSvc.pfmCtrlClient.SandboxControlApi.DeleteSandbox(context.TODO(), session.Sandbox)
		if err == nil {
			metricSessionActive.Dec()
			metricSessionDuration.Observe(time.Since(session.StartTime).Minutes())
		}
	}
}

// Generate a random state string
func generateState(n int) (string, error) {
	data := make([]byte, n)
	if _, err := io.ReadFull(rand.Reader, data); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(data), nil
}

func getUniqueState() (state string, err error) {
	for i := 0; i < 3; i++ {
		// Get random state
		randState, err := generateState(20)
		if err != nil {
			log.Error(err.Error())
			return "", err
		}

		// Make sure state is unique
		if _, found := authSvc.loginRequests[randState]; !found {
			return randState, nil
		}
	}
	return "", errors.New("Failed to generate a random state string")
}

func getLoginRequest(state string) *LoginRequest {
	mutex.Lock()
	defer mutex.Unlock()
	request, found := authSvc.loginRequests[state]
	if !found {
		return nil
	}
	return request
}

func setLoginRequest(state string, request *LoginRequest) {
	mutex.Lock()
	defer mutex.Unlock()
	authSvc.loginRequests[state] = request
}

func delLoginRequest(state string) {
	mutex.Lock()
	defer mutex.Unlock()
	request, found := authSvc.loginRequests[state]
	if !found {
		return
	}
	if request.timer != nil {
		request.timer.Stop()
	}
	delete(authSvc.loginRequests, state)
}

func getErrUrl(err string) string {
	return authSvc.uri + "?err=" + strings.ReplaceAll(err, " ", "+")
}

// ----------  REST API  ----------

func asAuthenticate(w http.ResponseWriter, r *http.Request) {

	// Get service & sandbox name from request query parameters
	query := r.URL.Query()
	svcName := query.Get("svc")
	var sboxName string

	// Get original request URL & method
	originalUrl := r.Header.Get("X-Original-URL")
	originalMethod := r.Header.Get("X-Original-Method")
	if originalUrl == "" || originalMethod == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Update request URL & method before running through matchers
	var err error
	r.Method = originalMethod
	r.URL, err = url.ParseRequestURI(originalUrl)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get permissions for matching route or default if not found
	var permission *Permission
	var match mux.RouteMatch
	if authSvc.router.Match(r, &match) {
		routeName := match.Route.GetName()
		sboxName = match.Vars["sbox"]
		mepName := match.Vars["mep"]
		log.Debug("routeName: ", routeName, " sboxName: ", sboxName, " mepName: ", mepName)

		// Check service-specific routes
		if svcName != "" {
			if svcPermissions, found := authSvc.cache.Services[svcName]; found {
				permission = svcPermissions[routeName]
			}
		}
		// Check file servers if not already found
		if permission == nil {
			if fsPermission, found := authSvc.cache.Fileservers[routeName]; found {
				permission = fsPermission
			}
		}
	} else {
		permission = authSvc.cache.Default
	}

	// Verify permission
	if permission == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Handle according to permission mode
	switch permission.Mode {
	case sm.ModeAllow:
		// break
	case sm.ModeBlock:
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	case sm.ModeVerify:
		// Retrieve user session, if any
		session, err := authSvc.sessionMgr.GetSessionStore().Get(r)
		if err != nil || session == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Verify role permissions
		role := session.Role
		if role == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		access := permission.Roles[role]
		if access != sm.AccessGranted {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// For non-admin users, verify session sandbox matches service sandbox, if any
		if session.Role != sm.RoleAdmin && sboxName != "" && sboxName != session.Sandbox {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

	default:
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Allow request
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func asAuthorize(w http.ResponseWriter, r *http.Request) {
	var metric met.SessionMetric

	// Retrieve query parameters
	query := r.URL.Query()
	code := query.Get("code")
	state := query.Get("state")

	// Validate request state
	request := getLoginRequest(state)
	if request == nil {
		err := errors.New("Invalid OAuth state")
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
		metricSessionFail.WithLabelValues("OAuth").Inc()
		return
	}

	// Get provider-specific OAuth config
	provider := request.provider
	config, found := authSvc.oauthConfigs[provider]
	if !found {
		err := errors.New("Provider config not found for: " + provider)
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
		metricSessionFail.WithLabelValues("Internal").Inc()
		return
	}
	metric.Provider = provider

	// Delete login request & timer
	delLoginRequest(state)

	// Retrieve access token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
		metricSessionFail.WithLabelValues("Internal").Inc()
		return
	}

	oauthClient := config.Client(context.Background(), token)
	if oauthClient == nil {
		err = errors.New("Failed to create new oauth client")
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
		metricSessionFail.WithLabelValues("OAuth").Inc()
		return
	}

	// Retrieve User ID
	var userId string
	switch provider {
	case OAUTH_PROVIDER_GITHUB:
		client := github.NewClient(oauthClient)
		if client == nil {
			err = errors.New("Failed to create new GitHub client")
			log.Error(err.Error())
			metric.Description = err.Error()
			_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
			http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
			metricSessionFail.WithLabelValues("OAuth").Inc()
			return
		}
		user, _, err := client.Users.Get(context.Background(), "")
		if err != nil {
			log.Error(err.Error())
			metric.Description = err.Error()
			_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
			http.Redirect(w, r, getErrUrl("Failed to retrieve GitHub user ID"), http.StatusFound)
			metricSessionFail.WithLabelValues("OAuth").Inc()
			return
		}
		userId = *user.Login

	case OAUTH_PROVIDER_GITLAB:
		client := gitlab.NewOAuthClient(oauthClient, token.AccessToken)
		if client == nil {
			err = errors.New("Failed to create new GitLab client")
			log.Error(err.Error())
			metric.Description = err.Error()
			_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
			http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
			metricSessionFail.WithLabelValues("OAuth").Inc()
			return
		}

		// Override default gitlab base URL
		if gitlabApiUrl != "" {
			err = client.SetBaseURL(gitlabApiUrl)
			if err != nil {
				log.Error(err.Error())
				metric.Description = err.Error()
				_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
				http.Redirect(w, r, getErrUrl("Failed to set GitLab API base url"), http.StatusFound)
				metricSessionFail.WithLabelValues("OAuth").Inc()
				return
			}
		}

		user, _, err := client.Users.CurrentUser()
		if err != nil {
			log.Error(err.Error())
			metric.Description = err.Error()
			_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
			http.Redirect(w, r, getErrUrl("Failed to retrieve GitLab user ID"), http.StatusFound)
			metricSessionFail.WithLabelValues("OAuth").Inc()
			return
		}
		userId = user.Username
	default:
	}
	metric.User = userId

	createSandbox, err := strconv.ParseBool(request.createSandbox)
	if err != nil {
		createSandbox = false
	}

	// Start user session
	sandboxName, isNew, userRole, err, errCode := startSession(provider, userId, w, r, createSandbox)
	if err != nil {
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Redirect(w, r, getErrUrl(err.Error()), errCode)
		metricSessionFail.WithLabelValues("Session").Inc()
		return
	}

	metric.Sandbox = sandboxName
	_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeLogin, metric)

	// Redirect user to sandbox
	http.Redirect(w, r, authSvc.uri+"?sbox="+sandboxName+"&user="+userId+"&role="+userRole, http.StatusFound)
	metricSessionSuccess.Inc()
	if isNew {
		metricSessionActive.Inc()
	}
}

func asLogin(w http.ResponseWriter, r *http.Request) {
	log.Info("----- OAUTH LOGIN -----")
	var metric met.SessionMetric
	metricSessionLogin.WithLabelValues("OAuth").Inc()

	// Retrieve query parameters
	query := r.URL.Query()
	provider := query.Get("provider")
	createSandbox := query.Get("sbox")
	metric.Provider = provider

	// Get provider-specific OAuth config
	config, found := authSvc.oauthConfigs[provider]
	if !found {
		err := errors.New("Provider config not found for: " + provider)
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
		metricSessionFail.WithLabelValues("Internal").Inc()
		return
	}

	// Generate unique random state string
	state, err := getUniqueState()
	if err != nil {
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Redirect(w, r, getErrUrl(err.Error()), http.StatusFound)
		metricSessionFail.WithLabelValues("Internal").Inc()
		return
	}

	// Track oauth request & handle
	request := &LoginRequest{
		provider:      provider,
		createSandbox: createSandbox,
		timer:         time.NewTimer(10 * time.Minute),
	}
	setLoginRequest(state, request)

	// Start timer to remove request from map
	go func() {
		<-request.timer.C
		delLoginRequest(state)
	}()

	// Generate provider-specific oauth redirect
	uri := config.AuthCodeURL(state, oauth2.AccessTypeOnline)
	http.Redirect(w, r, uri, http.StatusFound)
}

func asLoginUser(w http.ResponseWriter, r *http.Request) {
	log.Info("----- LOGIN -----")
	var metric met.SessionMetric
	metricSessionLogin.WithLabelValues("Basic").Inc()

	// Get form data
	username := r.FormValue("username")
	password := r.FormValue("password")

	metric.Provider = OAUTH_PROVIDER_LOCAL
	metric.User = username

	// Validate user credentials
	authenticated, err := authSvc.userStore.AuthenticateUser(OAUTH_PROVIDER_LOCAL, username, password)
	if err != nil || !authenticated {
		if err != nil {
			metric.Description = err.Error()
		} else {
			metric.Description = "Unauthorized"
		}
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Start user session
	sandboxName, isNew, _, err, errCode := startSession(OAUTH_PROVIDER_LOCAL, username, w, r, false)
	if err != nil {
		log.Error(err.Error())
		metric.Description = err.Error()
		_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeError, metric)
		http.Error(w, err.Error(), errCode)
		return
	}

	metric.Sandbox = sandboxName
	_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeLogin, metric)
	if isNew {
		metricSessionActive.Inc()
	}

	// Prepare response
	var sandbox dataModel.Sandbox
	sandbox.Name = sandboxName

	// Format response
	jsonResponse, err := json.Marshal(sandbox)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(jsonResponse))
}

// Retrieve existing user session or create a new one
func startSession(provider string, username string, w http.ResponseWriter, r *http.Request, createSandbox bool) (sandboxName string, isNew bool, userRole string, err error, code int) {

	// Get existing session by user name, if any
	sessionStore := authSvc.sessionMgr.GetSessionStore()
	session, err := sessionStore.GetByName(provider, username)
	if err != nil {
		// Check if max session count is reached before creating a new one
		count := sessionStore.GetCount()
		if count >= authSvc.maxSessions {
			err = errors.New("Maximum session count exceeded")
			return "", false, "", err, http.StatusServiceUnavailable
		}

		// Get requested sandbox name & role from user profile, if any
		providerMode := strings.TrimSpace(os.Getenv("MEEP_OAUTH_PROVIDER_MODE"))
		role := users.RoleUser
		user, err := authSvc.userStore.GetUser(provider, username)
		if err == nil {
			sandboxName = user.Sboxname
			role = user.Role
		}
		if err != nil && providerMode == providerModeSecure {
			return "", false, "", err, http.StatusUnauthorized
		}

		// Create sandbox
		if createSandbox {
			var sandboxConfig pcc.SandboxConfig
			if sandboxName == "" {
				sandbox, _, err := authSvc.pfmCtrlClient.SandboxControlApi.CreateSandbox(context.TODO(), sandboxConfig)
				if err != nil {
					return "", false, "", err, http.StatusInternalServerError
				}
				sandboxName = sandbox.Name
			} else {
				_, err := authSvc.pfmCtrlClient.SandboxControlApi.CreateSandboxWithName(context.TODO(), sandboxName, sandboxConfig)
				if err != nil {
					return "", false, "", err, http.StatusInternalServerError
				}
			}
		}

		// Create new session
		session = new(sm.Session)
		session.ID = ""
		session.Username = username
		session.Provider = provider
		session.Sandbox = sandboxName
		session.Role = role
		isNew = true
	} else {
		sandboxName = session.Sandbox
	}
	userRole = session.Role

	// Set session
	err, code = sessionStore.Set(session, w, r)
	if err != nil {
		log.Error("Failed to set session with err: ", err.Error())
		// Remove newly created sandbox on failure
		if session.ID == "" && createSandbox {
			_, _ = authSvc.pfmCtrlClient.SandboxControlApi.DeleteSandbox(context.TODO(), sandboxName)
		}
		return "", false, "", err, code
	}
	return sandboxName, isNew, userRole, nil, http.StatusOK
}

func asLogout(w http.ResponseWriter, r *http.Request) {
	log.Info("----- LOGOUT -----")
	var metric met.SessionMetric
	sandboxDeleted := false
	metricSessionLogout.Inc()

	// Get existing session
	sessionStore := authSvc.sessionMgr.GetSessionStore()
	session, err := sessionStore.Get(r)
	if err == nil {
		metric.Provider = session.Provider
		metric.User = session.Username
		metric.Sandbox = session.Sandbox

		// Delete sandbox
		if session.Sandbox != "" {
			_, err = authSvc.pfmCtrlClient.SandboxControlApi.DeleteSandbox(context.TODO(), session.Sandbox)
			if err == nil {
				sandboxDeleted = true
			}
		}
	}

	// Delete session
	err, code := sessionStore.Del(w, r)
	if err != nil {
		log.Error("Failed to delete session with err: ", err.Error())
		http.Error(w, err.Error(), code)
		return
	}

	_ = authSvc.metricStore.SetSessionMetric(met.SesMetTypeLogout, metric)
	if sandboxDeleted {
		metricSessionActive.Dec()
		metricSessionDuration.Observe(time.Since(session.StartTime).Minutes())
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func asTriggerWatchdog(w http.ResponseWriter, r *http.Request) {
	// Refresh session
	sessionStore := authSvc.sessionMgr.GetSessionStore()
	err, code := sessionStore.Refresh(w, r)
	if err != nil {
		log.Error("Failed to refresh session with err: ", err.Error())
		http.Error(w, err.Error(), code)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
* Response Code 200: Login is Supported and Session exists
* Response Code 401: Login is Supported and Session doesn't exists
* Response Code 404: Login is not Supported
 */
func asLoginSupported(w http.ResponseWriter, r *http.Request) {
	log.Info("----- LOGIN SUPPORTED-----")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Check if Github is enabled
	githubEnabledStr := strings.TrimSpace(os.Getenv("MEEP_OAUTH_GITHUB_ENABLED"))
	githubEnabled, err := strconv.ParseBool(githubEnabledStr)
	if err != nil || !githubEnabled {
		w.WriteHeader(http.StatusNotFound)
	} else {
		// Retrieve user session, if any
		session, err := authSvc.sessionMgr.GetSessionStore().Get(r)
		if err != nil || session == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
