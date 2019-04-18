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
	"fmt"
	"io"
)

// Zones retrieves the list of zones a user has access to.
func (c *Client) Zones() ([]Zone, error) {
	zones := &Zones{}
	err := resourceRequest(c, zonesEP(c), "GET", nil, zones)
	if err != nil {
		return []Zone{}, err
	}

	return zones.Zones, nil
}

// ZonesListAll retrieves the complete list of zones with the ListFilter criteria passed.
// Handles paging through results on the user's behalf.
func (c *Client) ZonesListAll(filter ListFilter) ([]Zone, error) {
	if filter.MaxItems > 100 {
		return nil, fmt.Errorf("MaxItems must be between 1 and 100")
	}

	var zones []Zone

	for {
		resp, err := c.zonesList(filter)
		if err != nil {
			return nil, err
		}

		zones = append(zones, resp.Zones...)
		filter.StartFrom = resp.NextID

		if len(filter.StartFrom) == 0 {
			return zones, nil
		}
	}
}

// Zone retrieves the Zone whose ID it's passed.
func (c *Client) Zone(id string) (Zone, error) {
	zone := &ZoneResponse{}
	err := resourceRequest(c, zoneEP(c, id), "GET", nil, zone)
	if err != nil {
		return Zone{}, err
	}

	return zone.Zone, nil
}

// ZoneCreate creates the Zone it's passed.
func (c *Client) ZoneCreate(z *Zone) (*ZoneUpdateResponse, error) {
	zJSON, err := json.Marshal(z)
	if err != nil {
		return nil, err
	}
	var resource = &ZoneUpdateResponse{}
	err = resourceRequest(c, zonesEP(c), "POST", zJSON, resource)
	if err != nil {
		return &ZoneUpdateResponse{}, err
	}

	return resource, nil
}

// ZoneUpdate updates the Zone whose ID it's passed.
func (c *Client) ZoneUpdate(zoneID string, z *Zone) (*ZoneUpdateResponse, error) {
	zJSON, err := json.Marshal(z)
	if err != nil {
		return nil, err
	}
	var resource = &ZoneUpdateResponse{}
	err = resourceRequest(c, zoneEP(c, zoneID), "PUT", zJSON, resource)
	if err != nil {
		return &ZoneUpdateResponse{}, err
	}

	return resource, nil
}

// ZoneDelete deletes the Zone whose ID it's passed.
func (c *Client) ZoneDelete(zoneID string) (*ZoneUpdateResponse, error) {
	resource := &ZoneUpdateResponse{}
	err := resourceRequest(c, zoneEP(c, zoneID), "DELETE", nil, resource)
	if err != nil {
		return &ZoneUpdateResponse{}, err
	}

	return resource, nil
}

// ZoneExists returns true if a zone request does not 404
// Otherwise, it returns false
func (c *Client) ZoneExists(zoneID string) (bool, error) {
	resp, err := doClientReq(c, "GET", zoneEP(c, zoneID))
	if err != nil {
		return false, err
	}

	code := resp.StatusCode
	if code == 404 {
		return false, nil
	}

	return true, nil
}

// ZoneHistory retrieves the ZoneHistory for the Zone whose ID it's passed.
func (c *Client) ZoneHistory(id string) (*ZoneHistory, error) {
	zh := &ZoneHistory{}
	err := resourceRequest(c, zoneHistoryEP(c, id), "GET", nil, zh)
	if err != nil {
		return &ZoneHistory{}, err
	}

	return zh, nil
}

// ZoneChange retrieves the ZoneChange matching the Zone ID and
// and ZoneChange ID it's passed.
func (c *Client) ZoneChange(zoneID, zoneChangeID string) (ZoneChange, error) {
	zc := ZoneChange{}
	history, err := c.ZoneHistory(zoneID)
	if err != nil {
		return zc, err
	}

	for _, each := range history.ZoneChanges {
		fmt.Println(each)
		if each.ID == zoneChangeID {
			return each, nil
		}
	}

	return zc, nil
}

// RecordSetLimit is the highest number of records the vinyldns server will allow at once
// TODO: is there a way to get this limit directly from vinyldns?
const RecordSetLimit = 100

// RecordSetCollector creates a function to retrieve the next set of recordsets.
// To retrieve *all* recordsets, call that function repeatedly until err == // io.EOF
func (c *Client) RecordSetCollector(zoneID string, limit int) (func() ([]RecordSet, error), error) {
	if limit > RecordSetLimit {
		return nil, fmt.Errorf("Limit must be zero or not greater than %d", RecordSetLimit)
	}

	var nextID string
	var recordSets []RecordSet
	var err error

	return func() ([]RecordSet, error) {
		if err != nil {
			return nil, err
		}

		for {
			rss := &RecordSetsResponse{}
			err = resourceRequest(c, recordSetsListEP(c, zoneID, ListFilter{
				StartFrom: nextID,
				MaxItems:  limit,
			}), "GET", nil, rss)
			if err != nil {
				return nil, err
			}
			recordSets = append(recordSets, rss.RecordSets...)

			nextID = rss.NextID
			if len(nextID) == 0 {
				// keep from trying to get more records
				err = io.EOF

				break
			}
		}

		// return at most `limit` records and remove those returned from recordSets
		max := limit
		if max == 0 || max > len(recordSets) {
			max = len(recordSets)
		}
		r := recordSets[:max]
		recordSets = recordSets[max:]
		return r, err
	}, nil
}

// RecordSets retrieves a list of RecordSets from a Zone.
func (c *Client) RecordSets(id string) ([]RecordSet, error) {
	collector, err := c.RecordSetCollector(id, 0)
	if err != nil {
		return nil, err
	}

	var recordSets []RecordSet
	for {
		rs, err := collector()
		if err != nil && err != io.EOF {
			return nil, err
		}
		recordSets = append(recordSets, rs...)
		if err == io.EOF {
			break
		}
	}

	return recordSets, nil
}

// RecordSetsListAll retrieves the complete list of record sets with the ListFilter criteria passed.
// Handles paging through results on the user's behalf.
func (c *Client) RecordSetsListAll(zoneID string, filter ListFilter) ([]RecordSet, error) {
	if filter.MaxItems > 100 {
		return nil, fmt.Errorf("MaxItems must be between 1 and 100")
	}

	var rss []RecordSet

	for {
		resp, err := c.recordSetsList(zoneID, filter)
		if err != nil {
			return nil, err
		}

		rss = append(rss, resp.RecordSets...)
		filter.StartFrom = resp.NextID

		if len(filter.StartFrom) == 0 {
			return rss, nil
		}
	}
}

// RecordSet retrieves the record matching the Zone ID and RecordSet ID it's passed.
func (c *Client) RecordSet(zoneID, recordSetID string) (RecordSet, error) {
	rs := &RecordSetResponse{}
	err := resourceRequest(c, recordSetEP(c, zoneID, recordSetID), "GET", nil, rs)
	if err != nil {
		return RecordSet{}, err
	}

	return rs.RecordSet, nil
}

// RecordSetCreate creates the RecordSet it's passed in the Zone whose ID it's passed.
func (c *Client) RecordSetCreate(zoneID string, rs *RecordSet) (*RecordSetUpdateResponse, error) {
	rsJSON, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	var resource = &RecordSetUpdateResponse{}
	err = resourceRequest(c, recordSetsEP(c, zoneID), "POST", rsJSON, resource)
	if err != nil {
		return &RecordSetUpdateResponse{}, err
	}

	return resource, nil
}

// RecordSetUpdate updates the RecordSet matching the Zone ID and RecordSetID it's passed.
func (c *Client) RecordSetUpdate(zoneID, recordSetID string, rs *RecordSet) (*RecordSetUpdateResponse, error) {
	rsJSON, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}
	var resource = &RecordSetUpdateResponse{}
	err = resourceRequest(c, recordSetEP(c, zoneID, recordSetID), "PUT", rsJSON, resource)
	if err != nil {
		return &RecordSetUpdateResponse{}, err
	}

	return resource, nil
}

// RecordSetDelete deletes the RecordSet matching the Zone ID and RecordSet ID it's passed.
func (c *Client) RecordSetDelete(zoneID, recordSetID string) (*RecordSetUpdateResponse, error) {
	resource := &RecordSetUpdateResponse{}
	err := resourceRequest(c, recordSetEP(c, zoneID, recordSetID), "DELETE", nil, resource)
	if err != nil {
		return &RecordSetUpdateResponse{}, err
	}

	return resource, nil
}

// RecordSetChange retrieves the RecordSetChange matching the Zone, RecordSet, and Change IDs
// it's passed.
func (c *Client) RecordSetChange(zoneID, recordSetID, changeID string) (*RecordSetChange, error) {
	rsc := &RecordSetChange{}
	err := resourceRequest(c, recordSetChangeEP(c, zoneID, recordSetID, changeID), "GET", nil, rsc)
	if err != nil {
		return &RecordSetChange{}, err
	}

	return rsc, nil
}

// Groups retrieves a list of Groups that the requester is a part of.
func (c *Client) Groups() ([]Group, error) {
	groups := &Groups{}
	err := resourceRequest(c, groupsEP(c), "GET", nil, groups)
	if err != nil {
		return []Group{}, err
	}

	return groups.Groups, nil
}

// GroupCreate creates the Group it's passed.
func (c *Client) GroupCreate(g *Group) (*Group, error) {
	gJSON, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	var resource = &Group{}
	err = resourceRequest(c, groupsEP(c), "POST", gJSON, resource)
	if err != nil {
		return &Group{}, err
	}

	return resource, nil
}

// Group gets the Group whose ID it's passed.
func (c *Client) Group(groupID string) (*Group, error) {
	group := &Group{}
	err := resourceRequest(c, groupEP(c, groupID), "GET", nil, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// GroupDelete deletes the Group whose ID it's passed.
func (c *Client) GroupDelete(groupID string) (*Group, error) {
	group := &Group{}
	err := resourceRequest(c, groupEP(c, groupID), "DELETE", nil, group)
	if err != nil {
		return nil, err
	}

	return group, nil
}

// GroupUpdate updates the Group whose ID it's passed.
func (c *Client) GroupUpdate(groupID string, g *Group) (*Group, error) {
	gJSON, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	var resource = &Group{}
	err = resourceRequest(c, groupEP(c, groupID), "PUT", gJSON, resource)
	if err != nil {
		return &Group{}, err
	}

	return resource, nil
}

// GroupAdmins returns an array of Users that are admins
// associated with the Group whose GroupID it's passed.
func (c *Client) GroupAdmins(groupID string) ([]User, error) {
	admins := &GroupAdmins{}
	err := resourceRequest(c, groupAdminsEP(c, groupID), "GET", nil, admins)
	if err != nil {
		return nil, err
	}

	return admins.GroupAdmins, nil
}

// GroupMembers returns an array of Users that are members
// associated with the Group whose GroupID it's passed.
func (c *Client) GroupMembers(groupID string) ([]User, error) {
	members := &GroupMembers{}
	err := resourceRequest(c, groupMembersEP(c, groupID), "GET", nil, members)
	if err != nil {
		return nil, err
	}

	return members.GroupMembers, nil
}

// GroupActivity returns group change activity
// associated with the Group whose GroupID it's passed.
func (c *Client) GroupActivity(groupID string) (*GroupChanges, error) {
	activity := &GroupChanges{}
	err := resourceRequest(c, groupActivityEP(c, groupID), "GET", nil, activity)
	if err != nil {
		return nil, err
	}

	return activity, nil
}

// BatchRecordChanges returns the list of batch record changes
func (c *Client) BatchRecordChanges() ([]RecordChange, error) {
	changes := &BatchRecordChanges{}
	err := resourceRequest(c, batchRecordChangesEP(c), "GET", nil, changes)
	if err != nil {
		return nil, err
	}

	return changes.BatchChanges, nil
}

// BatchRecordChange returns the batch record change
// associated with the change whose ID it's passed.
func (c *Client) BatchRecordChange(changeID string) (*BatchRecordChange, error) {
	change := &BatchRecordChange{}
	err := resourceRequest(c, batchRecordChangeEP(c, changeID), "GET", nil, change)
	if err != nil {
		return nil, err
	}

	return change, nil
}

// BatchRecordChangeCreate creates the batch record change it's passed.
func (c *Client) BatchRecordChangeCreate(change *BatchRecordChange) (*BatchRecordChangeUpdateResponse, error) {
	cJSON, err := json.Marshal(change)
	if err != nil {
		return nil, err
	}
	var resource = &BatchRecordChangeUpdateResponse{}
	err = resourceRequest(c, batchRecordChangesEP(c), "POST", cJSON, resource)
	if err != nil {
		return &BatchRecordChangeUpdateResponse{}, err
	}

	return resource, nil
}
