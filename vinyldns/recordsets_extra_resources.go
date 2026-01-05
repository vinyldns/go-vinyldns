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

// RecordSetCount represents the record set count response.
type RecordSetCount struct {
	Count int `json:"count"`
}

// RecordSetChangeFailuresResponse represents failed record set changes.
type RecordSetChangeFailuresResponse struct {
	FailedRecordSetChanges []RecordSetChange `json:"failedRecordSetChanges"`
	StartFrom              int               `json:"startFrom,omitempty"`
	NextID                 int               `json:"nextId,omitempty"`
	MaxItems               int               `json:"maxItems,omitempty"`
}
