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

func TestAbandonedZoneEP(t *testing.T) {
	az := abandonedZonesEP(c, ListFilter{})
	expected := "http://host.com/zones/deleted/changes"

	if az != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", az)
		t.Error("deletedZoneChangesEP should return the right endpoint")
	}

	az = abandonedZonesEP(c, ListFilter{
		StartFrom: "nextplease",
	})
	expected = "http://host.com/zones/deleted/changes?startFrom=nextplease"

	if az != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", az)
		t.Error("deletedZoneChangesEP should return the right endpoint")
	}

	az = abandonedZonesEP(c, ListFilter{
		StartFrom: "nextplease",
		MaxItems:  99,
	})
	expected = "http://host.com/zones/deleted/changes?startFrom=nextplease&maxItems=99"

	if az != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", az)
		t.Error("deletedZoneChangesEP should return the right endpoint")
	}

	az = abandonedZonesEP(c, ListFilter{
		NameFilter: "foo",
		StartFrom:  "nextplease",
		MaxItems:   99,
	})
	expected = "http://host.com/zones/deleted/changes?nameFilter=foo&startFrom=nextplease&maxItems=99"

	if az != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", az)
		t.Error("deletedZoneChangesEP should return the right endpoint")
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
	rsc := recordSetChangesEP(c, "123", ListFilter{})
	expected := "http://host.com/zones/123/recordsetchanges"

	if rsc != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rsc)
		t.Error("recordSetChangesEP should return the right endpoint")
	}
}

func TestRecordSetChangesEPWithQuery(t *testing.T) {
	rsc := recordSetChangesEP(c, "123", ListFilter{
		MaxItems:  3,
		StartFrom: "1",
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
