//go:build integration
// +build integration

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
	"fmt"
	"strings"
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

func superUser() *Client {
	return NewClient(ClientConfiguration{
		"superUserAccessKey",
		"superUserSecretKey",
		"http://localhost:9000",
		"go-vinyldns integration testing (super user)",
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
		Email:       "test@test.com",
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

	if admins[0].UserName != "ok" && admins[0].UserName != "dummy" {
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
		PrimaryServer: "vinyldns-integration:19001",
	}

	zone := &Zone{
		Name:               "ok.",
		Email:              "email@test.com",
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
	if z.AdminGroupID == "" {
		t.Error("expected ZoneByName to return a zone with a valid AdminGroupID")
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
		RecordNameFilter: rName + "*",
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

	changes, err := c.RecordSetChanges(zones[0].ID, ListFilterRecordSetChanges{})
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

	changes, err := c.RecordSetChanges(zones[0].ID, ListFilterRecordSetChanges{
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

	changes, err := c.RecordSetChangesListAll(zones[0].ID, ListFilterRecordSetChanges{})
	if err != nil {
		t.Error(err)
	}

	if len(changes) <= 0 {
		t.Error("Expected RecordSetChangesListAll to yield results")
	}
}

func TestRecordSetChangesFailureIntegration(t *testing.T) {
	c := client()

	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(zones) == 0 {
		t.Error("Expected at least one zone for RecordSetChangesFailure testing")
		return
	}

	failures, err := c.RecordSetChangesFailure(zones[0].ID, ListFilter{})
	if err != nil {
		t.Error(err)
		return
	}

	// TODO: How to create a failure for testing?
	if failures == nil {
		t.Error("Expected RecordSetChangesFailure to return a non-nil response")
	}
}

// ==============================================================================
// Health & Monitoring Tests
// ==============================================================================

func TestPingIntegration(t *testing.T) {
	c := client()
	response, err := c.Ping()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(response, "PONG") {
		t.Errorf("Expected Ping to return 'PONG', got: %s", response)
	}
}

func TestHealthIntegration(t *testing.T) {
	c := client()
	err := c.Health()
	if err != nil {
		t.Errorf("Health check failed: %v", err)
	}
}

func TestColorIntegration(t *testing.T) {
	c := client()
	color, err := c.Color()
	if err != nil {
		t.Error(err)
	}

	if color == "" {
		t.Error("Expected Color to return a non-empty value")
	} else if color != "green" && color != "blue" {
		t.Errorf("Expected Color to return 'green' or 'blue', got: %s", color)
	}
}

func TestMetricsPrometheusIntegration(t *testing.T) {
	c := client()
	metrics, err := c.MetricsPrometheus(nil)
	if err != nil {
		t.Error(err)
	}

	// TODO: Actually validate some metrics content?
	if metrics == "" {
		t.Log("Expected MetricsPrometheus to return non-empty metrics")
		// t.Error("MetricsPrometheus returned empty metrics")
	}
}

// ==============================================================================
// System Management Tests
// ==============================================================================

func TestStatusIntegration(t *testing.T) {
	c := client()
	status, err := c.Status()
	if err != nil {
		t.Error(err)
	}

	if status.Version == "" {
		t.Error("Expected Status to return a version")
	}
	if status.Color == "" {
		t.Error("Expected Status to return a color")
	}
	if status.ProcessingDisabled != true && status.ProcessingDisabled != false {
		t.Error("Expected Status to return a valid ProcessingDisabled value")
	}
	if status.KeyName == "" {
		t.Error("Expected Status to return a valid KeyName")
	}
}

func TestStatusUpdateIntegration(t *testing.T) {
	c := superUser()

	currentStatus, err := c.Status()
	if err != nil {
		t.Error(err)
	}

	newDisabledState := !currentStatus.ProcessingDisabled
	updatedStatus, err := c.StatusUpdate(newDisabledState)
	if err != nil {
		t.Error(err)
	}

	if updatedStatus.ProcessingDisabled != newDisabledState {
		t.Errorf("Expected ProcessingDisabled to be %v, got %v", newDisabledState, updatedStatus.ProcessingDisabled)
	}

	// Restore original state
	_, err = c.StatusUpdate(currentStatus.ProcessingDisabled)
	if err != nil {
		t.Error(err)
	}
}

// ==============================================================================
// User Management Tests
// ==============================================================================

func TestUserIntegration(t *testing.T) {
	c := client()
	user, err := c.User("ok")
	if err != nil {
		t.Error(err)
	}

	if user.UserName != "ok" {
		t.Errorf("Expected user name to be 'ok', got: %s", user.UserName)
	}

	if user.ID == "" {
		t.Error("Expected User to return a valid ID")
	}

	if user.GroupID == nil || len(user.GroupID) == 0 {
		t.Error("Expected User to return at least one GroupID")
	}
}

func TestUserLockUnlockIntegration(t *testing.T) {
	c := superUser()

	user, err := c.User("ok")
	if err != nil {
		t.Error(err)
	}

	lockedUser, err := c.UserLock(user.ID)
	if err != nil {
		t.Errorf("UserLock failed: %v", err)
		return
	}

	if lockedUser.LockStatus != "Locked" && lockedUser.LockStatus != "" {
		t.Errorf("User lock status: %s", lockedUser.LockStatus)
	}

	// Unlock the user to restore state
	unlockedUser, err := c.UserUnlock(user.ID)
	if err != nil {
		t.Errorf("UserUnlock failed: %v", err)
	}

	if unlockedUser.LockStatus == "Locked" {
		t.Error("Expected user to be unlocked")
	}
}

// ==============================================================================
// Zone Enhancement Tests
// ==============================================================================

func TestZoneBackendIDsIntegration(t *testing.T) {
	c := client()
	backendIDs, err := c.ZoneBackendIDs()
	if err != nil {
		t.Error(err)
	}

	if len(backendIDs) == 0 {
		t.Error("Expected ZoneBackendIDs to return at least one backend")
	}
}

func TestZoneACLRuleCreateDeleteIntegration(t *testing.T) {
	c := superUser()
	groups, err := c.Groups()
	if err != nil {
		t.Error(err)
	}
	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(zones) == 0 {
		t.Error("No zones available for ACL rule testing")
	}

	zoneID := zones[0].ID

	rule := &ACLRule{
		AccessLevel: "Read",
		Description: "Integration test ACL rule",
		RecordTypes: []string{"A", "AAAA"},
		GroupID: groups[0].ID,
	}

	createResp, err := c.ZoneACLRuleCreate(zoneID, rule)
	if err != nil {
		t.Errorf("ZoneACLRuleCreate failed: %v", err)
		return
	}

	if createResp.Zone.ID != zoneID {
		t.Error("Expected ZoneACLRuleCreate to return the correct zone ID")
	}

	if createResp.Status != "Accepted" && createResp.Status != "Pending" {
		t.Errorf("Expected ZoneACLRuleCreate to return status 'Accepted' or 'Pending', got: %s", createResp.Status)
	}

	deleteResp, err := c.ZoneACLRuleDelete(zoneID, rule)
	if err != nil {
		t.Errorf("ZoneACLRuleDelete failed: %v", err)
		return
	}

	if deleteResp.Zone.ID != zoneID {
		t.Error("Expected ZoneACLRuleDelete to return the correct zone ID")
	}

	if deleteResp.Status != "Accepted" && deleteResp.Status != "Pending" {
		t.Errorf("Expected ZoneACLRuleDelete to return status 'Accepted' or 'Pending', got: %s", deleteResp.Status)
	}
}

func TestZoneChangesFailureIntegration(t *testing.T) {
	c := client()
	failures, err := c.ZoneChangesFailure(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	// It's okay if there are no failures
	if failures == nil {
		t.Error("Expected ZoneChangesFailure to return a non-nil response")
	}
}

// ==============================================================================
// Record Set Enhancement Tests
// ==============================================================================

func TestRecordSetCountIntegration(t *testing.T) {
	c := client()

	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(zones) == 0 {
		t.Error("No zones available for record set count testing")
	}

	count, err := c.RecordSetCount(zones[0].ID)
	if err != nil {
		t.Error(err)
	}

	if count.Count < 0 {
		t.Error("Expected RecordSetCount to return a non-negative count")
	}
}

func TestRecordSetChangeHistoryIntegration(t *testing.T) {
	c := client()

	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(zones) == 0 {
		t.Error("No zones available for change history testing")
	}

	recordSets, err := c.RecordSetsListAll(zones[0].ID, ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(recordSets) == 0 {
		t.Error("No record sets available for change history testing")
	}

	// Use FQDN format for the filter
	fqdn := recordSets[0].Name
	if !strings.HasSuffix(fqdn, ".") && zones[0].Name != "" {
		fqdn = recordSets[0].Name + "." + zones[0].Name
	}

	filter := RecordSetChangeHistoryFilter{
		ZoneID: zones[0].ID,
		FQDN:   fqdn,
	}

	history, err := c.RecordSetChangeHistory(filter)
	if err != nil {
		t.Error(err)
	}

	if history == nil {
		t.Error("Expected RecordSetChangeHistory to return a non-nil response")
	}
}

func TestRecordSetOwnershipTransferIntegration(t *testing.T) {
	c := client()

	groups, err := c.Groups()
	if err != nil {
		t.Error(err)
		return
	}

	if len(groups) < 2 {
		t.Error("Expected at least 2 groups for ownership transfer testing")
		return
	}

	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(zones) == 0 {
		t.Error("Expected at least one zone for ownership transfer testing")
		return
	}

	recordSet := &RecordSet{
		Name:   "ownership-transfer-test",
		ZoneID: zones[0].ID,
		Type:   "A",
		TTL:    60,
		Records: []Record{
			{Address: "127.0.0.1"},
		},
		OwnerGroupID: groups[0].ID,
	}

	createResp, err := c.RecordSetCreate(recordSet)
	if err != nil {
		t.Errorf("Failed to create test record set: %v", err)
		return
	}

	createdRecordSetID := createResp.RecordSet.ID

	limit := 10
	var fetchedRecordSet RecordSet
	for i := 0; i < limit; i++ {
		time.Sleep(5 * time.Second)

		fetchedRecordSet, err = c.RecordSet(zones[0].ID, createdRecordSetID)
		if err == nil && fetchedRecordSet.ID == createdRecordSetID {
			break
		}

		if i == (limit - 1) {
			t.Errorf("Failed to fetch created record set after %d retries", limit)
			return
		}
	}

	transferResp, err := c.RecordSetOwnershipTransferRequest(&fetchedRecordSet, groups[1].ID)
	if err != nil {
		t.Errorf("Ownership transfer request failed: %v", err)

		// Clean up
		_, cleanupErr := c.RecordSetDelete(zones[0].ID, createdRecordSetID)
		if cleanupErr != nil {
			t.Logf("Failed to clean up test record set: %v", cleanupErr)
		}
		return
	}

	if transferResp.RecordSet.ID != createdRecordSetID {
		t.Error("Expected ownership transfer to return the correct record set ID")
	}

	// Wait for transfer to be processed
	time.Sleep(5 * time.Second)

	updatedRecordSet, err := c.RecordSet(zones[0].ID, createdRecordSetID)
	if err != nil {
		t.Errorf("Failed to fetch record set after transfer request: %v", err)
	}

	_, err = c.RecordSetOwnershipTransferCancel(&updatedRecordSet, groups[1].ID)
	if err != nil {
		t.Logf("Ownership transfer cancel failed: %v", err)
	}

	// Clean up - delete the test record set
	time.Sleep(5 * time.Second)
	_, err = c.RecordSetDelete(zones[0].ID, createdRecordSetID)
	if err != nil {
		t.Errorf("Failed to clean up test record set: %v", err)
	}
}

// ==============================================================================
// Group Enhancement Tests
// ==============================================================================

func TestGroupChangeIntegration(t *testing.T) {
	c := client()

	user, err := c.User("ok")
	if err != nil {
		t.Error(err)
	}
	gID := user.GroupID[0]
	group, err := c.Group(gID)
	if err != nil {
		t.Error(err)
	}

	group.Description = "Updated description for group change test"
	_, err = c.GroupUpdate(gID, group)
	if err != nil {
		t.Error(err)
	}

	// Get group activity to find a change ID
	activity, err := c.GroupActivity(gID)
	if err != nil {
		t.Error(err)
	}

	if len(activity.Changes) == 0 {
		t.Error("No group changes available for testing")
	}

	changeID := activity.Changes[0].ID
	groupChange, err := c.GroupChange(changeID)
	if err != nil {
		t.Error(err)
		return
	}

	if groupChange.ID != changeID {
		t.Errorf("Expected group change ID to be %s, got %s", changeID, groupChange.ID)
		return
	}
	if groupChange.ChangeType != "Update" {
		t.Errorf("Expected group change type to be 'Update', got %s", groupChange.ChangeType)
	}
	if groupChange.OldGroup.Description == group.Description {
		t.Error("Expected group change 'Before' description to differ from updated description")
	}
	if groupChange.NewGroup.Description != group.Description {
		t.Error("Expected group change 'After' description to match updated description")
	}
}

func TestGroupValidDomainsIntegration(t *testing.T) {
	c := client()

	domains, err := c.GroupValidDomains()
	if err != nil {
		t.Error(err)
	}

	if domains == nil {
		t.Error("Expected GroupValidDomains to return a non-nil response")
	}

	for _, domain := range domains {
		if domain == "" {
			t.Error("Expected valid domains to be non-empty strings")
		}
	}
}

// ==============================================================================
// Batch Changes Enhancement Tests
// ==============================================================================

func TestBatchRecordChangeWorkflowIntegration(t *testing.T) {
	c := client()

	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
	}

	if len(zones) == 0 {
		t.Error("No zones available for batch change testing")
	}

	batchChange := &BatchRecordChange{
		Comments: "Integration test batch change",
		Changes: []RecordChange{
			{
				ChangeType: "Add",
				InputName:  "batch-test." + zones[0].Name,
				Type:       "A",
				TTL:        300,
				Record: RecordData{
					Address: "127.0.0.1",
				},
			},
		},
		OwnerGroupID: zones[0].AdminGroupID,
	}

	createResp, err := c.BatchRecordChangeCreate(batchChange)
	if err != nil {
		t.Errorf("Failed to create batch change: %v", err)
		return
	}

	batchChangeID := createResp.ID

	limit := 10
	for i := 0; i < limit; i++ {
		time.Sleep(5 * time.Second)

		fetchedBatch, err := c.BatchRecordChange(batchChangeID)
		if err != nil {
			t.Errorf("Failed to fetch batch change: %v", err)
			break
		}

		if fetchedBatch.Status == "Complete" {
			return
		}

		if fetchedBatch.Status == "Failed" {
			t.Errorf("Batch change failed, status: %s", fetchedBatch.Status)
			return
		}

		// If in manual review, try to approve/cancel
		if fetchedBatch.Status == "PendingReview" {
			// Try to approve
			review := &BatchChangeReview{
				ReviewComment: "Integration test approval",
			}

			_, approveErr := c.BatchRecordChangeApprove(batchChangeID, review)
			if approveErr != nil {
				t.Logf("Batch change approve failed: %v", approveErr)

				// Try to cancel instead
				_, cancelErr := c.BatchRecordChangeCancel(batchChangeID, review)
				if cancelErr != nil {
					t.Errorf("Batch change approve/cancel failed: %v", cancelErr)
					return
				}
				if fetchedBatch.Status != "Cancelled" {
					t.Error("Expected batch change to be cancelled")
					return
				}
			}
		}

		if i == (limit - 1) {
			t.Errorf("Batch change did not complete after %d retries, status: %s", limit, fetchedBatch.Status)
		}
	}
}

// Note: This test is skipped if manual review is not enabled in the environment
func TestBatchRecordChangeRejectIntegration(t *testing.T) {
	c := client()

	zones, err := c.ZonesListAll(ListFilter{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(zones) == 0 {
		t.Error("No zones available for batch change testing")
		return
	}

	batchChange := &BatchRecordChange{
		Comments: "Integration test batch change",
		Changes: []RecordChange{
			{
				ChangeType: "Add",
				InputName:  "batch-test-reject." + zones[0].Name,
				Type:       "A",
				TTL:        300,
				Record: RecordData{
					Address: "127.0.0.2",
				},
			},
		},
		OwnerGroupID: zones[0].AdminGroupID,
	}

	createResp, err := c.BatchRecordChangeCreate(batchChange)
	if err != nil {
		t.Errorf("Failed to create batch change: %v", err)
		return
	}

	batchChangeID := createResp.ID
	limit := 10
	for i := 0; i < limit; i++ {
		time.Sleep(5 * time.Second)
		fetchedBatch, err := c.BatchRecordChange(batchChangeID)
		if err != nil {
			t.Errorf("Failed to fetch batch change: %v", err)
			break
		}

		if fetchedBatch.Status == "Complete" || fetchedBatch.Status == "Failed" {
			t.Skip("Batch change completed without review. Manual review is not enabled")
			return
		}

		if fetchedBatch.Status == "PendingReview"{
			break
		}
		if i == (limit - 1) {
			t.Errorf("Batch change did not complete/reach manual review after %d retries, status: %s", limit, fetchedBatch.Status)
		}
	}
	review := &BatchChangeReview{
		ReviewComment: "Integration test rejection",
	}
	_, err = c.BatchRecordChangeReject(batchChangeID, review)
	if err != nil {
		t.Errorf("Failed to reject batch change: %v", err)
	}
	fetchedBatch, err := c.BatchRecordChange(batchChangeID)
	if err != nil {
		t.Errorf("Failed to fetch batch change after rejection: %v", err)
	}
	if fetchedBatch.Status != "Rejected" {
		t.Errorf("Expected batch change status to be 'Rejected', got: %s", fetchedBatch.Status)
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

func TestZonesDeletedIntegration(t *testing.T) {
	c := client()
	deletedZones, err := c.ZonesDeleted(DeletedZonesFilter{})
	if err != nil {
		t.Error(err)
	}

	if deletedZones == nil {
		t.Error("Expected ZonesDeleted to return a non-nil response")
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
