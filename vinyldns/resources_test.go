/*
Copyright 2019 Comcast Cable Communications Management, LLC
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
	"testing"
)

func TestZoneMarshaling(t *testing.T) {
	zone := &Zone{
		ID: "123",
	}
	expected := "{\"id\":\"123\"}"
	z, err := json.Marshal(zone)
	if err != nil {
		t.Error(err)
	}

	if string(z) != expected {
		fmt.Println(string(z))
		t.Error("Failed to correctly marshal Zone")
	}
}

func TestRecordSetMarshaling(t *testing.T) {
	rs := &RecordSet{
		Records: []Record{
			{
				Type: "text",
			},
		},
	}
	expected := "{\"zoneId\":\"\",\"type\":\"\",\"ttl\":0,\"account\":\"\",\"records\":[{\"type\":\"text\"}]}"
	r, err := json.Marshal(rs)
	if err != nil {
		t.Error(err)
	}

	if string(r) != expected {
		fmt.Println(string(r))
		t.Error("Failed to correctly marshal RecordSet")
	}
}
