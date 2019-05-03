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
)
