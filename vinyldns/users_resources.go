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

// UserInfo represents user details from user endpoints.
type UserInfo struct {
	ID         string   `json:"id,omitempty"`
	UserName   string   `json:"userName,omitempty"`
	GroupID    []string `json:"groupId,omitempty"`
	LockStatus string   `json:"lockStatus,omitempty"`
}
