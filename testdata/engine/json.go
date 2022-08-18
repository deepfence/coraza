// Copyright 2022 Juan Pablo Tosso
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package engine

import (
	"github.com/corazawaf/coraza/v3/testing/profile"
)

var _ = profile.RegisterProfile(profile.Profile{
	Meta: profile.Meta{
		Author:      "jptosso",
		Description: "Test if the json request body work",
		Enabled:     true,
		Name:        "jsonyaml",
	},
	Tests: []profile.Test{
		{
			Title: "json",
			Stages: []profile.Stage{
				{
					Stage: profile.SubStage{
						Input: profile.StageInput{
							URI:    "/index.php?json.test=456",
							Method: "POST",
							Headers: map[string]string{
								"content-type": "application/json",
							},
							Data: `{"test":123, "test2": 456, "test3": [22, 44, 55]}`,
						},
						Output: profile.ExpectedOutput{
							TriggeredRules: []int{
								100,
								101,
								1100,
								1101,
								1010,
							},
							NonTriggeredRules: []int{
								1111,
								1102,
								103,
							},
						},
					},
				},
			},
		},
	},
	Rules: `
SecRequestBodyAccess On
SecRule REQUEST_HEADERS:content-type "application/json" "id: 100, phase:1, pass, log, ctl:requestBodyProcessor=JSON"
SecRule REQBODY_PROCESSOR "JSON" "id: 101,phase:2,log,block"

SecRule REQBODY_ERROR "!@eq 0" "id:1111, phase:2, log, block"

SecRule REQUEST_BODY "456" "id:103, phase:2, log"
SecRule ARGS:json.test "@eq 123" "id:1100, phase:2, log, block"
SecRule ARGS:json.test3.2 "@eq 55" "id:1101, phase:2, log, block"

# We test for some vulnerability
SecRule ARGS:json.test "@eq 456" "id:1102, phase:2, log, block"

SecRule ARGS:json.test3 "@eq 3" "id: 1010, phase:2, log, block"
`,
})