// +build integration

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
	"testing"
	"time"
)

// client() assumes a VinylDNS is running on localhost:9000 witht he default access and secret keys
// see `make start-api` for a Make task in starting VinylDNS
func client() *Client {
	client := NewClient(ClientConfiguration{
		"okAccessKey",
		"okSecretKey",
		"http://localhost:9000",
	})

	return client
}

func TestIntegration(t *testing.T) {
	c := client()
	users := []User{
		User{
			UserName:  "ok",
			FirstName: "ok",
			LastName:  "ok",
			Email:     "test@test.com",
			ID:        "ok",
		},
	}
	group, err := c.GroupCreate(&Group{
		Name:        "test-group",
		Description: "a test group",
		Email:       "test@vinyldns.com",
		Admins:      users,
		Members:     users,
	})

	if err != nil {
		t.Error(err)
	}

	connection := &ZoneConnection{
		Name:          "ok.",
		KeyName:       "vinyldns.",
		Key:           "nzisn+4G2ldMn0q1CV3vsg==",
		PrimaryServer: "vinyldns-bind9",
	}

	zone := &Zone{
		Name:               "ok.",
		Email:              "email@email.com",
		AdminGroupID:       group.ID,
		Connection:         connection,
		TransferConnection: connection,
	}

	zc, err := c.ZoneCreate(zone)
	if err != nil {
		t.Error(err)
	}

	createdZoneID := zc.Zone.ID
	limit := 10
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		zg, err := c.Zone(createdZoneID)
		if err == nil && zg.ID != createdZoneID {
			t.Error(fmt.Sprintf("unable to get zone %s", createdZoneID))
		}
		if err == nil && zg.ID == createdZoneID {
			break
		}

		if i == (limit - 1) {
			fmt.Printf("%d retries reached in polling for zone %s", limit, createdZoneID)
			t.Error(err)
		}
	}

	_, err = c.ZoneDelete(createdZoneID)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < limit; time.Sleep(10 * time.Second) {
		i++

		exists, err := c.ZoneExists(createdZoneID)
		if err != nil {
			t.Error(err)
			break
		}

		if !exists {
			break
		}

		if i == (limit - 1) {
			fmt.Printf("%d retries reached in waiting for zone deletion of %s", limit, createdZoneID)
			t.Error(err)
		}
	}

	_, err = c.GroupDelete(group.ID)
	if err != nil {
		t.Error(err)
	}
}
