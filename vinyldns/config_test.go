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

func TestNewConfigFromEnv(t *testing.T) {
	accessKey := envOrDefault("VINYLDNS_ACCESS_KEY", "accesskey123")
	secretKey := envOrDefault("VINYLDNS_SECRET_KEY", "secretkey123")
	host := envOrDefault("VINYLDNS_HOST", "host.name.com")

	defaultConfig := NewConfigFromEnv()

	expectSame(t, defaultConfig.AccessKey, accessKey, "defaultConfig.AccessKey")
	expectSame(t, defaultConfig.SecretKey, secretKey, "defaultConfig.SecretKey")
	expectSame(t, defaultConfig.Host, host, "defaultConfig.Host")
}

func TestNewClient(t *testing.T) {
	const (
		testAccessKey = "access granted"
		testSecretKey = "this is very secret"
		testHost      = "certainly.a.unique.host"
	)

	client := NewClient(ClientConfiguration{
		AccessKey: testAccessKey,
		SecretKey: testSecretKey,
		Host:      testHost,
	})

	expectSame(t, client.AccessKey, testAccessKey, "client.AccessKey")
	expectSame(t, client.SecretKey, testSecretKey, "client.SecretKey")
	expectSame(t, client.Host, testHost, "client.Host")
}

func envOrDefault(envVar, defaultValue string) string {
	if value, has := os.LookupEnv(envVar); has {
		return value
	}
	os.Setenv(envVar, defaultValue)
	return envOrDefault(envVar, defaultValue)
}

func expectSame(t *testing.T, got, expected interface{}, name string) {
	if got != expected {
		t.Errorf("Expected %s to be equal to \"%v\"; got: \"%s\"", name, expected, got)
	}
}
