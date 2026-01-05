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

// DeletedZoneInfo represents details for a deleted zone response.
type DeletedZoneInfo struct {
	ZoneChange     ZoneChange `json:"zoneChange"`
	AdminGroupName string     `json:"adminGroupName,omitempty"`
	UserName       string     `json:"userName,omitempty"`
	AccessLevel    string     `json:"accessLevel,omitempty"`
}

// DeletedZonesResponse represents the deleted zones response.
type DeletedZonesResponse struct {
	ZonesDeletedInfo []DeletedZoneInfo `json:"zonesDeletedInfo"`
	StartFrom        string            `json:"startFrom,omitempty"`
	NextID           string            `json:"nextId,omitempty"`
	MaxItems         int               `json:"maxItems,omitempty"`
	IgnoreAccess     bool              `json:"ignoreAccess,omitempty"`
}

// ZoneChangeFailuresResponse represents the failed zone changes response.
type ZoneChangeFailuresResponse struct {
	FailedZoneChanges []ZoneChange `json:"failedZoneChanges"`
	StartFrom         int          `json:"startFrom,omitempty"`
	NextID            int          `json:"nextId,omitempty"`
	MaxItems          int          `json:"maxItems,omitempty"`
}
