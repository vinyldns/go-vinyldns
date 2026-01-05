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

func TestRecordSetCount(t *testing.T) {
	countJSON := `{"count": 10}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsetcount",
			code:     200,
			body:     countJSON,
		},
	})
	defer server.Close()

	count, err := client.RecordSetCount("123")
	if err != nil {
		t.Error(err)
	}
	if count.Count != 10 {
		t.Error("Expected record set count to be 10")
	}
}

func TestRecordSetChangeHistory(t *testing.T) {
	historyJSON := `{
		"zoneId": "z1",
		"recordSetChanges": [
			{
				"zone": {"name": "ok.", "email": "test@test.com", "status": "Active", "id": "z1"},
				"recordSet": {"name": "ok.", "type": "A", "zoneId": "z1", "ttl": 300, "id": "rs1", "account": "system", "records": []},
				"userId": "u1",
				"changeType": "Update",
				"status": "Complete",
				"created": "now",
				"id": "c1",
				"userName": "testuser"
			}
		],
		"maxItems": 100
	}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/recordsetchange/history?fqdn=ok.&recordType=A&zoneId=z1",
			code:     200,
			body:     historyJSON,
		},
	})
	defer server.Close()

	history, err := client.RecordSetChangeHistory(RecordSetChangeHistoryFilter{
		ZoneID:     "z1",
		FQDN:       "ok.",
		RecordType: "A",
	})
	if err != nil {
		t.Error(err)
	}
	if len(history.RecordSetChanges) != 1 {
		t.Error("Expected record set changes to have one entry")
	}
}

func TestRecordSetChangesFailure(t *testing.T) {
	failuresJSON := `{
		"failedRecordSetChanges": [
			{
				"zone": {"name": "ok.", "email": "test@test.com", "status": "Active", "id": "z1"},
				"recordSet": {"name": "ok.", "type": "A", "zoneId": "z1", "ttl": 300, "id": "rs1", "account": "system", "records": []},
				"userId": "u1",
				"changeType": "Create",
				"status": "Failed",
				"created": "now",
				"id": "c1"
			}
		],
		"maxItems": 100
	}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/metrics/health/zones/z1/recordsetchangesfailure",
			code:     200,
			body:     failuresJSON,
		},
	})
	defer server.Close()

	failures, err := client.RecordSetChangesFailure("z1", ListFilter{})
	if err != nil {
		t.Error(err)
	}
	if len(failures.FailedRecordSetChanges) != 1 {
		t.Error("Expected failed record set changes to have one entry")
	}
}
