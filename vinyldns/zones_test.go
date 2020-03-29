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

func TestZones(t *testing.T) {
	zonesJSON, err := readFile("test-fixtures/zones/zones.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones",
			code:     202,
			body:     zonesJSON,
		},
	})
	defer server.Close()

	zones, err := client.Zones()
	if err != nil {
		t.Log(pretty.PrettyFormat(zones))
		t.Error(err)
	}
	if len(zones) != 2 {
		t.Error("Expected 2 Domains")
	}
	for _, z := range zones {
		if z.Name == "" {
			t.Error("Expected zone.Name to have a value")
		}
		if z.Email == "" {
			t.Error("Expected zone.Email to have a value")
		}
		if z.Status == "" {
			t.Error("Expected zone.Status to have a value")
		}
		if z.Created == "" {
			t.Error("Expected zone.Created to have a value")
		}
		if z.ID == "" {
			t.Error("Expected zone.ID to have a value")
		}
		if z.AdminGroupID == "" {
			t.Error("Expected zone.AdminGroupID to have a value")
		}
		if z.Account == "" {
			t.Error("Expected zone.Account to have a value")
		}
		if z.BackendID == "" {
			t.Error("Expected zone.BackendID to have a value")
		}
		if z.AccessLevel == "" {
			t.Error("Expected zone.AccessLevel to have a value")
		}
	}
}

func TestZonesListAll(t *testing.T) {
	zonesListJSON1, err := readFile("test-fixtures/zones/zones-list-1.json")
	if err != nil {
		t.Error(err)
	}
	zonesListJSON2, err := readFile("test-fixtures/zones/zones-list-2.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones?maxItems=1",
			code:     200,
			body:     zonesListJSON1,
		},
		{
			endpoint: "http://host.com/zones?startFrom=2&maxItems=1",
			code:     200,
			body:     zonesListJSON2,
		},
	})

	defer server.Close()

	if _, err := client.ZonesListAll(ListFilter{
		MaxItems: 200,
	}); err == nil {
		t.Error("Expected error -- MaxItems must be between 1 and 100")
	}

	zones, err := client.ZonesListAll(ListFilter{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(zones) != 2 {
		t.Error("Expected 2 Zones; got ", len(zones))
	}

	if zones[0].ID != "1" {
		t.Error("Expected Zone.ID to be 1")
	}

	if zones[1].ID != "2" {
		t.Error("Expected Zone.ID to be 2")
	}
}

func TestZonesListAllWhenNone(t *testing.T) {
	zonesListNoneJSON, err := readFile("test-fixtures/zones/zones-list-none.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones",
			code:     200,
			body:     zonesListNoneJSON,
		},
	})

	defer server.Close()

	zones, err := client.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(zones) != 0 {
		t.Error("Expected 0 Zones; got ", len(zones))
	}

	j, err := json.Marshal(zones)
	if err != nil {
		t.Error(err)
	}

	if string(j) != "[]" {
		t.Error("Expected string-converted marshaled JSON to be '[]'; got ", string(j))
	}
}

func TestZone(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123",
			code:     200,
			body:     zoneJSON,
		},
	})

	defer server.Close()

	z, err := client.Zone("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}

	if z.Name != "vinyldns." {
		t.Error("Expected zone.Name to have a value")
	}
	if z.Email != "some_user@foo.com" {
		t.Error("Expected zone.Email to have a value")
	}
	if z.Status != "Active" {
		t.Error("Expected zone.Status to have a value")
	}
	if z.Created != "2015-10-30T01:25:46Z" {
		t.Error("Expected zone.Created to have a value")
	}
	if z.ID != "123" {
		t.Error("Expected zone.ID to have a value")
	}
	if z.LatestSync == "" {
		t.Error("Expected zone.LatestSync to have a value")
	}
	if z.Updated == "" {
		t.Error("Expected zone.Updated to have a value")
	}
	if z.AdminGroupID == "" {
		t.Error("Expected zone.AdminGroupID to have a value")
	}
	if z.Account == "" {
		t.Error("Expected zone.Account to have a value")
	}
	if z.BackendID == "" {
		t.Error("Expected zone.BackendID to have a value")
	}
	if z.AccessLevel == "" {
		t.Error("Expected zone.AccessLevel to have a value")
	}
	if z.Connection.Name != "vinyldns." {
		t.Error("Expected zone.Connection.Name to have a value")
	}
	if z.Connection.KeyName != "vinyldns." {
		t.Error("Expected zone.Connection.KeyName to have a value")
	}
	if z.Connection.Key != "OBF:1:ABC" {
		t.Error("Expected zone.Connection.Key to have a value")
	}
	if z.Connection.PrimaryServer != "127.0.0.1" {
		t.Error("Expected zone.Connection.PrimaryServer to have a value")
	}
	if z.TransferConnection.Name != "vinyldns." {
		t.Error("Expected zone.TransferConnection.Name to have a value")
	}
	if z.TransferConnection.KeyName != "vinyldns." {
		t.Error("Expected zone.TransferConnection.KeyName to have a value")
	}
	if z.TransferConnection.Key != "OBF:1:ABC+5" {
		t.Error("Expected zone.TransferConnection.Key to have a value")
	}
	if z.TransferConnection.PrimaryServer != "127.0.0.1" {
		t.Error("Expected zone.TransferConnection.PrimaryServer to have a value")
	}

	rule := z.ACL.Rules[0]
	if rule.AccessLevel != "Read" {
		t.Error("Expected rule.AccessLevel to be Read")
	}
	if rule.Description != "test-acl-group-id" {
		t.Error("Expected rule.Description to be test-acl-group-id")
	}
	if rule.GroupID != "123" {
		t.Error("Expected rule.GroupId to be 123")
	}
	if rule.RecordMask != "www-*" {
		t.Error("Expected rule.RecordMask to be www-*")
	}
	for _, rt := range rule.RecordTypes {
		if rt != "A" && rt != "AAAA" && rt != "CNAME" {
			t.Error("Expected rule.RecordTypes to be A, AAAA, CNAME")
		}
	}
}

func TestZoneByName(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/name/vinyldns",
			code:     200,
			body:     zoneJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneByName("vinyldns")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}

	if z.Name != "vinyldns." {
		t.Error("Expected zone.Name to have a value")
	}
	if z.Email != "some_user@foo.com" {
		t.Error("Expected zone.Email to have a value")
	}
	if z.Status != "Active" {
		t.Error("Expected zone.Status to have a value")
	}
	if z.Created != "2015-10-30T01:25:46Z" {
		t.Error("Expected zone.Created to have a value")
	}
	if z.ID != "123" {
		t.Error("Expected zone.ID to have a value")
	}
	if z.LatestSync == "" {
		t.Error("Expected zone.LatestSync to have a value")
	}
	if z.Updated == "" {
		t.Error("Expected zone.Updated to have a value")
	}
	if z.AdminGroupID == "" {
		t.Error("Expected zone.AdminGroupID to have a value")
	}
	if z.Account == "" {
		t.Error("Expected zone.Account to have a value")
	}
	if z.BackendID == "" {
		t.Error("Expected zone.BackendID to have a value")
	}
	if z.AccessLevel == "" {
		t.Error("Expected zone.AccessLevel to have a value")
	}
	if z.Connection.Name != "vinyldns." {
		t.Error("Expected zone.Connection.Name to have a value")
	}
	if z.Connection.KeyName != "vinyldns." {
		t.Error("Expected zone.Connection.KeyName to have a value")
	}
	if z.Connection.Key != "OBF:1:ABC" {
		t.Error("Expected zone.Connection.Key to have a value")
	}
	if z.Connection.PrimaryServer != "127.0.0.1" {
		t.Error("Expected zone.Connection.PrimaryServer to have a value")
	}
	if z.TransferConnection.Name != "vinyldns." {
		t.Error("Expected zone.TransferConnection.Name to have a value")
	}
	if z.TransferConnection.KeyName != "vinyldns." {
		t.Error("Expected zone.TransferConnection.KeyName to have a value")
	}
	if z.TransferConnection.Key != "OBF:1:ABC+5" {
		t.Error("Expected zone.TransferConnection.Key to have a value")
	}
	if z.TransferConnection.PrimaryServer != "127.0.0.1" {
		t.Error("Expected zone.TransferConnection.PrimaryServer to have a value")
	}

	rule := z.ACL.Rules[0]
	if rule.AccessLevel != "Read" {
		t.Error("Expected rule.AccessLevel to be Read")
	}
	if rule.Description != "test-acl-group-id" {
		t.Error("Expected rule.Description to be test-acl-group-id")
	}
	if rule.GroupID != "123" {
		t.Error("Expected rule.GroupId to be 123")
	}
	if rule.RecordMask != "www-*" {
		t.Error("Expected rule.RecordMask to be www-*")
	}
	for _, rt := range rule.RecordTypes {
		if rt != "A" && rt != "AAAA" && rt != "CNAME" {
			t.Error("Expected rule.RecordTypes to be A, AAAA, CNAME")
		}
	}
}

func TestZoneCreate(t *testing.T) {
	zoneUpdateResponseJSON, err := readFile("test-fixtures/zones/zone-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones",
			code:     200,
			body:     zoneUpdateResponseJSON,
		},
	})

	defer server.Close()

	zone := &Zone{
		Name:         "test.",
		Email:        "email@email.com",
		AdminGroupID: "1234",
		Connection: &ZoneConnection{
			Name:          "connectionName",
			KeyName:       "keyName",
			Key:           "key",
			PrimaryServer: "1.2.3.4",
		},
	}

	z, err := client.ZoneCreate(zone)
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z.Zone.Name != "test." {
		t.Error("Expected zoneResponse.Zone.Name to have a value")
	}
	if z.UserID != "pclear" {
		t.Error("Expected zoneResponse.Zone.UserId to have a value")
	}
}

func TestZoneUpdate(t *testing.T) {
	zoneUpdateResponseJSON, err := readFile("test-fixtures/zones/zone-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123",
			code:     200,
			body:     zoneUpdateResponseJSON,
		},
	})

	defer server.Close()

	zone := &Zone{
		ID:           "123",
		Name:         "test.",
		Email:        "email@email.com",
		AdminGroupID: "123",
		Connection: &ZoneConnection{
			Name:          "connectionName",
			KeyName:       "keyName",
			Key:           "key",
			PrimaryServer: "1.2.3.4",
		},
	}

	z, err := client.ZoneUpdate(zone)
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z.Zone.Name != "test." {
		t.Error("Expected zoneResponse.Zone.Name to have a value")
	}
	if z.UserID != "pclear" {
		t.Error("Expected zoneResponse.Zone.UserId to have a value")
	}
}

func TestZoneDelete(t *testing.T) {
	zoneUpdateResponseJSON, err := readFile("test-fixtures/zones/zone-update.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123",
			code:     200,
			body:     zoneUpdateResponseJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneDelete("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z.Zone.Name != "test." {
		t.Error("Expected zoneResponse.Zone.Name to have a value")
	}
	if z.UserID != "pclear" {
		t.Error("Expected zoneResponse.Zone.UserId to have a value")
	}
}

func TestZoneExists_yes(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123",
			code:     200,
			body:     zoneJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneExists("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z != true {
		t.Error("Expected ZoneExists to be true")
	}
}

func TestZoneExists_no(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123",
			code:     404,
			body:     zoneJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneExists("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z != false {
		t.Error("Expected ZoneExists to be false")
	}
}

func TestZoneNameExists_yes(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/name/.ok",
			code:     200,
			body:     zoneJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneNameExists(".ok")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z != true {
		t.Error("Expected ZoneNameExists to be true")
	}
}

func TestZoneNameExists_no(t *testing.T) {
	zoneJSON, err := readFile("test-fixtures/zones/zone.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/name/.ok",
			code:     404,
			body:     zoneJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneNameExists(".ok")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z != false {
		t.Error("Expected ZoneNameExists to be false")
	}
}

func TestZoneChanges(t *testing.T) {
	zoneChangesJSON, err := readFile("test-fixtures/zones/zone-changes.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/changes",
			code:     200,
			body:     zoneChangesJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneChanges("123")
	zc := z.ZoneChanges[0]
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z.ZoneID != "123" {
		t.Error("Expected ZoneChanges.ZoneID to have a value")
	}
	if zc.UserID != "userId1" {
		t.Error("Expected ZoneChanges.ZoneChanges[0].UserID to have a value")
	}
	if zc.ChangeType != "Create" {
		t.Error("Expected ZoneChanges.ZoneChanges[0].ChangeType to have a value")
	}
	if zc.Status != "Complete" {
		t.Error("Expected ZoneChanges.ZoneChanges[0].Status to have a value")
	}
	if zc.ID != "change123" {
		t.Error("Expected ZoneChanges.ZoneChanges[0].ID to have a value")
	}
	if zc.Zone.Name != "vinyldnstest.sys.vinyldns.net." {
		t.Error("Expected ZoneChanges.ZoneChanges[0].Zone.Name to have a value")
	}
}

func TestZoneSync(t *testing.T) {
	zoneSyncJSON, err := readFile("test-fixtures/zones/zone-sync.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/sync",
			code:     200,
			body:     zoneSyncJSON,
		},
	})

	defer server.Close()
	z, err := client.ZoneSync("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z.Zone.ID != "123" {
		t.Error("Expected ZoneChange.Zone.ID to have a value")
	}
	if z.Status != "Pending" {
		t.Error("Expected ZoneChange.Status to have a value")
	}
	if z.ChangeType != "Sync" {
		t.Error("Expected ZoneChange.ChangeType to have a value")
	}
	if z.Zone.Status != "Syncing" {
		t.Error("Expected ZoneChange.Zone.Status to have a value")
	}
	if z.Zone.Name != "sync-test." {
		t.Error("Expected ZoneChange.Zone.Name to have a value")
	}
}

func TestZoneChangesListAll(t *testing.T) {
	zoneChangesListJSON1, err := readFile("test-fixtures/zones/zone-changes-list-1.json")
	if err != nil {
		t.Error(err)
	}
	zoneChangesListJSON2, err := readFile("test-fixtures/zones/zone-changes-list-2.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/changes?maxItems=1",
			code:     200,
			body:     zoneChangesListJSON1,
		},
		{
			endpoint: "http://host.com/zones/123/changes?startFrom=2&maxItems=1",
			code:     200,
			body:     zoneChangesListJSON2,
		},
	})

	defer server.Close()

	if _, err := client.ZoneChangesListAll("123", ListFilter{
		MaxItems: 200,
	}); err == nil {
		t.Error("Expected error -- MaxItems must be between 1 and 100")
	}

	changes, err := client.ZoneChangesListAll("123", ListFilter{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(changes) != 2 {
		t.Error("Expected 2 ZoneChanges; got ", len(changes))
	}

	if changes[0].ID != "1" {
		t.Error("Expected Zone.ID to be 1")
	}

	if changes[1].ID != "2" {
		t.Error("Expected Zone.ID to be 2")
	}
}

func TestZoneChange(t *testing.T) {
	zoneChangesJSON, err := readFile("test-fixtures/zones/zone-changes.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/zones/123/changes",
			code:     200,
			body:     zoneChangesJSON,
		},
	})

	defer server.Close()

	z, err := client.ZoneChange("123", "change123")
	if err != nil {
		t.Log(pretty.PrettyFormat(z))
		t.Error(err)
	}
	if z.UserID != "userId1" {
		t.Error("Expected ZoneChange.UserID to have a value")
	}
}
