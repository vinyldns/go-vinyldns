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

import "testing"

func TestUser(t *testing.T) {
	userJSON := `{"id":"ok","userName":"ok","groupId":["ok-group"]}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/users/ok",
			code:     200,
			body:     userJSON,
		},
	})
	defer server.Close()

	user, err := client.User("ok")
	if err != nil {
		t.Error(err)
	}
	if user.ID != "ok" {
		t.Error("Expected user ID to be ok")
	}
}

func TestUserLock(t *testing.T) {
	lockJSON := `{"id":"ok","userName":"ok","lockStatus":"Locked"}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/users/ok/lock",
			code:     200,
			body:     lockJSON,
		},
	})
	defer server.Close()

	user, err := client.UserLock("ok")
	if err != nil {
		t.Error(err)
	}
	if user.LockStatus != "Locked" {
		t.Error("Expected lock status to be Locked")
	}
}

func TestUserUnlock(t *testing.T) {
	unlockJSON := `{"id":"ok","userName":"ok","lockStatus":"Unlocked"}`
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/users/ok/unlock",
			code:     200,
			body:     unlockJSON,
		},
	})
	defer server.Close()

	user, err := client.UserUnlock("ok")
	if err != nil {
		t.Error(err)
	}
	if user.LockStatus != "Unlocked" {
		t.Error("Expected lock status to be Unlocked")
	}
}
