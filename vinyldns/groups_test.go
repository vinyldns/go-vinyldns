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
	"testing"

	"github.com/gobs/pretty"
)

func TestGroupCreate(t *testing.T) {
	groupJSON, err := readFile("test-fixtures/groups/group.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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
	groupJSON, err := readFile("test-fixtures/groups/group.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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
	groupJSON, err := readFile("test-fixtures/groups/group.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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
	groupsJSON, err := readFile("test-fixtures/groups/groups.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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

func TestGroupsListAll(t *testing.T) {
	groupsListJSON1, err := readFile("test-fixtures/groups/groups-list-1.json")
	if err != nil {
		t.Error(err)
	}
	groupsListJSON2, err := readFile("test-fixtures/groups/groups-list-2.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/groups?maxItems=1",
			code:     200,
			body:     groupsListJSON1,
		},
		{
			endpoint: "http://host.com/groups?startFrom=2&maxItems=1",
			code:     200,
			body:     groupsListJSON2,
		},
	})

	defer server.Close()

	if _, err := client.GroupsListAll(ListFilter{
		MaxItems: 200,
	}); err == nil {
		t.Error("Expected error -- MaxItems must be between 1 and 100")
	}

	groups, err := client.GroupsListAll(ListFilter{
		MaxItems: 1,
	})
	if err != nil {
		t.Error(err)
	}

	if len(groups) != 2 {
		t.Error("Expected 2 Groups; got ", len(groups))
	}

	if groups[0].ID != "1" {
		t.Error("Expected Group.ID to be 1")
	}

	if groups[1].ID != "2" {
		t.Error("Expected Group.ID to be 2")
	}
}

func TestGroup(t *testing.T) {
	groupJSON, err := readFile("test-fixtures/groups/group.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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
	groupAdminsJSON, err := readFile("test-fixtures/groups/group-admins.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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
	groupMembersJSON, err := readFile("test-fixtures/groups/group-members.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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
	groupActivityJSON, err := readFile("test-fixtures/groups/group-activity.json")
	if err != nil {
		t.Error(err)
	}
	server, client := testTools([]testToolsConfig{
		{
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
