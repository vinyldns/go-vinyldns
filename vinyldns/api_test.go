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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gobs/pretty"
)

type testToolsConfig struct {
	endpoint string
	code     int
	body     string
}

func testTools(configs []testToolsConfig) (*httptest.Server, *Client) {
	host := "http://host.com"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, c := range configs {
			if c.endpoint == r.RequestURI {
				w.WriteHeader(c.code)
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, c.body)
				return
			}
		}

		fmt.Printf("Requested: %s\n", r.RequestURI)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}

	client := &Client{
		"accessToken",
		"secretToken",
		host,
		&http.Client{Transport: tr},
	}

	return server, client
}

func TestRecordSetsListAll(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/123/recordsets?maxItems=1",
			code:     200,
			body:     recordSetsListJSON1,
		},
		testToolsConfig{
			endpoint: "http://host.com/zones/123/recordsets?startFrom=2&maxItems=1",
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
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
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

func TestRecordSetCollector(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/123/recordsets?maxItems=3",
			code:     200,
			body:     recordSetsJSON,
		},
		testToolsConfig{
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
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
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
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
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

	r, err := client.RecordSetCreate("123", rs)
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
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/123/recordsets/456",
			code:     202,
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

	r, err := client.RecordSetUpdate("123", "456", rs)
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
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
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

func TestRecordSetChange(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
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

func TestGroupCreate(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups",
			code:     200,
			body:     groupJSON,
		},
	})

	defer server.Close()

	group := &Group{
		Name:        "test-group",
		Email:       "test@test.com",
		Description: "this is a description",
	}

	g, err := client.GroupCreate(group)
	if err != nil {
		t.Log(pretty.PrettyFormat(g))
		t.Error(err)
	}
	if g.Name != "test-group" {
		t.Error("Expected groupResponse.Name to be test-group")
	}
	if g.Email != "test@test.com" {
		t.Error("Expected groupResponse.Email to be test@test.com")
	}
	if g.Description != "this is a description" {
		t.Error("Expected groupResponse.Description to be this is a description")
	}
}

func TestGroupUpdate(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups/123",
			code:     200,
			body:     groupJSON,
		},
	})

	defer server.Close()

	group := &Group{
		ID:          "123",
		Name:        "test-group",
		Email:       "test@test.com",
		Description: "this is a description",
	}

	g, err := client.GroupUpdate("123", group)
	if err != nil {
		t.Log(pretty.PrettyFormat(g))
		t.Error(err)
	}
	if g.Name != "test-group" {
		t.Error("Expected groupResponse.Name to be test-group")
	}
	if g.Email != "test@test.com" {
		t.Error("Expected groupResponse.Email to be test@test.com")
	}
	if g.Description != "this is a description" {
		t.Error("Expected groupResponse.Description to be 'this is a description'")
	}
}

func TestGroupDelete(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups/123",
			code:     200,
			body:     groupJSON,
		},
	})

	defer server.Close()

	g, err := client.GroupDelete("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(g))
		t.Error(err)
	}
	if g.Name != "test-group" {
		t.Error("Expected groupResponse.Name to be test-group")
	}
	if g.Email != "test@test.com" {
		t.Error("Expected groupResponse.Email to be test@test.com")
	}
	if g.Description != "this is a description" {
		t.Error("Expected groupResponse.Description to be 'this is a description'")
	}
}

func TestGroups(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups",
			code:     200,
			body:     groupsJSON,
		},
	})

	defer server.Close()

	groups, err := client.Groups()
	if err != nil {
		t.Log(pretty.PrettyFormat(groups))
		t.Error(err)
	}
	if len(groups) != 2 {
		t.Error("Expected 2 Groups")
	}
	for _, g := range groups {
		if g.Name == "" {
			t.Error("Expected group.Name to have a value")
		}
		if g.Email == "" {
			t.Error("Expected group.Email to have a value")
		}
		if g.Created == "" {
			t.Error("Expected group.Created to have a value")
		}
		if g.Status == "" {
			t.Error("Expected group.Status to have a value")
		}
		if g.ID == "" {
			t.Error("Expected group.Id to have a value")
		}
	}
}

func TestGroup(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups/123",
			code:     200,
			body:     groupJSON,
		},
	})

	defer server.Close()

	g, err := client.Group("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(g))
		t.Error(err)
	}
	if g.Name != "test-group" {
		t.Error("Expected groupResponse.Name to be test-group")
	}
	if g.Email != "test@test.com" {
		t.Error("Expected groupResponse.Email to be test@test.com")
	}
	if g.Description != "this is a description" {
		t.Error("Expected groupResponse.Description to be 'this is a description'")
	}
}

func TestGroupAdmins(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups/123/admins",
			code:     200,
			body:     groupAdminsJSON,
		},
	})

	defer server.Close()

	admins, err := client.GroupAdmins("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(admins))
		t.Error(err)
	}

	a := admins[0]
	if a.UserName != "jdoe201" {
		t.Error("Expected GroupAdmins[0].Username to be 'jdoe201'")
	}
}

func TestGroupMembers(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups/123/members",
			code:     200,
			body:     groupMembersJSON,
		},
	})

	defer server.Close()

	members, err := client.GroupMembers("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(members))
		t.Error(err)
	}

	m := members[0]
	if m.UserName != "jdoe201" {
		t.Error("Expected GroupMembers[0].Username to be 'jdoe201'")
	}
}

func TestGroupActivity(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/groups/123/activity",
			code:     200,
			body:     groupActivityJSON,
		},
	})

	defer server.Close()

	activity, err := client.GroupActivity("123")
	if err != nil {
		t.Log(pretty.PrettyFormat(activity))
		t.Error(err)
	}

	c := activity.Changes[0]
	if c.NewGroup.Name != "test-list-group-activity-max-item-success" {
		t.Error("Expected GroupActivity.Changes[0].NewGroup.Name to be 'test-list-group-activity-max-item-success'")
	}
	if c.UserID != "some-user" {
		t.Error("Expected GroupActivity.Changes[0].UserID to be 'some-user'")
	}
	if c.ChangeType != "Update" {
		t.Error("Expected GroupActivity.Changes[0].UserID to be 'some-user'")
	}
	if c.Created != "1488480605378" {
		t.Error("Expected GroupActivity.Changes[0].Created to be '1488480605378'")
	}
}

func TestBatchRecordChanges(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/batchrecordchanges",
			code:     200,
			body:     batchRecordChangesJSON,
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

func TestBatchRecordChange(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/batchrecordchanges/123",
			code:     200,
			body:     batchRecordChangeJSON,
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
	server, client := testTools([]testToolsConfig{
		testToolsConfig{
			endpoint: "http://host.com/zones/batchrecordchanges",
			code:     200,
			body:     batchRecordChangeCreateJSON,
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
