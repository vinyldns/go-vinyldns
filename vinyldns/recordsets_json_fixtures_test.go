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
	recordSetsJSON = `{
		"recordSets": [{
      "id": "6eb42765-818b-4966-89e8-1a34720c82a2",
      "zoneId": "2d9f4ec0-0596-4040-a953-d14e2cca8982",
      "name": "bind9",
      "type": "A",
      "status": "Active",
      "created": "2015-11-02T14:02:08Z",
      "updated": "2015-11-02T14:02:09Z",
      "ttl": 300,
      "records": [
        {
          "address": "127.0.0.1"
        }
      ],
      "account": "account-test-2"
    },
    {
      "id": "bba979b5-da30-4deb-9d8b-3f4a40e622c4",
      "zoneId": "2d9f4ec0-0596-4040-a953-d14e2cca8982",
      "name": "pjc-cname",
      "type": "CNAME",
      "status": "Active",
      "created": "2015-11-02T15:51:59Z",
      "updated": "2015-11-02T15:51:59Z",
      "ttl": 200,
      "records": [
        {
          "cname": "pjc-test."
        }
      ],
      "account": "vinyldns"
    }
  ]}`

	recordSetsListJSON1 = `{
		"nextId": "2",
		"maxItems": 1,
		"recordSets": [{
			"id": "1",
			"zoneId": "123",
			"name": "bind9",
			"type": "A",
			"status": "Active",
			"created": "2015-11-02T14:02:08Z",
			"updated": "2015-11-02T14:02:09Z",
			"ttl": 300,
			"records": [{
				"address": "127.0.0.1"
			}],
			"account": "account-test-2"
		}]
	}`

	recordSetsListJSON2 = `{
		"maxItems": 1,
		"recordSets": [{
			"id": "2",
			"zoneId": "123",
			"name": "bind9",
			"type": "A",
			"status": "Active",
			"created": "2015-11-02T14:02:08Z",
			"updated": "2015-11-02T14:02:09Z",
			"ttl": 300,
			"records": [{
				"address": "127.0.0.1"
			}],
			"account": "account-test-2"
		}]
	}`

	recordSetsListNoneJSON = `{
		"maxItems": 100,
		"recordSets": []
	}`

	recordSetUpdateResponseJSON = `{
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
		"recordSet": {
			"id": "e7739220-2a57-45ef-b69f-fdf42a262b87",
			"zoneId": "2d9f4ec0-0596-4040-a953-d14e2cca8982",
			"name": "foo.",
			"type": "A",
			"status": "Pending",
			"created": "2015-11-02T16:08:50Z",
			"ttl": 200,
			"records": [
				{
					"address": "127.0.0.1"
				}
			],
			"account": "vinyldns"
		},
		"userId": "pclear",
		"changeType": "Create",
		"status": "Pending",
		"created": "2015-11-02T16:08:50Z",
		"id": "b3d4e0a9-a081-4adc-9a95-3ec2e7d26635"
	}`

	recordSetJSON = `{
		"recordSet":{
			"id":"123",
			"ownerGroupId":"789",
			"zoneId":"456",
			"name":"test-01",
			"type":"A",
			"status":"Active",
			"created":"2015-11-02T13:41:54Z",
			"updated":"2015-11-02T13:41:57Z",
			"ttl":200,
			"records":[{
				"address":"127.0.0.1"
			}],
			"account":"vinyldns"
	}}`

	recordSetChangesListNoneJSON = `{
		"zoneId": "123",
		"maxItems": 100,
		"recordSetChanges": []
	}`

	recordSetChangesJSON1 = `{
		"zoneId": "123",
		"startFrom": "1",
		"nextId": "2",
		"maxItems": 1,
		"recordSetChanges": [{
			"status": "Complete",
			"zone": {
				"status": "Active",
				"updated": "2016-12-30T15:37:57Z",
				"name": "system-test-history.",
				"adminGroupId": "67b4da23-6832-4600-8450-9fa0664caeeb",
				"created": "2016-12-30T15:37:56Z",
				"account": "67b4da23-6832-4600-8450-9fa0664caeeb",
				"email": "i.changed.this.10.times@history-test.com",
				"shared": true,
				"acl": {
					"rules": []
				},
				"id": "9f353bc7-cb8d-491c-b074-34afafc97c5f"
			},
			"created": "2016-12-30T15:37:58Z",
			"recordSet": {
				"status": "Active",
				"updated": "2016-12-30T15:37:58Z",
				"name": "test-create-cname-ok-1",
				"created": "2016-12-30T15:37:57Z",
				"account": "history-id",
				"zoneId": "9f353bc7-cb8d-491c-b074-34afafc97c5f",
				"records": [
					{
						"cname": "changed-cname."
					}
				],
				"ttl": 200,
				"type": "CNAME",
				"id": "f62235df-5372-443c-9ba4-bdd3fca452f4"
			},
			"changeType": "Delete",
			"userId": "history-id",
			"updates": {
				"status": "Active",
				"updated": "2016-12-30T15:37:58Z",
				"name": "test-create-cname-ok",
				"created": "2016-12-30T15:37:57Z",
				"account": "history-id",
				"zoneId": "9f353bc7-cb8d-491c-b074-34afafc97c5f",
				"records": [{
					"cname": "changed-cname."
				}],
				"ttl": 200,
				"type": "CNAME",
				"id": "f62235df-5372-443c-9ba4-bdd3fca452f4"
			},
			"id": "1"
		}]
	}`

	recordSetChangesJSON2 = `{
		"zoneId": "123",
		"startFrom": "2",
		"maxItems": 1,
		"recordSetChanges": [{
			"status": "Complete",
			"zone": {
				"status": "Active",
				"updated": "2016-12-30T15:37:57Z",
				"name": "system-test-history.",
				"adminGroupId": "67b4da23-6832-4600-8450-9fa0664caeeb",
				"created": "2016-12-30T15:37:56Z",
				"account": "67b4da23-6832-4600-8450-9fa0664caeeb",
				"email": "i.changed.this.10.times@history-test.com",
				"shared": true,
				"acl": {
					"rules": []
				},
				"id": "9f353bc7-cb8d-491c-b074-34afafc97c5f"
			},
			"created": "2016-12-30T15:37:58Z",
			"recordSet": {
				"status": "Active",
				"updated": "2016-12-30T15:37:58Z",
				"name": "test-create-cname-ok-1",
				"created": "2016-12-30T15:37:57Z",
				"account": "history-id",
				"zoneId": "9f353bc7-cb8d-491c-b074-34afafc97c5f",
				"records": [
					{
						"cname": "changed-cname."
					}
				],
				"ttl": 200,
				"type": "CNAME",
				"id": "f62235df-5372-443c-9ba4-bdd3fca452f4"
			},
			"changeType": "Delete",
			"userId": "history-id",
			"updates": {
				"status": "Active",
				"updated": "2016-12-30T15:37:58Z",
				"name": "test-create-cname-ok",
				"created": "2016-12-30T15:37:57Z",
				"account": "history-id",
				"zoneId": "9f353bc7-cb8d-491c-b074-34afafc97c5f",
				"records": [{
					"cname": "changed-cname."
				}],
				"ttl": 200,
				"type": "CNAME",
				"id": "f62235df-5372-443c-9ba4-bdd3fca452f4"
			},
			"id": "2"
		}]
	}`

	recordSetChangeJSON = `{
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
		"recordSet": {
			"id": "e7739220-2a57-45ef-b69f-fdf42a262b87",
			"zoneId": "2d9f4ec0-0596-4040-a953-d14e2cca8982",
			"name": "foo.",
			"type": "A",
			"status": "Pending",
			"created": "2015-11-02T16:08:50Z",
			"ttl": 200,
			"records": [
				{
					"address": "127.0.0.1"
				}
			],
			"account": "vinyldns"
		},
		"userId": "pclear",
		"changeType": "Create",
		"status": "Pending",
		"created": "2015-11-02T16:08:50Z",
		"id": "b3d4e0a9-a081-4adc-9a95-3ec2e7d26635"
	}`
)
