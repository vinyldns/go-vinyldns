/*
Copyright 2026 Comcast Cable Communications Management, LLC
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

import "testing"

func TestPing(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/ping",
			code:     200,
			body:     "PONG",
		},
	})
	defer server.Close()

	resp, err := client.Ping()
	if err != nil {
		t.Error(err)
	}
	if resp != "PONG" {
		t.Error("Expected ping response to be PONG")
	}
}

func TestHealth(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/health",
			code:     200,
			body:     "",
		},
	})
	defer server.Close()

	if err := client.Health(); err != nil {
		t.Error(err)
	}
}

func TestColor(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/color",
			code:     200,
			body:     "blue",
		},
	})
	defer server.Close()

	resp, err := client.Color()
	if err != nil {
		t.Error(err)
	}
	if resp != "blue" {
		t.Error("Expected color response to be blue")
	}
}

func TestMetricsPrometheus(t *testing.T) {
	server, client := testTools([]testToolsConfig{
		{
			endpoint: "http://host.com/metrics/prometheus?name=jvm_memory_bytes_used",
			code:     200,
			body:     "# HELP test\n",
		},
	})
	defer server.Close()

	resp, err := client.MetricsPrometheus([]string{"jvm_memory_bytes_used"})
	if err != nil {
		t.Error(err)
	}
	if resp == "" {
		t.Error("Expected metrics response to have a value")
	}
}
