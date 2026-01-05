/*
Copyright 2026 Comcast Cable Communications Management, LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package vinyldns

import "testing"

func TestStatus(t *testing.T) {
	statusJSON := `{"processingDisabled":false,"color":"blue","keyName":"vinyldns.","version":"0.21.3"}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/status",
			code:     200,
			body:     statusJSON,
		},
	})
	defer server.Close()

	status, err := client.Status()
	if err != nil {
		t.Error(err)
	}
	if status.Color != "blue" {
		t.Error("Expected status.Color to be blue")
	}
}

func TestStatusUpdate(t *testing.T) {
	statusJSON := `{"processingDisabled":true,"color":"blue","keyName":"vinyldns.","version":"0.21.3"}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/status?processingDisabled=true",
			code:     200,
			body:     statusJSON,
		},
	})
	defer server.Close()

	status, err := client.StatusUpdate(true)
	if err != nil {
		t.Error(err)
	}
	if !status.ProcessingDisabled {
		t.Error("Expected status.ProcessingDisabled to be true")
	}
}
