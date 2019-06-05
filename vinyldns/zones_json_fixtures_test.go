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
	zoneChangesListJSON1 = `{
		"zoneId": "123",
		"nextId": "2",
		"maxItems": 1,
		"zoneChanges": [
			{
				"zone": {
					"name": "vinyldnstest.sys.vinyldns.net.",
					"email": "paul_cleary@foo.com",
					"status": "Active",
					"created": "2015-10-30T22:47:38Z",
					"id": "2d9f4ec0-0596-4040-a953-d14e2cca8982",
					"connection": {
						"name": "vinyldnstest.sys.vinyldns.net.",
						"keyName": "cap_all.vinyldns.com",
						"key": "xxx",
						"primaryServer": "int-ddns01.resource.vinyldns.net"
					},
					"transferConnection": {
						"name": "vinyldns.",
						"keyName": "vinyldns.",
						"key": "OBF:1:ABC+5",
						"primaryServer": "127.0.0.1"
					},
					"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7",
					"acl": {
						"rules": []
					}
				},
				"userId": "userId1",
				"changeType": "Create",
				"status": "Complete",
				"created": "2015-10-30T22:47:38Z",
				"id": "1"
			}
		]
	}`

	zoneChangesListJSON2 = `{
		"zoneId": "123",
		"startFrom": "2",
		"zoneChanges": [
			{
				"zone": {
					"name": "vinyldnstest.sys.vinyldns.net.",
					"email": "paul_cleary@foo.com",
					"status": "Active",
					"created": "2015-10-30T22:47:38Z",
					"id": "2d9f4ec0-0596-4040-a953-d14e2cca8982",
					"connection": {
						"name": "vinyldnstest.sys.vinyldns.net.",
						"keyName": "cap_all.vinyldns.com",
						"key": "xxx",
						"primaryServer": "int-ddns01.resource.vinyldns.net"
					},
					"transferConnection": {
						"name": "vinyldns.",
						"keyName": "vinyldns.",
						"key": "OBF:1:ABC+5",
						"primaryServer": "127.0.0.1"
					},
					"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7",
					"acl": {
						"rules": []
					}
				},
				"userId": "userId1",
				"changeType": "Create",
				"status": "Complete",
				"created": "2015-10-30T22:47:38Z",
				"id": "2"
			}
		]
	}`
)
