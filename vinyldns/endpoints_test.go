/*
Copyright 2019 Comcast Cable Communications Management, LLC
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

func TestZoneEP(t *testing.T) {
	zone := zoneEP(c, "123")
	expected := "http://host.com/zones/123"

	if zone != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zone)
		t.Error("zoneEP should return the right endpoint")
	}
}

func TestZoneHistoryEP(t *testing.T) {
	zone := zoneHistoryEP(c, "123")
	expected := "http://host.com/zones/123/history"

	if zone != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", zone)
		t.Error("zoneHistoryEP should return the right endpoint")
	}
}

func TestRecordSetsEp(t *testing.T) {
	rs := recordSetsEp(c, "123", "", 0)
	expected := "http://host.com/zones/123/recordsets"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetsEp should return the right endpoint")
	}

	rs = recordSetsEp(c, "543", "nextplease", 0)
	expected = "http://host.com/zones/543/recordsets?startFrom=nextplease"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetsEp should return the right endpoint")
	}

	rs = recordSetsEp(c, "7", "nextplease", 99)
	expected = "http://host.com/zones/7/recordsets?startFrom=nextplease?limit=99"

	if rs != expected {
		fmt.Printf("Expected: %s", expected)
		fmt.Printf("Actual: %s", rs)
		t.Error("recordSetsEp should return the right endpoint")
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
