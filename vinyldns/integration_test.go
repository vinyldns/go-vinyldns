//go:build integration
// +build integration

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
	"fmt"
	"testing"
	"time"
)

// client() assumes a VinylDNS is running on localhost:9000 with the default access and secret keys
// see `make start-api` for a Make task in starting VinylDNS
func client() *Client {
	return NewClient(ClientConfiguration{
		"okAccessKey",
		"okSecretKey",
		"http://localhost:9000",
		"go-vinyldns integration testing",
	})
}

func TestGroupCreateIntegration(t *testing.T) {
	c := client()
	users := []User{
		{
			UserName:  "ok",
			FirstName: "ok",
			LastName:  "ok",
			Email:     "test@test.com",
			ID:        "ok",
		},
	}
	gc, err := c.GroupCreate(&Group{
		Name:        "test-group",
		Description: "a test group",
		Email:       "test@vinyldns.com",
		Admins:      users,
		Members:     users,
	})
	if err != nil {
		t.Error(err)
	}

	gg, err := c.Group(gc.ID)
	if err != nil {
		t.Error(err)
	}

	if gg.ID != gc.ID {
		t.Error(err)
	}
}

func TestGroupsListAllIntegrationFilterForNonexistentName(t *testing.T) {
	c := client()
	zones, err := c.ZonesListAll(ListFilter{
		NameFilter: "foo",
	})
	if err != nil {
		t.Error(err)
	}

	if len(zones) > 0 {
		t.Error("Expected GroupsListAll for groups named 'foo' to yield no results")
	}
}

func TestGroupAdminsIntegration(t *testing.T) {
	c := client()
	groups, err := c.Groups()
	if err != nil {
		t.Error(err)
	}

	gID := groups[0].ID
	admins, err := c.GroupAdmins(gID)
	if err != nil {
		t.Error(err)
	}

	if admins[0].UserName != "ok" {
		t.Error(fmt.Sprintf("unable to get group admins for group %s", gID))
	}
}

func TestZoneNameExistsForNonexistentZoneIntegration(t *testing.T) {
	c := client()
	exists, err := c.ZoneNameExists("foo")
	if err != nil {
		t.Error(err)
	}
	if exists != false {
		t.Error(fmt.Sprintf("expected ZoneNameExists to return false; got %t", exists))
	}
}

func TestZoneCreateIntegration(t *testing.T) {
	c := client()
	groups, err := c.Groups()
	if err != nil {
		t.Error(err)
	}
	connection := &ZoneConnection{
		Name:          "ok.",
		KeyName:       "vinyldns.",
		Key:           "nzisn+4G2ldMn0q1CV3vsg==",
		PrimaryServer: "127.0.0.1:19001",
	}

	zone := &Zone{
		Name:               "ok.",
		Email:              "email@email.com",
		AdminGroupID:       groups[0].ID,
		Connection:         connection,
		TransferConnection: connection,
	}

	zc, err := c.ZoneCreate(zone)
	if err != nil {
		t.Error(err)
	}

	createdZoneID := zc.Zone.ID
	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		zg, err := c.Zone(createdZoneID)
		if err == nil && zg.ID != createdZoneID {
			t.Error(fmt.Sprintf("unable to get zone %s", createdZoneID))
		}
		if err == nil && zg.ID == createdZoneID {
			break
		}

		if i == (limit - 1) {
			fmt.Printf("%d retries reached in polling for zone %s", limit, createdZoneID)
			t.Error(err)
		}
	}
}

func TestZoneNameExistsForExistentZoneIntegration(t *testing.T) {
	c := client()
	exists, err := c.ZoneNameExists("ok.")
	if err != nil {
		t.Error(err)
	}
	if exists != true {
		t.Error(fmt.Sprintf("expected ZoneNameExists to return true; got %t", exists))
	}
}

func TestZoneByNameIntegration(t *testing.T) {
	c := client()
	z, err := c.ZoneByName("ok")
	if err != nil {
		t.Error(err)
	}

	if z.Name != "ok." {
		t.Error(fmt.Sprintf("unable to get ZoneByName %s", "ok."))
	}
}

func TestZonesListAllIntegrationFilterForNonexistentName(t *testing.T) {
	c := client()
	zones, err := c.ZonesListAll(ListFilter{
		NameFilter: "foo",
	})
	if err != nil {
		t.Error(err)
	}

	if len(zones) > 0 {
		t.Error("Expected ZonesListAll for zones named 'foo' to yield no results")
	}
}

func TestZoneChangesIntegration(t *testing.T) {
	c := client()
	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	changes, err := c.ZoneChanges(zones[0].ID)
	if err != nil {
		t.Error(err)
	}

	if changes.ZoneID != zones[0].ID {
		t.Error("Expected ZoneChanges to yield correct ID")
	}

	if len(changes.ZoneChanges) <= 0 {
		t.Error("Expected ZoneChanges to yield results")
	}
}

func TestZoneChangesListAllIntegration(t *testing.T) {
	c := client()
	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	changes, err := c.ZoneChangesListAll(zones[0].ID, ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if changes[0].Zone.ID != zones[0].ID {
		t.Error("Expected ZoneChangesListAll to yield correct ID")
	}

	if len(changes) <= 0 {
		t.Error("Expected ZoneChangesListAll to yield results")
	}
}

func TestRecordSetCreateIntegrationARecord(t *testing.T) {
	c := client()
	zs, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}
	rc, err := c.RecordSetCreate(&RecordSet{
		Name:   "integration-test",
		ZoneID: zs[0].ID,
		Type:   "A",
		TTL:    60,
		Records: []Record{
			{
				Address: "127.0.0.1",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	createdID := rc.RecordSet.ID
	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		rg, err := c.RecordSet(zs[0].ID, createdID)
		if err == nil && rg.ID != createdID {
			t.Error(fmt.Sprintf("unable to get record set %s", createdID))
		}
		if err == nil && rg.ID == createdID {
			break
		}

		if i == (limit - 1) {
			fmt.Printf("%d retries reached in polling for record set %s", limit, createdID)
			t.Error(err)
		}
	}
}

func TestRecordSetCreateIntegrationNSRecord(t *testing.T) {
	c := client()
	zs, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}
	rc, err := c.RecordSetCreate(&RecordSet{
		Name:   "integration-test",
		ZoneID: zs[0].ID,
		Type:   "NS",
		TTL:    60,
		Records: []Record{
			{
				NSDName: "ns1.parent.com.",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	createdID := rc.RecordSet.ID
	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		rg, err := c.RecordSet(zs[0].ID, createdID)
		if err == nil && rg.ID != createdID {
			t.Error(fmt.Sprintf("unable to get record set %s", createdID))
		}
		if err == nil && rg.ID == createdID {
			break
		}

		if i == (limit - 1) {
			fmt.Printf("%d retries reached in polling for record set %s", limit, createdID)
			t.Error(err)
		}
	}
}

func TestRecordSetUpdateIntegrationARecord(t *testing.T) {
	c := client()
	zs, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}
	rs := &RecordSet{
		Name:   "integration-test-record-set-update",
		ZoneID: zs[0].ID,
		Type:   "A",
		TTL:    60,
		Records: []Record{
			{
				Address: "127.0.0.1",
			},
		},
	}
	rc, err := c.RecordSetCreate(rs)
	if err != nil {
		t.Error(err)
	}
	createdID := rc.RecordSet.ID
	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		rg, err := c.RecordSet(zs[0].ID, createdID)
		if err == nil && rg.ID != createdID {
			t.Error(fmt.Sprintf("unable to get record set %s", createdID))
		}
		if err == nil && rg.ID == createdID {
			updatedName := "updated-integration-test-record-set-update"
			rs.ID = createdID
			rs.Name = updatedName
			u, err := c.RecordSetUpdate(rs)
			if err == nil && u.RecordSet.ID != createdID && u.RecordSet.Name != updatedName {
				t.Error(fmt.Sprintf("unable to updated record set %s", createdID))
			}
			break
		}

		if i == (limit - 1) {
			fmt.Printf("%d retries reached in polling for record set %s", limit, createdID)
			t.Error(err)
		}
	}
}

func TestRecordSetsListAllIntegrationFilterForExistentName(t *testing.T) {
	c := client()
	zs, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	records, err := c.RecordSetsListAll(zs[0].ID, ListFilter{
		NameFilter: "foo",
	})
	if err != nil {
		t.Error(err)
	}

	if len(records) < 1 {
		t.Error("Expected RecordSetsListAll for records named 'foo' to yield results")
	}
}

func TestRecordSetsListAllIntegrationFilterForNonexistentName(t *testing.T) {
	c := client()
	zs, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	records, err := c.RecordSetsListAll(zs[0].ID, ListFilter{
		NameFilter: "thisdoesnotexist",
	})
	if err != nil {
		t.Error(err)
	}

	if len(records) > 0 {
		t.Error("Expected RecordSetsListAll for records named 'thisdoesnotexist' to yield no results")
	}
}

func TestRecordSetsGlobalListAllIntegrationFilterForExistentName(t *testing.T) {
	c := client()
	rName := "foo"

	records, err := c.RecordSetsGlobalListAll(GlobalListFilter{
		RecordNameFilter: "*" + rName + "*",
	})
	if err != nil {
		t.Error(err)
	}

	if len(records) < 1 {
		t.Error(fmt.Sprintf("Expected RecordSetsGlobalListAll for records named '%s' to yield results", rName))
	}

	if records[0].Name != rName {
		t.Error(fmt.Sprintf("Expected RecordSetsGlobalListAll for records named '%s' to return the matching record", rName))
	}
}

func TestRecordSetsGlobalListAllIntegrationFilterForNonexistentName(t *testing.T) {
	c := client()
	records, err := c.RecordSetsGlobalListAll(GlobalListFilter{
		RecordNameFilter: "thisdoesnotexist",
	})
	if err != nil {
		t.Error(err)
	}

	if len(records) > 0 {
		t.Error("Expected RecordSetsListAll for records named 'thisdoesnotexist' to yield no results")
	}
}

func TestRecordSetDeleteIntegration(t *testing.T) {
	c := client()
	zs, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}
	z := zs[0].ID

	rs, err := c.RecordSetsListAll(z, ListFilter{})
	if err != nil {
		t.Error(err)
	}

	var r string
	for _, rec := range rs {
		if rec.Name == "integration-test" {
			r = rec.ID
			break
		}
	}

	_, err = c.RecordSetDelete(z, r)
	if err != nil {
		t.Error(err)
	}
}

func TestRecordSetChangesIntegration(t *testing.T) {
	c := client()
	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	changes, err := c.RecordSetChanges(zones[0].ID, ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(changes.RecordSetChanges) <= 0 {
		t.Error("Expected RecordSetChanges to return results")
	}
}

func TestRecordSetChangesIntegrationWithMaxItems(t *testing.T) {
	c := client()
	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	changes, err := c.RecordSetChanges(zones[0].ID, ListFilter{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(changes.RecordSetChanges) > 1 {
		t.Error("Expected RecordSetChanges to return only 1 results")
	}
}

func TestRecordSetChangesListAllIntegration(t *testing.T) {
	c := client()
	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	changes, err := c.RecordSetChangesListAll(zones[0].ID, ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(changes) <= 0 {
		t.Error("Expected RecordSetChangesListAll to yield results")
	}
}

func TestZoneDeleteIntegration(t *testing.T) {
	c := client()
	zs, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}
	z := zs[0].ID

	_, err = c.ZoneDelete(z)
	if err != nil {
		t.Error(err)
	}

	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		exists, err := c.ZoneExists(z)
		if err != nil {
			t.Error(err)
			break
		}

		if !exists {
			break
		}

		if i == (limit - 1) {
			fmt.Printf("%d retries reached in waiting for zone deletion of %s", limit, z)
			t.Error(err)
		}
	}
}

func TestGroupDeleteIntegration(t *testing.T) {
	c := client()
	gs, err := c.Groups()
	if err != nil {
		t.Error(err)
	}
	g := gs[0].ID
	_, err = c.GroupDelete(g)
	if err != nil {
		t.Error(err)
	}
}
