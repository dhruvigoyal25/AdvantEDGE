/*
 * Copyright (c) 2022  InterDigital Communications, Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"encoding/json"

	log "github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger"
)

func ConvertMobilityProcedureSubscriptionToInlineSubscription(src *MobilityProcedureSubscription) *InlineSubscription {
	var dst InlineSubscription
	srcJson, err := json.Marshal(*src)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	err = json.Unmarshal([]byte(srcJson), &dst)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return &dst
}

func ConvertAdjacentAppInfoSubscriptionToInlineSubscription(src *AdjacentAppInfoSubscription) *InlineSubscription {
	var dst InlineSubscription
	srcJson, err := json.Marshal(*src)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	err = json.Unmarshal([]byte(srcJson), &dst)
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	return &dst
}
