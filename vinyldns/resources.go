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
	"strconv"
	"strings"
)

// Error represents an error from the
// vinyldns API
type Error struct {
	RequestURL    string
	RequestMethod string
	RequestBody   string
	ResponseBody  string
	ResponseCode  int
}

func (d Error) Error() string {
	components := []string{
		"Request URL:",
		d.RequestURL,
		"Request Method:",
		d.RequestMethod,
		"Request body:",
		d.RequestBody,
		"Response code: ",
		strconv.Itoa(d.ResponseCode),
		"Response body:",
		d.ResponseBody}
	return strings.Join(components, "\n")
}

// ListFilter represents the list query parameters that may be passed to
// VinylDNS API endpoints such as /zones and /recordsets
type ListFilter struct {
	NameFilter string
	StartFrom  string
	MaxItems   int
}

// User represents a vinyldns user.
type User struct {
	ID        string `json:"id,omitempty"`
	UserName  string `json:"userName,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
	Created   string `json:"created,omitempty"`
}

// Groups is a slice of groups
type Groups struct {
	Groups []Group `json:"groups"`
}

// Group represents a vinyldns group.
type Group struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Email       string `json:"email,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	Created     string `json:"created,omitempty"`
	Members     []User `json:"members"`
	Admins      []User `json:"admins"`
}

// GroupAdmins is a slice of Users
type GroupAdmins struct {
	GroupAdmins []User `json:"admins"`
}

// GroupMembers is a slice of Users
type GroupMembers struct {
	GroupMembers []User `json:"members"`
}

// GroupChange represents a group change event object.
type GroupChange struct {
	UserID     string `json:"userId,omitempty"`
	Created    string `json:"created,omitempty"`
	ChangeType string `json:"changeType,omitempty"`
	NewGroup   Group  `json:"newGroup,omitempty"`
	OldGroup   Group  `json:"oldGroup,omitempty"`
}

// GroupChanges is represents the group changes.
type GroupChanges struct {
	Changes []GroupChange `json:"changes"`
}

// BatchRecordChanges represents a list of record changes,
// as returned by the list batch changes VinylDNS API endpoint.
type BatchRecordChanges struct {
	BatchChanges []RecordChange `json:"batchChanges,omitempty"`
}

// RecordChange represents an individual batch record change.
type RecordChange struct {
	ID               string     `json:"id,omitempty"`
	Status           string     `json:"status,omitempty"`
	ChangeType       string     `json:"changeType,omitempty"`
	RecordName       string     `json:"recordName,omitempty"`
	TTL              int        `json:"ttl,omitempty"`
	Type             string     `json:"type,omitempty"`
	ZoneName         string     `json:"zoneName,omitempty"`
	InputName        string     `json:"inputName,omitempty"`
	ZoneID           string     `json:"zoneId,omitempty"`
	TotalChanges     int        `json:"totalChanges,omitempty"`
	UserName         string     `json:"userName,omitempty"`
	Comments         string     `json:"comments,omitempty"`
	UserID           string     `json:"userId,omitempty"`
	CreatedTimestamp string     `json:"createdTimestamp,omitempty"`
	Record           RecordData `json:"data,omitempty"`
	OwnerGroupID     string     `json:"ownerGroupId,omitempty"`
}

// BatchRecordChangeUpdateResponse is represents a batch record change create or update response
type BatchRecordChangeUpdateResponse struct {
	Comments     string         `json:"comments,omitempty"`
	OwnerGroupID string         `json:"ownerGroupId,omitempty"`
	Changes      []RecordChange `json:"changes,omitempty"`
}

// RecordData is represents a batch record change record data.
type RecordData struct {
	Address  string `json:"address,omitempty"`
	CName    string `json:"cname,omitempty"`
	PTRDName string `json:"ptrdname,omitempty"`
}

// BatchRecordChange represents a batch record change API response.
type BatchRecordChange struct {
	ID               string         `json:"id,omitempty"`
	UserName         string         `json:"userName,omitempty"`
	UserID           string         `json:"userId,omitempty"`
	Status           string         `json:"status,omitempty"`
	Comments         string         `json:"comments,omitempty"`
	CreatedTimestamp string         `json:"createdTimestamp,omitempty"`
	OwnerGroupID     string         `json:"ownerGroupId,omitempty"`
	Changes          []RecordChange `json:"changes,omitempty"`
}
