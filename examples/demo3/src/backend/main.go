/*
 * MEC Use-Case 3 API
 *
 * This section describes a use case that the user can accomplish using the MEC Sandbox APIs from a MEC application
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package main

import (
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	server "github.com/InterDigitalInc/AdvantEDGE/example/demo3/src/server"
	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
	"github.com/gorilla/handlers"
)

// Initalize customized logger
func init() {
	log.MeepTextLogInit("Demo-3")
}

func main() {
	var (
		dir      string
		fileName string
		run      bool = true
		done     chan bool
	)

	// First element in os.Args is always the program name
	// Require at least 2 arugments to have a file name argument
	if len(os.Args) < 2 {
		log.Fatal("Missing parameter, provide file name!")
	}

	// Read configuration file path in command line arugments
	configPath := os.Args[1]
	dir = filepath.Dir(configPath)
	fileName = filepath.Base(configPath)

	go func() {

		port, err := server.Init(dir, fileName)
		if err != nil {
			log.Fatal("Failed to initalize Demo 3 ", err)
		}

		// Start demo 3 ticker
		server.Run(done)

		// Start demo 3 server
		router := server.NewRouter()
		methods := handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"})
		header := handlers.AllowedHeaders([]string{"content-type"})
		log.Fatal(http.ListenAndServe(port, handlers.CORS(methods, header)(router)))
		run = false
	}()

	// Listen for SIGKILL
	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
		<-sigchan
		log.Info("Waiting to shut down program !")
		run = false
	}()

	// Listen for demo 3 error exit program
	go func() {
		<-done
		run = false
	}()

	for {
		// Invoke graceful termination upon program kill
		if !run {
			log.Info("Invoking demo 3 graceful termination")
			server.Terminate()
			break
		}
		time.Sleep(time.Second)
	}
}
