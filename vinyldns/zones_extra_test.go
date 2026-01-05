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

func TestZoneBackendIDs(t *testing.T) {
	idsJSON := `["backend-1","backend-2"]`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/backendids",
			code:     200,
			body:     idsJSON,
		},
	})
	defer server.Close()

	ids, err := client.ZoneBackendIDs()
	if err != nil {
		t.Error(err)
	}
	if len(ids) != 2 {
		t.Error("Expected backend IDs to have length 2")
	}
}

func TestZonesDeleted(t *testing.T) {
	deletedJSON := `{
		"zonesDeletedInfo": [
			{
				"zoneChange": {
					"zone": {
						"name": "ok.",
						"email": "test@test.com",
						"status": "Deleted",
						"id": "z1"
					},
					"userId": "u1",
					"changeType": "Delete",
					"status": "Synced",
					"created": "now",
					"id": "c1"
				},
				"adminGroupName": "admins",
				"userName": "test",
				"accessLevel": "NoAccess"
			}
		],
		"maxItems": 100,
		"ignoreAccess": true
	}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/deleted/changes",
			code:     200,
			body:     deletedJSON,
		},
	})
	defer server.Close()

	resp, err := client.ZonesDeleted(DeletedZonesFilter{})
	if err != nil {
		t.Error(err)
	}
	if len(resp.ZonesDeletedInfo) != 1 {
		t.Error("Expected zonesDeletedInfo to have one entry")
	}
}

func TestZoneACLRuleCreate(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/acl/rules",
			code:     202,
			body:     zoneJSON,
		},
	})
	defer server.Close()

	resp, err := client.ZoneACLRuleCreate("123", &ACLRule{
		AccessLevel: "Write",
		GroupID:     "456",
	})
	if err != nil {
		t.Error(err)
	}
	if resp.Status == "" {
		t.Error("Expected zone update status to have a value")
	}
}

func TestZoneACLRuleDelete(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/acl/rules",
			code:     202,
			body:     zoneJSON,
		},
	})
	defer server.Close()

	resp, err := client.ZoneACLRuleDelete("123", &ACLRule{
		AccessLevel: "Write",
		GroupID:     "456",
	})
	if err != nil {
		t.Error(err)
	}
	if resp.Status == "" {
		t.Error("Expected zone update status to have a value")
	}
}

func TestZoneChangesFailure(t *testing.T) {
	failuresJSON := `{
		"failedZoneChanges": [
			{
				"zone": {
					"name": "ok.",
					"email": "test@test.com",
					"status": "Active",
					"id": "z1"
				},
				"userId": "u1",
				"changeType": "Sync",
				"status": "Failed",
				"created": "now",
				"id": "c1"
			}
		],
		"maxItems": 100
	}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/metrics/health/zonechangesfailure",
			code:     200,
			body:     failuresJSON,
		},
	})
	defer server.Close()

	resp, err := client.ZoneChangesFailure(ListFilter{})
	if err != nil {
		t.Error(err)
	}
	if len(resp.FailedZoneChanges) != 1 {
		t.Error("Expected failedZoneChanges to have one entry")
	}
}
