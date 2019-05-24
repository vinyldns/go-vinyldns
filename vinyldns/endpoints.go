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
	"strings"
)

func zonesEP(c *Client) string {
	return concatStrs("", c.Host, "/zones")
}

func zonesListEP(c *Client, f ListFilter) string {
	params := []string{}
	query := "?"

	if f.NameFilter != "" {
		params = append(params, fmt.Sprintf("nameFilter=%s", f.NameFilter))
	}

	if f.StartFrom != "" {
		params = append(params, fmt.Sprintf("startFrom=%s", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	if len(params) == 0 {
		query = ""
	}

	return concatStrs("", zonesEP(c), query, strings.Join(params, "&"))
}

func zoneEP(c *Client, id string) string {
	return concatStrs("", zonesEP(c), "/", id)
}

func zoneNameEP(c *Client, name string) string {
	return concatStrs("", zonesEP(c), "/name/", name)
}

func zoneHistoryEP(c *Client, id string) string {
	return concatStrs("", zoneEP(c, id), "/history")
}

func recordSetsEP(c *Client, zoneID string) string {
	return concatStrs("", zoneEP(c, zoneID), "/recordsets")
}

func recordSetsListEP(c *Client, zoneID string, f ListFilter) string {
	params := []string{}
	query := "?"

	if f.NameFilter != "" {
		params = append(params, fmt.Sprintf("recordNameFilter=%s", f.NameFilter))
	}

	if f.StartFrom != "" {
		params = append(params, fmt.Sprintf("startFrom=%s", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	if len(params) == 0 {
		query = ""
	}

	return concatStrs("", recordSetsEP(c, zoneID), query, strings.Join(params, "&"))
}

func recordSetEP(c *Client, zoneID, recordSetID string) string {
	return concatStrs("", recordSetsEP(c, zoneID), "/", recordSetID)
}

func recordSetChangeEP(c *Client, zoneID, recordSetID, changeID string) string {
	return concatStrs("", recordSetEP(c, zoneID, recordSetID), "/changes/", changeID)
}

func groupsEP(c *Client) string {
	return concatStrs("", c.Host, "/groups")
}

func groupEP(c *Client, groupID string) string {
	return concatStrs("", groupsEP(c), "/", groupID)
}

func groupAdminsEP(c *Client, groupID string) string {
	return concatStrs("", groupEP(c, groupID), "/admins")
}

func groupMembersEP(c *Client, groupID string) string {
	return concatStrs("", groupEP(c, groupID), "/members")
}

func groupActivityEP(c *Client, groupID string) string {
	return concatStrs("", groupEP(c, groupID), "/activity")
}

func batchRecordChangesEP(c *Client) string {
	return concatStrs("", zonesEP(c), "/batchrecordchanges")
}

func batchRecordChangeEP(c *Client, changeID string) string {
	return concatStrs("", batchRecordChangesEP(c), "/", changeID)
}
