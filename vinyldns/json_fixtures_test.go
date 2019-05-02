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
	zonesJSON = `{
		"zones": [{
			"name": "vinyldns.",
			"email": "some_user@foo.com",
			"status": "Active",
			"created": "2015-10-30T01:25:46Z",
			"id": "8f922062-25f2-4a9d-b5ed-d9368f32bd29",
			"connection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"transferConnection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7",
			"acl": {
				"rules": [
					{
						"accessLevel": "Read",
						"description": "test-acl-group-id",
						"groupId": "123",
						"recordMask": "www-*",
						"recordTypes": ["A", "AAAA", "CNAME"]
					}
				]
			}
		},{
			"name": "vinyldnstest.sys.vinyldns.net.",
			"email": "another_user@foo.com",
			"status": "Active",
			"created": "2015-10-30T22:47:38Z",
			"id": "2d9f4ef0-0596-4040-a953-d14e2cca8982",
			"connection": {
				"name": "vinyldnstest.sys.vinyldns.net.",
				"keyName": "all.vinyldns.com",
				"key": "OBF:1:QRS==",
				"primaryServer": "int-ttns01.resource.vinyldns.net"
			},
			"transferConnection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7",
			"acl": {
				"rules": [
					{
						"accessLevel": "Read",
						"description": "test-acl-group-id",
						"groupId": "123",
						"recordMask": "www-*",
						"recordTypes": ["A", "AAAA", "CNAME"]
					}
				]
			}
		}]}`

	zonesListJSON1 = `{
		"nextId": "2",
		"maxItems": 1,
		"zones": [{
			"name": "vinyldns-one.",
			"email": "some_user@foo.com",
			"status": "Active",
			"created": "2015-10-30T01:25:46Z",
			"id": "1",
			"connection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"transferConnection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7",
			"acl": {
				"rules": [
					{
						"accessLevel": "Read",
						"description": "test-acl-group-id",
						"groupId": "123",
						"recordMask": "www-*",
						"recordTypes": ["A", "AAAA", "CNAME"]
					}
				]
			}
		}]}`

	zonesListJSON2 = `{
		"startFrom": "2",
		"maxItems": 1,
		"zones": [{
			"name": "vinyldns-two.",
			"email": "some_user@foo.com",
			"status": "Active",
			"created": "2015-10-30T01:25:46Z",
			"id": "2",
			"connection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"transferConnection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7",
			"acl": {
				"rules": [
					{
						"accessLevel": "Read",
						"description": "test-acl-group-id",
						"groupId": "123",
						"recordMask": "www-*",
						"recordTypes": ["A", "AAAA", "CNAME"]
					}
				]
			}
		}]}`

	zonesListNoneJSON = `{
		"maxItems": 100,
		"zones": []
	}`

	zoneJSON = `{
		"zone":{
			"name":"vinyldns.",
			"email":"some_user@foo.com",
			"status":"Active",
			"created":"2015-10-30T01:25:46Z",
			"updated":"2015-10-30T01:25:46Z",
			"latestSync":"2015-10-30T01:25:46Z",
			"id":"123",
			"connection":{
				"name":"vinyldns.",
				"keyName":"vinyldns.",
				"key":"OBF:1:ABC",
				"primaryServer":"127.0.0.1"
			},
			"transferConnection": {
				"name": "vinyldns.",
				"keyName": "vinyldns.",
				"key": "OBF:1:ABC+5",
				"primaryServer": "127.0.0.1"
			},
			"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7",
			"acl": {
				"rules": [
					{
						"accessLevel": "Read",
						"description": "test-acl-group-id",
						"groupId": "123",
						"recordMask": "www-*",
						"recordTypes": ["A", "AAAA", "CNAME"]
					}
				]
			}
		}
	}`

	zoneUpdateResponseJSON = `{
		"zone": {
			"name": "test.",
			"email": "paul_cleary@foo.com",
			"status": "Active",
			"created": "2015-11-02T15:25:29Z",
			"id": "beb024cb-d31a-4fb3-bf9c-f08bf378d404",
			"adminGroupId": "c314836d-17db-4a57-b849-eb1feffe0ae7"
		},
		"userId": "pclear",
		"changeType": "Update",
		"status": "Complete",
		"created": "2015-11-02T15:25:29Z",
		"id": "ccf116b8-f72b-4507-b042-3c6cc64c58fd"
	}`

	zoneHistoryJSON = `{
		"zoneId": "123",
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
				"id": "change123"
			}
		],
		"recordSetChanges": [
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
				"recordSet": {
					"id": "rs123",
					"zoneId": "2d9f4ec0-0596-4040-a953-d14e2cca8982",
					"name": "bind9",
					"type": "A",
					"status": "Active",
					"created": "2015-11-02T13:59:48Z",
					"updated": "2015-11-02T13:59:51Z",
					"ttl": 300,
					"records": [
						{
							"address": "127.0.0.1"
						}
					],
					"account": "account-test-2"
				},
				"userId": "account",
				"changeType": "Create",
				"status": "Complete",
				"created": "2015-11-02T13:59:48Z",
				"id": "13c0f664-58c2-4b1a-9c46-086c3658f861"
			}
		]
	}`

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
			"ownerGroupId":"456",
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

	groupsJSON = `{
		"maxItems": 100,
		"groups": [
			{
				"id": "93887728-2b26-4749-ba69-98871dda9cc0",
				"name": "some-other-group",
				"email": "test@vinyldns.com",
				"created": "2017-03-02T16:23:07Z",
				"status": "Active",
				"members": [
					{
						"id": "2764183c-5e75-4ae6-8833-503cd5f4dcb0"
					}
				],
				"admins": [
					{
						"id": "2764183c-5e75-4ae6-8833-503cd5f4dcb0"
					}
				]
			},
			{
				"id": "aa1ea217-70a7-4350-b22b-c7e2f2158fb9",
				"name": "some-group",
				"email": "test@vinyldns.com",
				"created": "2017-03-02T16:22:57Z",
				"status": "Active",
				"members": [
					{
						"id": "2764183c-5e75-4ae6-8833-503cd5f4dcb0"
					}
				],
				"admins": [
					{
						"id": "2764183c-5e75-4ae6-8833-503cd5f4dcb0"
					}
				]
			}
		]
	}`

	groupJSON = `{
		"id":"3094e07b-3b82-4fff-ac9d-f60dff223c2c",
		"name":"test-group",
		"email":"test@test.com",
		"description":"this is a description",
		"created":"2017-01-18T14:39:26Z",
		"status":"Active",
		"members":[{"id":"test-id"}],
		"admins":[{"id":"test-id"}]
	}`

	groupAdminsJSON = `{
		"admins": [
			{
				"userName": "jdoe201",
				"firstName": "john",
				"created": "2017-03-02T16:39:02Z",
				"lastName": "doe",
				"email": "john_doe@vinyldns.com",
				"id": "2764183c-5e75-4ae6-8833-503cd5f4dcb0"
			},
			{
				"userName": "jdoe202",
				"firstName": "jane",
				"created": "2017-03-02T16:50:02Z",
				"lastName": "doe",
				"email": "jane_doe@vinyldns.com",
				"id": "1764183c-5e75-4ae6-8833-503cd5f4dcb4"
			}
		]
	}`

	groupMembersJSON = `{
		"members": [
			{
				"userName": "jdoe201",
				"firstName": "john",
				"created": "2017-03-02T16:39:02Z",
				"lastName": "doe",
				"email": "john_doe@vinyldns.com",
				"id": "2764183c-5e75-4ae6-8833-503cd5f4dcb0"
			},
			{
				"userName": "jdoe202",
				"firstName": "jane",
				"created": "2017-03-02T16:50:02Z",
				"lastName": "doe",
				"email": "jane_doe@vinyldns.com",
				"id": "1764183c-5e75-4ae6-8833-503cd5f4dcb4"
			}
		]
	}`

	groupActivityJSON = `{
		"maxItems": 100,
		"changes": [
			{
				"newGroup": {
					"status": "Active",
					"name": "test-list-group-activity-max-item-success",
					"created": "2017-03-02T18:49:58Z",
					"id": "1555bac7-0343-4d11-800f-955afb481818",
					"admins": [
						{
							"id": "ok"
						}
					],
					"members": [
						{
							"id": "dummy199"
						},
						{
							"id": "ok"
						}
					],
					"email": "test@test.com"
				},
				"created": "1488480605378",
				"userId": "some-user",
				"changeType": "Update",
				"oldGroup": {
					"status": "Active",
					"name": "test-list-group-activity-max-item-success",
					"created": "2017-03-02T18:49:58Z",
					"id": "1555bac7-0343-4d11-800f-955afb481818",
					"admins": [
						{
							"id": "ok"
						}
					],
					"members": [
						{
							"id": "dummy198"
						},
						{
							"id": "ok"
						}
					],
					"email": "test@test.com"
				},
				"id": "11abb88b-c47d-469b-bc2d-6656e00711cf"
			}
		]
	}`

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
