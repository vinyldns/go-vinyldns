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

import "strconv"

func zonesEP(c *Client) string {
	return concatStrs("", c.Host, "/zones")
}

func zoneEP(c *Client, id string) string {
	return concatStrs("", zonesEP(c), "/", id)
}

func zoneHistoryEP(c *Client, id string) string {
	return concatStrs("", zoneEP(c, id), "/history")
}

func recordSetsEP(c *Client, id string, startFrom string, limit int) string {
	ep := concatStrs("", zoneEP(c, id), "/recordsets")
	if len(startFrom) != 0 {
		ep += "?startFrom=" + startFrom
	}
	if limit > 0 {
		ep += "?limit=" + strconv.Itoa(limit)
	}
	return ep
}

func recordSetEP(c *Client, zoneID, recordSetID string) string {
	return concatStrs("", recordSetsEP(c, zoneID, "", 0), "/", recordSetID)
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
