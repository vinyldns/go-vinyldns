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

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/gobs/pretty"
)

func TestBatchRecordChanges(t *testing.T) {
	batchChangesJSON, err := readFile("test-fixtures/batch-changes/batch-changes.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/batchrecordchanges",
			code:     200,
			body:     batchChangesJSON,
		},
	})

	defer server.Close()

	changes, err := client.BatchRecordChanges()
	if err != nil {
		t.Log(pretty.PrettyFormat(changes))
		t.Error(err)
	}

	c := changes[0]
	if c.UserName != "vinyl201" {
		t.Error("Expected BatchRecordChanges[0].UserName to be 'vinyl201'")
	}
	if c.TotalChanges != 5 {
		t.Error("Expected BatchRecordChanges[0].TotalChanges to be '5'")
	}
}

func TestBatchRecordChangeEncoding(t *testing.T) {
	recordChangeJSON, err := readFile("test-fixtures/batch-changes/batch-recordchange-format.json")
	if err != nil {
		t.Error(err)
	}

	var shouldBe RecordChange
	err = json.Unmarshal([]byte(recordChangeJSON), &shouldBe)
	if err != nil {
		t.Error(err)
	}

	reference := RecordChange{
		ID:               "id",
		Status:           "status",
		Comments:         "comments",
		ChangeType:       "changeType",
		RecordName:       "recordName",
		TTL:              1,
		Type:             "type",
		ZoneName:         "zoneName",
		InputName:        "inputName",
		ZoneID:           "zoneId",
		TotalChanges:     1,
		UserName:         "userName",
		UserID:           "userId",
		CreatedTimestamp: "createdTimeStamp",
		Record: RecordData{
			Address:  "address",
			CName:    "cname",
			PTRDName: "ptrdname",
		},
		OwnerGroupID: "ownerGroupId",
	}

	if !reflect.DeepEqual(shouldBe, reference) {
		t.Log(shouldBe)
		t.Log(reference)
		t.Error("Expected unmarshalled batch recordchange to match manually created one")
	}
}

func TestBatchRecordChange(t *testing.T) {
	batchChangeJSON, err := readFile("test-fixtures/batch-changes/batch-change.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/batchrecordchanges/123",
			code:     200,
			body:     batchChangeJSON,
		},
	})

	defer server.Close()

	change, err := client.BatchRecordChange("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(change))
		t.Error(err)
	}

	c := change.Changes[0]
	if c.RecordName != "parent.com." {
		t.Error("Expected BatchRecordChange.Changes[0].RecordName to be 'parent.com.'")
	}
	if c.ZoneName != "parent.com." {
		t.Error("Expected BatchRecordChange.Changes[0].ZoneName to be 'parent.com.'")
	}
}

func TestBatchRecordChangeCreate(t *testing.T) {
	batchChangeCreateJSON, err := readFile("test-fixtures/batch-changes/batch-change-create.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/batchrecordchanges",
			code:     200,
			body:     batchChangeCreateJSON,
		},
	})

	defer server.Close()

	change := &BatchRecordChange{}
	changeResult, err := client.BatchRecordChangeCreate(change)
	if err != nil {
		t.Log(pretty.PrettyFormat(changeResult))
		t.Error(err)
	}

	c := changeResult.Changes[0]
	if c.ChangeType != "Add" {
		t.Error("Expected BatchRecordChangeCreate.Changes[0].ChangeType to be 'Add'")
	}
	if changeResult.Comments != "this is optional" {
		t.Error("Expected BatchRecordChangeCreate.Comments to be 'this is optional'")
	}
}
