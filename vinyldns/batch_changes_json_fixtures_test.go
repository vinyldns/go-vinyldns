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

const (
	batchRecordChangesJSON = `{
		"batchChanges": [{
			"userName": "vinyl201",
			"status": "Complete",
			"totalChanges": 5,
			"userId": "vinyl",
			"comments": "this is optional",
			"createdTimestamp": "2018-05-11T18:12:13Z",
			"id": "bd03175c-6fd7-419e-991c-3d5d1441d995"
		},
		{
			"userName": "vinyl201",
			"status": "Complete",
			"totalChanges": 10,
			"userId": "vinyl",
			"comments": "this is optional",
			"createdTimestamp": "2018-05-11T18:12:12Z",
			"id": "c2ad84b0-e6de-4a70-aa28-e808d33deaa5"
		},
		{
			"userName": "vinyl201",
			"status": "Complete",
			"totalChanges": 7,
			"userId": "vinyl",
			"comments": "this is optional",
			"createdTimestamp": "2018-05-11T18:12:12Z",
			"id": "2b827a33-7c4f-4623-8dd9-277c6fba0e54"
		}]
	}`

	batchRecordChangeCreateJSON = `{
		"comments": "this is optional",
		"changes": [
			{
				"changeType": "Add",
				"inputName": "example.com.",
				"type": "A",
				"ttl": 3600,
				"record": {
					"address": "127.0.0.1"
				}
			}
		]
	}`

	batchRecordChangeJSON = `{
		"userName": "vinyl201",
		"status": "Pending",
		"userId": "vinyl",
		"comments": "this is optional",
		"createdTimestamp": "2018-05-09T14:19:34Z",
		"changes": [
			{
				"status": "Pending",
				"changeType": "Add",
				"ttl": 200,
				"recordName": "parent.com.",
				"type": "A",
				"id": "7573ca11-3e30-45a8-9ba5-791f7d6ae7a7",
				"zoneName": "parent.com.",
				"inputName": "parent.com.",
				"record": {
					"address": "127.0.0.1"
				},
				"zoneId": "74e93bfc-7296-4b86-83d3-1ffcb0eb3d13"
			}
		],
		"id": "02bd95f4-a32c-443b-82eb-54dbaa55b31a"
	}`
)
