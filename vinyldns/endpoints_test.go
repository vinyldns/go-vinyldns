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
	"net/http"
	"testing"
)

var c = &Client{
	"accessKey",
	"secretKey",
	"http://host.com",
	&http.Client{},
	"go-vinyldns testing",
}

func TestZonesEP(t *testing.T) {
	zones := zonesEP(c)
	expected := "http://host.com/zones"

	if zones != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zones)
		t.Error("zonesEP should return the right endpoint")
	}
}

func TestZonesListEP(t *testing.T) {
	zones := zonesListEP(c, ListFilter{
		NameFilter: "foo",
		MaxItems:   2,
		StartFrom:  "123",
	})
	expected := "http://host.com/zones?nameFilter=foo&startFrom=123&maxItems=2"

	if zones != expected {
		fmt.Printf("\nExpected: %s", expected)
		fmt.Printf("\nActual: %s", zones)
		t.Error("zonesListEP should return the right endpoint")
	}
}

func TestZonesListEPWithoutAllFilterParams(t *testing.T) {
	zones := zonesListEP(c, ListFilter{
		NameFilter: "foo",
	})
	expected := "http://host.com/zones?nameFilter=foo"

	if zones != expected {
		fmt.Printf("\nExpected: %s", expected)
		fmt.Printf("\nActual: %s", zones)
		t.Error("zonesListEP should return the right endpoint")
	}
}

func TestZonesListEPWithoutAnyFilterParams(t *testing.T) {
	zones := zonesListEP(c, ListFilter{})
	expected := "http://host.com/zones"

	if zones != expected {
		fmt.Printf("\nExpected: %s", expected)
		fmt.Printf("\nActual: %s", zones)
		t.Error("zonesListEP should return the right endpoint")
	}
}

func TestZoneEP(t *testing.T) {
	zone := zoneEP(c, "123")
	expected := "http://host.com/zones/123"

	if zone != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zone)
		t.Error("zoneEP should return the right endpoint")
	}
}
func TestZoneDetailsEP(t *testing.T) {
	zone := zoneDetailsEP(c, "123")
	expected := "http://host.com/zones/123/details"

	if zone != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zone)
		t.Error("zoneDetailsEP should return the right endpoint")
	}
}

func TestZoneSyncEP(t *testing.T) {
	zone := zoneSyncEP(c, "123")
	expected := "http://host.com/zones/123/sync"

	if zone != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zone)
		t.Error("zoneSyncEP should return the right endpoint")
	}
}

func TestZoneNameEP(t *testing.T) {
	zone := zoneNameEP(c, "foo")
	expected := "http://host.com/zones/name/foo"

	if zone != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zone)
		t.Error("zoneNameEP should return the right endpoint")
	}
}

func TestZoneChangesEP(t *testing.T) {
	zone := zoneChangesEP(c, "123", ListFilter{})
	expected := "http://host.com/zones/123/changes"

	if zone != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zone)
		t.Error("zoneChangesEP should return the right endpoint")
	}
}

func TestRecordSetsListEP(t *testing.T) {
	rs := recordSetsListEP(c, "123", ListFilter{})
	expected := "http://host.com/zones/123/recordsets"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetsListEP should return the right endpoint")
	}

	rs = recordSetsListEP(c, "543", ListFilter{
		StartFrom: "nextplease",
	})
	expected = "http://host.com/zones/543/recordsets?startFrom=nextplease"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetsListEP should return the right endpoint")
	}

	rs = recordSetsListEP(c, "7", ListFilter{
		StartFrom: "nextplease",
		MaxItems:  99,
	})
	expected = "http://host.com/zones/7/recordsets?startFrom=nextplease&maxItems=99"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetsListEP should return the right endpoint")
	}

	rs = recordSetsListEP(c, "7", ListFilter{
		NameFilter: "foo",
		StartFrom:  "nextplease",
		MaxItems:   99,
	})
	expected = "http://host.com/zones/7/recordsets?recordNameFilter=foo&startFrom=nextplease&maxItems=99"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetsListEP should return the right endpoint")
	}
}

func TestRecordSetsGlobalListEP(t *testing.T) {
	rs := recordSetsGlobalListEP(c, GlobalListFilter{})
	expected := "http://host.com/recordsets"
	msg := "recordSetsGlobalListEP should return the right endpoint"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error(msg)
	}

	rs = recordSetsGlobalListEP(c, GlobalListFilter{
		StartFrom: "nextplease",
	})
	expected = "http://host.com/recordsets?startFrom=nextplease"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error(msg)
	}

	rs = recordSetsGlobalListEP(c, GlobalListFilter{
		StartFrom: "nextplease",
		MaxItems:  99,
	})
	expected = "http://host.com/recordsets?startFrom=nextplease&maxItems=99"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error(msg)
	}

	rs = recordSetsGlobalListEP(c, GlobalListFilter{
		RecordNameFilter: "foo",
		StartFrom:        "nextplease",
		MaxItems:         99,
	})
	expected = "http://host.com/recordsets?recordNameFilter=foo&startFrom=nextplease&maxItems=99"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error(msg)
	}
}

func TestRecordSetEP(t *testing.T) {
	rs := recordSetEP(c, "123", "456")
	expected := "http://host.com/zones/123/recordsets/456"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetEP should return the right endpoint")
	}
}

func TestRecordSetChangesEP(t *testing.T) {
	rsc := recordSetChangesEP(c, "123", ListFilterRecordSetChanges{})
	expected := "http://host.com/zones/123/recordsetchanges"

	if rsc != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rsc)
		t.Error("recordSetChangesEP should return the right endpoint")
	}
}

func TestRecordSetChangesEPWithQuery(t *testing.T) {
	rsc := recordSetChangesEP(c, "123", ListFilterRecordSetChanges{
		MaxItems:  3,
		StartFrom: 1,
	})
	expected := "http://host.com/zones/123/recordsetchanges?startFrom=1&maxItems=3"

	if rsc != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rsc)
		t.Error("recordSetChangesEP should return the right endpoint")
	}
}

func TestRecordSetChangeEP(t *testing.T) {
	rsc := recordSetChangeEP(c, "123", "456", "789")
	expected := "http://host.com/zones/123/recordsets/456/changes/789"

	if rsc != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rsc)
		t.Error("recordSetChangeEP should return the right endpoint")
	}
}

func TestGroupsEP(t *testing.T) {
	rs := groupsEP(c)
	expected := "http://host.com/groups"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("groupsEP should return the right endpoint")
	}
}

func TestGroupEp(t *testing.T) {
	rs := groupEP(c, "123")
	expected := "http://host.com/groups/123"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("groupEp should return the right endpoint")
	}
}

func TestGroupsListEP(t *testing.T) {
	groups := groupsListEP(c, ListFilter{
		NameFilter: "foo",
		MaxItems:   2,
		StartFrom:  "123",
	})
	expected := "http://host.com/groups?groupNameFilter=foo&startFrom=123&maxItems=2"

	if groups != expected {
		fmt.Printf("\nExpected: %s", expected)
		fmt.Printf("\nActual: %s", groups)
		t.Error("groupsListEP should return the right endpoint")
	}
}

func TestGroupsListEPWithoutAllFilterParams(t *testing.T) {
	groups := groupsListEP(c, ListFilter{
		NameFilter: "foo",
	})
	expected := "http://host.com/groups?groupNameFilter=foo"

	if groups != expected {
		fmt.Printf("\nExpected: %s", expected)
		fmt.Printf("\nActual: %s", groups)
		t.Error("groupsListEP should return the right endpoint")
	}
}

func TestGroupsListEPWithoutAnyFilterParams(t *testing.T) {
	groups := groupsListEP(c, ListFilter{})
	expected := "http://host.com/groups"

	if groups != expected {
		fmt.Printf("\nExpected: %s", expected)
		fmt.Printf("\nActual: %s", groups)
		t.Error("groupsListEP should return the right endpoint")
	}
}

func TestGroupAdminsEp(t *testing.T) {
	ga := groupAdminsEP(c, "123")
	expected := "http://host.com/groups/123/admins"

	if ga != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", ga)
		t.Error("groupAdminsEp should return the right endpoint")
	}
}

func TestGroupMembersEp(t *testing.T) {
	gm := groupMembersEP(c, "123")
	expected := "http://host.com/groups/123/members"

	if gm != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", gm)
		t.Error("groupMembersEp should return the right endpoint")
	}
}

func TestGroupActivityEp(t *testing.T) {
	ga := groupActivityEP(c, "123")
	expected := "http://host.com/groups/123/activity"

	if ga != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", ga)
		t.Error("groupActivityEp should return the right endpoint")
	}
}

func TestBatchRecordChangesEP(t *testing.T) {
	brc := batchRecordChangesEP(c)
	expected := "http://host.com/zones/batchrecordchanges"

	if brc != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", brc)
		t.Error("batchRecordChangesEP should return the right endpoint")
	}
}

func TestBatchRecordChangeEP(t *testing.T) {
	brc := batchRecordChangeEP(c, "123")
	expected := "http://host.com/zones/batchrecordchanges/123"

	if brc != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", brc)
		t.Error("batchRecordChangeEP should return the right endpoint")
	}
}

func TestHealthEndpoints(t *testing.T) {
	ping := pingEP(c)
	expected := "http://host.com/ping"
	if ping != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", ping)
		t.Error("pingEP should return the right endpoint")
	}

	health := healthEP(c)
	expected = "http://host.com/health"
	if health != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", health)
		t.Error("healthEP should return the right endpoint")
	}

	color := colorEP(c)
	expected = "http://host.com/color"
	if color != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", color)
		t.Error("colorEP should return the right endpoint")
	}

	metrics := prometheusMetricsEP(c, []string{"foo", "bar"})
	expected = "http://host.com/metrics/prometheus?name=foo&name=bar"
	if metrics != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", metrics)
		t.Error("prometheusMetricsEP should return the right endpoint")
	}

	status := statusEP(c)
	expected = "http://host.com/status"
	if status != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", status)
		t.Error("statusEP should return the right endpoint")
	}

	update := statusUpdateEP(c, true)
	expected = "http://host.com/status?processingDisabled=true"
	if update != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", update)
		t.Error("statusUpdateEP should return the right endpoint")
	}
}

func TestZoneExtrasEP(t *testing.T) {
	backendIDs := zoneBackendIDsEP(c)
	expected := "http://host.com/zones/backendids"
	if backendIDs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", backendIDs)
		t.Error("zoneBackendIDsEP should return the right endpoint")
	}

	ignoreAccess := true
	deleted := zoneDeletedChangesEP(c, DeletedZonesFilter{
		NameFilter:   "foo*",
		StartFrom:    "next",
		MaxItems:     5,
		IgnoreAccess: &ignoreAccess,
	})
	expected = "http://host.com/zones/deleted/changes?nameFilter=foo*&startFrom=next&maxItems=5&ignoreAccess=true"
	if deleted != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", deleted)
		t.Error("zoneDeletedChangesEP should return the right endpoint")
	}

	aclRules := zoneACLRulesEP(c, "123")
	expected = "http://host.com/zones/123/acl/rules"
	if aclRules != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", aclRules)
		t.Error("zoneACLRulesEP should return the right endpoint")
	}
}

func TestRecordSetExtrasEP(t *testing.T) {
	count := recordSetCountEP(c, "123")
	expected := "http://host.com/zones/123/recordsetcount"
	if count != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", count)
		t.Error("recordSetCountEP should return the right endpoint")
	}

	history := recordSetChangeHistoryEP(c, RecordSetChangeHistoryFilter{
		ZoneID:     "z1",
		FQDN:       "ok.",
		RecordType: "A",
		StartFrom:  "1",
		MaxItems:   2,
	})
	expected = "http://host.com/recordsetchange/history?zoneId=z1&fqdn=ok.&recordType=A&startFrom=1&maxItems=2"
	if history != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", history)
		t.Error("recordSetChangeHistoryEP should return the right endpoint")
	}

	failures := recordSetChangesFailureEP(c, "z2", ListFilter{
		StartFrom: "2",
		MaxItems:  3,
	})
	expected = "http://host.com/metrics/health/zones/z2/recordsetchangesfailure?startFrom=2&maxItems=3"
	if failures != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", failures)
		t.Error("recordSetChangesFailureEP should return the right endpoint")
	}
}

func TestGroupExtrasEP(t *testing.T) {
	change := groupChangeEP(c, "123")
	expected := "http://host.com/groups/change/123"
	if change != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", change)
		t.Error("groupChangeEP should return the right endpoint")
	}

	domains := groupValidDomainsEP(c)
	expected = "http://host.com/groups/valid/domains"
	if domains != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", domains)
		t.Error("groupValidDomainsEP should return the right endpoint")
	}
}

func TestUserExtrasEP(t *testing.T) {
	user := userEP(c, "ok")
	expected := "http://host.com/users/ok"
	if user != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", user)
		t.Error("userEP should return the right endpoint")
	}

	lock := userLockEP(c, "ok")
	expected = "http://host.com/users/ok/lock"
	if lock != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", lock)
		t.Error("userLockEP should return the right endpoint")
	}

	unlock := userUnlockEP(c, "ok")
	expected = "http://host.com/users/ok/unlock"
	if unlock != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", unlock)
		t.Error("userUnlockEP should return the right endpoint")
	}
}

func TestBatchReviewEP(t *testing.T) {
	approve := batchRecordChangeApproveEP(c, "123")
	expected := "http://host.com/zones/batchrecordchanges/123/approve"
	if approve != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", approve)
		t.Error("batchRecordChangeApproveEP should return the right endpoint")
	}

	reject := batchRecordChangeRejectEP(c, "123")
	expected = "http://host.com/zones/batchrecordchanges/123/reject"
	if reject != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", reject)
		t.Error("batchRecordChangeRejectEP should return the right endpoint")
	}

	cancel := batchRecordChangeCancelEP(c, "123")
	expected = "http://host.com/zones/batchrecordchanges/123/cancel"
	if cancel != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", cancel)
		t.Error("batchRecordChangeCancelEP should return the right endpoint")
	}
}

func TestZoneChangesFailureEP(t *testing.T) {
	failures := zoneChangesFailureEP(c, ListFilter{
		StartFrom: "2",
		MaxItems:  3,
	})
	expected := "http://host.com/metrics/health/zonechangesfailure?startFrom=2&maxItems=3"
	if failures != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", failures)
		t.Error("zoneChangesFailureEP should return the right endpoint")
	}
}

func TestBuildQuery(t *testing.T) {
	query := buildQuery(ListFilter{
		MaxItems:   1,
		NameFilter: "foo",
	}, "theNameFilter")
	expected := "?theNameFilter=foo&maxItems=1"

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildQuery should return the right string")
	}
}

func TestBuildQueryWithNoQuery(t *testing.T) {
	query := buildQuery(ListFilter{}, "")
	expected := ""

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildQuery should return the right string")
	}
}

func TestBuildStartMaxQuery(t *testing.T) {
	query := buildStartMaxQuery(ListFilter{
		StartFrom: "2",
		MaxItems:  3,
	})
	expected := "?startFrom=2&maxItems=3"

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildStartMaxQuery should return the right string")
	}
}

func TestBuildDeletedZonesQuery(t *testing.T) {
	ignoreAccess := true
	query := buildDeletedZonesQuery(DeletedZonesFilter{
		NameFilter:   "foo*",
		StartFrom:    "next",
		MaxItems:     5,
		IgnoreAccess: &ignoreAccess,
	})
	expected := "?nameFilter=foo*&startFrom=next&maxItems=5&ignoreAccess=true"

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildDeletedZonesQuery should return the right string")
	}
}

func TestBuildRecordSetChangeHistoryQuery(t *testing.T) {
	query := buildRecordSetChangeHistoryQuery(RecordSetChangeHistoryFilter{
		ZoneID:     "zone-1",
		FQDN:       "ok.",
		RecordType: "A",
		StartFrom:  "1",
		MaxItems:   2,
	})
	expected := "?zoneId=zone-1&fqdn=ok.&recordType=A&startFrom=1&maxItems=2"

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildRecordSetChangeHistoryQuery should return the right string")
	}
}

func TestBuildPrometheusQuery(t *testing.T) {
	query := buildPrometheusQuery([]string{"one", "two"})
	expected := "?name=one&name=two"

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildPrometheusQuery should return the right string")
	}
}

func TestBuildGlobalListQuery(t *testing.T) {
	query := buildGlobalListQuery(GlobalListFilter{
		MaxItems:         1,
		RecordNameFilter: "foo",
	})
	expected := "?recordNameFilter=foo&maxItems=1"

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildGlobalListQuery should return the right string")
	}
}

func TestBuildGlobalListQueryWithNoQuery(t *testing.T) {
	query := buildGlobalListQuery(GlobalListFilter{})
	expected := ""

	if query != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", query)
		t.Error("buildGlobalListQuery should return the right string")
	}
}
