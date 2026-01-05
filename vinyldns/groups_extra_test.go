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

func TestGroupChange(t *testing.T) {
	changeJSON := `{
		"id": "c1",
		"userId": "u1",
		"userName": "testuser",
		"changeType": "Update",
		"created": "now",
		"groupChangeMessage": "updated group",
		"newGroup": {"id": "g1", "name": "test-group"},
		"oldGroup": {"id": "g1", "name": "test-group"}
	}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/groups/change/c1",
			code:     200,
			body:     changeJSON,
		},
	})
	defer server.Close()

	change, err := client.GroupChange("c1")
	if err != nil {
		t.Error(err)
	}
	if change.ID != "c1" {
		t.Error("Expected group change ID to be c1")
	}
}

func TestGroupValidDomains(t *testing.T) {
	domainsJSON := `["gmail.com","test.com"]`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/groups/valid/domains",
			code:     200,
			body:     domainsJSON,
		},
	})
	defer server.Close()

	domains, err := client.GroupValidDomains()
	if err != nil {
		t.Error(err)
	}
	if len(domains) != 2 {
		t.Error("Expected valid domains length to be 2")
	}
}
