/*
Copyright 2018-2026 Comcast Cable Communications Management, LLC
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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClientTLSVerification(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.Header.Get("Authorization"), "AWS4-HMAC-SHA256 ") {
			t.Error("expected an AWS SigV4 authorization header")
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"zones":[]}`)
	}))
	defer server.Close()

	c := NewClient(ClientConfiguration{
		AccessKey: "accessKey",
		SecretKey: "secretKey",
		Host:      server.URL,
	})
	if _, err := c.Zones(); err == nil {
		t.Fatal("expected the default client to reject an untrusted TLS certificate")
	}

	// This client trusts only the test server's certificate and still verifies TLS.
	c.HTTPClient = server.Client()
	if _, err := c.Zones(); err != nil {
		t.Fatalf("expected a signed request with a trusted TLS certificate to succeed: %v", err)
	}
}

func TestNewClientFromEnv(t *testing.T) {
	t.Setenv("VINYLDNS_ACCESS_KEY", "accessKey")
	t.Setenv("VINYLDNS_SECRET_KEY", "secretKey")
	t.Setenv("VINYLDNS_HOST", "https://vinyldns-api.com")
	t.Setenv("VINYLDNS_USER_AGENT", "")

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
	if client.UserAgent != fmt.Sprintf("go-vinyldns/%s", Version) {
		t.Error("NewClientFromEnv should set a default UserAgent if one is not present in the environment")
	}

}

func TestNewClientFromEnvWithUserAgent(t *testing.T) {
	t.Setenv("VINYLDNS_USER_AGENT", "foo")

	client := NewClientFromEnv()

	if client.UserAgent != "foo" {
		t.Error("NewClientFromEnv should set a UserAgent from the environment if one is present")
	}

}
