/*
Copyright 2018 Comcast Cable Communications Management, LLC
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
	"testing"

	"github.com/gobs/pretty"
)

func TestRecordSets(t *testing.T) {
	recordSetsJSON, err := readFile("test-fixtures/recordsets/recordsets.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets",
			code:     200,
			body:     recordSetsJSON,
		},
	})

	defer server.Close()

	rs, err := client.RecordSets("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(rs))
		t.Error(err)
	}
	if len(rs) != 2 {
		t.Error("Expected 2 Record Sets")
	}
	for _, r := range rs {
		if r.ID == "" {
			t.Error("Expected RecordSet.Id to have a value")
		}
	}
}

func TestRecordSetsListAll(t *testing.T) {
	recordSetsListJSON1, err := readFile("test-fixtures/recordsets/recordsets-list-json-1.json")
	if err != nil {
		t.Error(err)
	}
	recordSetsListJSON2, err := readFile("test-fixtures/recordsets/recordsets-list-json-2.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets?maxItems=1",
			code:     200,
			body:     recordSetsListJSON1,
		},
		{
			endpoint: "http://host.com/zones/123/recordsets?maxItems=1&startFrom=2",
			code:     200,
			body:     recordSetsListJSON2,
		},
	})

	defer server.Close()

	if _, err := client.RecordSetsListAll("123", ListFilter{
		MaxItems: 200,
	}); err == nil {
		t.Error("Expected error -- MaxItems must be between 1 and 100")
	}

	records, err := client.RecordSetsListAll("123", ListFilter{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(records) != 2 {
		t.Error("Expected 2 records; got ", len(records))
	}

	if records[0].ID != "1" {
		t.Error("Expected RecordSet.ID to be 1")
	}

	if records[1].ID != "2" {
		t.Error("Expected RecordSet.ID to be 2")
	}
}

func TestRecordSetsListAllWhenNoneExist(t *testing.T) {
	recordSetsListNoneJSON, err := readFile("test-fixtures/recordsets/recordsets-list-none.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets",
			code:     200,
			body:     recordSetsListNoneJSON,
		},
	})

	defer server.Close()

	records, err := client.RecordSetsListAll("123", ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(records) != 0 {
		t.Error("Expected 0 records; got ", len(records))
	}

	j, err := json.Marshal(records)
	if err != nil {
		t.Error(err)
	}

	if string(j) != "[]" {
		t.Error("Expected string-converted marshaled JSON to be '[]'; got ", string(j))
	}
}

func TestRecordSetsGlobalListAll(t *testing.T) {
	recordSetsListJSON1, err := readFile("test-fixtures/recordsets/recordsets-list-json-1.json")
	if err != nil {
		t.Error(err)
	}
	recordSetsListJSON2, err := readFile("test-fixtures/recordsets/recordsets-list-json-2.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/recordsets?maxItems=1",
			code:     200,
			body:     recordSetsListJSON1,
		},
		{
			endpoint: "http://host.com/recordsets?maxItems=1&startFrom=2",
			code:     200,
			body:     recordSetsListJSON2,
		},
	})

	defer server.Close()

	if _, err := client.RecordSetsGlobalListAll(GlobalListFilter{
		MaxItems: 200,
	}); err == nil {
		t.Error("Expected error -- MaxItems must be between 1 and 100")
	}

	records, err := client.RecordSetsGlobalListAll(GlobalListFilter{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(records) != 2 {
		t.Error("Expected 2 records; got ", len(records))
	}

	if records[0].ID != "1" {
		t.Error("Expected RecordSet.ID to be 1")
	}

	if records[1].ID != "2" {
		t.Error("Expected RecordSet.ID to be 2")
	}
}

func TestRecordSetsGlobal(t *testing.T) {
	recordSetsListJSON1, err := readFile("test-fixtures/recordsets/recordsets-list-json-1.json")
	if err != nil {
		t.Error(err)
	}
	recordSetsListJSON2, err := readFile("test-fixtures/recordsets/recordsets-list-json-2.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/recordsets?maxItems=1",
			code:     200,
			body:     recordSetsListJSON1,
		},
		{
			endpoint: "http://host.com/recordsets?startFrom=2&maxItems=1",
			code:     200,
			body:     recordSetsListJSON2,
		},
	})

	defer server.Close()

	if _, _, err := client.RecordSetsGlobal(GlobalListFilter{
		MaxItems: 200,
	}); err == nil {
		t.Error("Expected error -- MaxItems must be between 1 and 100")
	}

	records, nextID, err := client.RecordSetsGlobal(GlobalListFilter{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(records) != 1 {
		t.Error("Expected 1 records; got ", len(records))
	}

	if records[0].ID != "1" {
		t.Error("Expected RecordSet.ID to be 1")
	}
	if nextID != "2" {
		t.Error("Expected nextId to be 2")
	}
}

func TestRecordSetsGlobalListAllWhenNoneExist(t *testing.T) {
	recordSetsListNoneJSON, err := readFile("test-fixtures/recordsets/recordsets-list-none.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/recordsets",
			code:     200,
			body:     recordSetsListNoneJSON,
		},
	})

	defer server.Close()

	records, err := client.RecordSetsGlobalListAll(GlobalListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(records) != 0 {
		t.Error("Expected 0 records; got ", len(records))
	}

	j, err := json.Marshal(records)
	if err != nil {
		t.Error(err)
	}

	if string(j) != "[]" {
		t.Error("Expected string-converted marshaled JSON to be '[]'; got ", string(j))
	}
}

func TestRecordSetCollector(t *testing.T) {
	recordSetsJSON, err := readFile("test-fixtures/recordsets/recordsets.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets?maxItems=3",
			code:     200,
			body:     recordSetsJSON,
		},
		{
			endpoint: "http://host.com/zones/123/recordsets?maxItems=1",
			code:     200,
			body:     recordSetsJSON,
		},
	})

	defer server.Close()

	if _, err := client.RecordSetCollector("123", 999); err == nil {
		t.Error("Expected error -- over max number of records allowed")
	}

	// test under limit
	collector, err := client.RecordSetCollector("123", 3)
	if err != nil {
		t.Error(err)
	}

	rs, _ := collector()
	if len(rs) != 2 {
		t.Log(pretty.PrettyFormat(rs))
		t.Error("Expected 2 Record Sets, got ", len(rs))
	}
	for _, r := range rs {
		if r.ID == "" {
			t.Error("Expected RecordSet.Id to have a value")
		}
	}

	// test over limit, but under max
	collector, err = client.RecordSetCollector("123", 1)
	if err != nil {
		t.Error(err)
	}

	rs, _ = collector()
	if len(rs) != 1 {
		t.Log(pretty.PrettyFormat(rs))
		t.Error("Expected 1 Record Sets, got ", len(rs))
	}
}

func TestRecordSet(t *testing.T) {
	recordSetJSON, err := readFile("test-fixtures/recordsets/recordset.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets/456",
			code:     200,
			body:     recordSetJSON,
		},
	})

	defer server.Close()

	rs, err := client.RecordSet("123", "456")
	if err != nil {
		t.Log(pretty.PrettyFormat(rs))
		t.Error(err)
	}
	if rs.ID != "123" {
		t.Error("Expected RecordSet.Id to have a value")
	}
	if rs.OwnerGroupID != "789" {
		t.Error("Expected RecordSet.OwnerGroupID to have a value")
	}
	if rs.ZoneID != "456" {
		t.Error("Expected RecordSet.ZoneId to have a value")
	}
	if rs.Name != "test-01" {
		t.Error("Expected RecordSet.Name to have a value")
	}
	if rs.Type != "A" {
		t.Error("Expected RecordSet.Type to have a value")
	}
	if rs.Status != "Active" {
		t.Error("Expected RecordSet.Status to have a value")
	}
	if rs.Created != "2015-11-02T13:41:54Z" {
		t.Error("Expected RecordSet.Status to have a value")
	}
	if rs.Updated != "2015-11-02T13:41:57Z" {
		t.Error("Expected RecordSet.Status to have a value")
	}
	if rs.TTL != 200 {
		t.Error("Expected RecordSet.Ttl to have a value")
	}
	if rs.Records[0].Address != "127.0.0.1" {
		t.Error("Expected RecordSet.Address to have a value")
	}
	if rs.Account != "vinyldns" {
		t.Error("Expected RecordSet.Account to have a value")
	}
}

func TestRecordSetCreate(t *testing.T) {
	recordSetUpdateResponseJSON, err := readFile("test-fixtures/recordsets/recordset-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets",
			code:     200,
			body:     recordSetUpdateResponseJSON,
		},
	})

	defer server.Close()

	rs := &RecordSet{
		ZoneID: "123",
		Name:   "name",
		Type:   "CNAME",
		TTL:    200,
		Records: []Record{{
			CName: "cname",
		}},
	}

	r, err := client.RecordSetCreate(rs)
	if err != nil {
		t.Log(pretty.PrettyFormat(r))
		t.Error(err)
	}
	if r.ChangeID != "b3d4e0a9-a081-4adc-9a95-3ec2e7d26635" {
		t.Error("Expected recordSetUpdateResponse.ChangeId to have a value")
	}
	if r.Status != "Pending" {
		t.Error("Expected recordSetUpdateResponse.Status to have a value")
	}
	if r.Zone.Name != "vinyldnstest.sys.vinyldns.net." {
		t.Error("Expected recordSetUpdateResponse.Zone.Name to have a value")
	}
	if r.RecordSet.Name != "foo." {
		t.Error("Expected recordSetUpdateResponse.RecordSet.Name to have a value")
	}
}

func TestRecordSetUpdate(t *testing.T) {
	recordSetUpdateResponseJSON, err := readFile("test-fixtures/recordsets/recordset-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets/456",
			code:     202,
			body:     recordSetUpdateResponseJSON,
		},
	})

	defer server.Close()

	rs := &RecordSet{
		ZoneID: "123",
		ID:     "456",
		Name:   "name",
		Type:   "CNAME",
		TTL:    200,
		Records: []Record{{
			CName: "cname",
		}},
	}

	r, err := client.RecordSetUpdate(rs)
	if err != nil {
		t.Log(pretty.PrettyFormat(r))
		t.Error(err)
	}
	if r.ChangeID != "b3d4e0a9-a081-4adc-9a95-3ec2e7d26635" {
		t.Error("Expected recordSetUpdateResponse.ChangeId to have a value")
	}
	if r.Status != "Pending" {
		t.Error("Expected recordSetUpdateResponse.Status to have a value")
	}
	if r.Zone.Name != "vinyldnstest.sys.vinyldns.net." {
		t.Error("Expected recordSetUpdateResponse.Zone.Name to have a value")
	}
	if r.RecordSet.Name != "foo." {
		t.Error("Expected recordSetUpdateResponse.RecordSet.Name to have a value")
	}
}

func TestRecordSetDelete(t *testing.T) {
	recordSetUpdateResponseJSON, err := readFile("test-fixtures/recordsets/recordset-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets/456",
			code:     202,
			body:     recordSetUpdateResponseJSON,
		},
	})

	defer server.Close()

	r, err := client.RecordSetDelete("123", "456")
	if err != nil {
		t.Log(pretty.PrettyFormat(r))
		t.Error(err)
	}
	if r.ChangeID != "b3d4e0a9-a081-4adc-9a95-3ec2e7d26635" {
		t.Error("Expected recordSetUpdateResponse.ChangeId to have a value")
	}
	if r.Status != "Pending" {
		t.Error("Expected recordSetUpdateResponse.Status to have a value")
	}
	if r.Zone.Name != "vinyldnstest.sys.vinyldns.net." {
		t.Error("Expected recordSetUpdateResponse.Zone.Name to have a value")
	}
	if r.RecordSet.Name != "foo." {
		t.Error("Expected recordSetUpdateResponse.RecordSet.Name to have a value")
	}
}

func TestRecordSetChangessListAllWhenNoneExist(t *testing.T) {
	recordSetChangesListNoneJSON, err := readFile("test-fixtures/recordsets/recordset-changes-list-none.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsetchanges",
			code:     200,
			body:     recordSetChangesListNoneJSON,
		},
	})

	defer server.Close()

	changes, err := client.RecordSetChangesListAll("123", ListFilterInt{})
	if err != nil {
		t.Error(err)
	}

	if len(changes) != 0 {
		t.Error("Expected 0 changes; got ", len(changes))
	}

	j, err := json.Marshal(changes)
	if err != nil {
		t.Error(err)
	}

	if string(j) != "[]" {
		t.Error("Expected string-converted marshaled JSON to be '[]'; got ", string(j))
	}
}

func TestRecordSetChangesListAll(t *testing.T) {
	recordSetChangesJSON1, err := readFile("test-fixtures/recordsets/recordset-changes-list-1.json")
	if err != nil {
		t.Error(err)
	}
	recordSetChangesJSON2, err := readFile("test-fixtures/recordsets/recordset-changes-list-2.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsetchanges?maxItems=1",
			code:     200,
			body:     recordSetChangesJSON1,
		},
		{
			endpoint: "http://host.com/zones/123/recordsetchanges?maxItems=1&startFrom=2",
			code:     200,
			body:     recordSetChangesJSON2,
		},
	})

	defer server.Close()

	if _, err := client.RecordSetChangesListAll("123", ListFilterInt{
		MaxItems: 200,
	}); err == nil {
		t.Error("Expected error -- MaxItems must be between 1 and 100")
	}

	changes, err := client.RecordSetChangesListAll("123", ListFilterInt{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(changes) != 2 {
		t.Error("Expected 2 records; got ", len(changes))
	}

	if changes[0].ID != "1" {
		t.Error("Expected RecordSetChange.ID to be 1")
	}

	if changes[1].ID != "2" {
		t.Error("Expected RecordSetChange.ID to be 2")
	}
}

func TestRecordSetChange(t *testing.T) {
	recordSetChangeJSON, err := readFile("test-fixtures/recordsets/recordset-change.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/recordsets/456/changes/789",
			code:     200,
			body:     recordSetChangeJSON,
		},
	})

	defer server.Close()

	rsc, err := client.RecordSetChange("123", "456", "789")
	if err != nil {
		t.Log(pretty.PrettyFormat(rsc))
		t.Error(err)
	}
	if rsc.ChangeType != "Create" {
		t.Error("Expected RecordSetChange.ChangeType to have a value")
	}
}
