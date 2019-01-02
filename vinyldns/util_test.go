/*
Copyright 2019 Comcast Cable Communications Management, LLC
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

func TestConcat(t *testing.T) {
	res := concat([]string{"I ", "am ", "a ", "sentence"})

	if res != "I am a sentence" {
		t.Error("concat failed")
	}
}

func TestConcatStrs(t *testing.T) {
	res := concatStrs("", "I ", "am ", "a ", "sentence")

	if res != "I am a sentence" {
		t.Error("concatStrs failed")
	}
}

func TestConcatStrsWithDelimiter(t *testing.T) {
	res := concatStrs(", ", "I", "am", "a", "list")

	if res != "I, am, a, list" {
		t.Error("concatStrs with delimiter failed")
	}
}
