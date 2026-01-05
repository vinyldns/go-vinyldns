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

// User retrieves a user by ID or username.
func (c *Client) User(userIdentifier string) (UserInfo, error) {
	user := &UserInfo{}
	err := resourceRequest(c, userEP(c, userIdentifier), "GET", nil, user)
	if err != nil {
		return UserInfo{}, err
	}

	return *user, nil
}

// UserLock locks a user account.
func (c *Client) UserLock(userID string) (UserInfo, error) {
	user := &UserInfo{}
	err := resourceRequest(c, userLockEP(c, userID), "PUT", nil, user)
	if err != nil {
		return UserInfo{}, err
	}

	return *user, nil
}

// UserUnlock unlocks a user account.
func (c *Client) UserUnlock(userID string) (UserInfo, error) {
	user := &UserInfo{}
	err := resourceRequest(c, userUnlockEP(c, userID), "PUT", nil, user)
	if err != nil {
		return UserInfo{}, err
	}

	return *user, nil
}
