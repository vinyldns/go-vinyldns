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
	"os"
	"testing"
)

func TestNewClientFromEnv(t *testing.T) {
	os.Setenv("VINYLDNS_ACCESS_KEY", "accessKey")
	os.Setenv("VINYLDNS_SECRET_KEY", "secretKey")
	os.Setenv("VINYLDNS_HOST", "https://vinyldns-api.com")

	client := NewClientFromEnv()

	if client.AccessKey != "accessKey" {
		t.Error("NewClientFromEnv should set an AccessKey from the environment")
	}
	if client.SecretKey != "secretKey" {
		t.Error("NewClientFromEnv should set a SecretKey from the environment")
	}
	if client.Host != "https://vinyldns-api.com" {
		t.Error("NewClientFromEnv should set a Host from the environment")
	}
	if client.UserAgent != "go-vinyldns" {
		t.Error("NewClientFromEnv should set a default UserAgent if one is not present in the environment")
	}

	os.Setenv("VINYLDNS_ACCESS_KEY", "")
	os.Setenv("VINYLDNS_SECRET_KEY", "")
	os.Setenv("VINYLDNS_HOST", "")
}

func TestNewClientFromEnvWithUserAgent(t *testing.T) {
	os.Setenv("VINYLDNS_USER_AGENT", "foo")

	client := NewClientFromEnv()

	if client.UserAgent != "foo" {
		t.Error("NewClientFromEnv should set a UserAgent from the environment if one is present")
	}

	os.Setenv("VINYLDNS_USER_AGENT", "")
}
