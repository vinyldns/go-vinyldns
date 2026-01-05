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

// Ping performs a health check that returns "PONG".
func (c *Client) Ping() (string, error) {
	_, body, err := resourceRequestRaw(c, pingEP(c), "GET", nil)
	if err != nil {
		return "", err
	}

	return body, nil
}

// Health performs a comprehensive health check.
func (c *Client) Health() error {
	_, _, err := resourceRequestRaw(c, healthEP(c), "GET", nil)
	return err
}

// Color returns the current blue/green deployment color.
func (c *Client) Color() (string, error) {
	_, body, err := resourceRequestRaw(c, colorEP(c), "GET", nil)
	if err != nil {
		return "", err
	}

	return body, nil
}

// MetricsPrometheus returns metrics in Prometheus text format.
func (c *Client) MetricsPrometheus(names []string) (string, error) {
	_, body, err := resourceRequestRaw(c, prometheusMetricsEP(c, names), "GET", nil)
	if err != nil {
		return "", err
	}

	return body, nil
}
