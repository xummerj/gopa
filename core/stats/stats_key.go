/*
Copyright 2016 Medcl (m AT medcl.net)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stats

const STATS_FETCH_COUNT = "stats.count.fetch.total"
const STATS_FETCH_SUCCESS_COUNT = "stats.count.fetch.success"
const STATS_FETCH_FAIL_COUNT = "stats.count.fetch.fail"
const STATS_FETCH_TIMEOUT_COUNT = "stats.count.fetch.timeout"
const STATS_FETCH_IGNORE_COUNT = "stats.count.fetch.ignore"
const STATS_PARSE_COUNT = "stats.count.parse.total"
const STATS_PARSE_SUCCESS_COUNT = "stats.count.parse.success"
const STATS_PARSE_FAIL_COUNT = "stats.count.parse.fail"
const STATS_PARSE_IGNORE_COUNT = "stats.count.ignore.fail"


type StatsCount struct {
	TotalCount   int `json:"total,omitempty"`
	SuccessCount int `json:"success,omitempty"`
	FailCount    int `json:"fail,omitempty"`
	Ignore       int `json:"ignore,omitempty"`
	Timeout      int `json:"timeout,omitempty"`
}

type TaskStatus struct {
	Fetch StatsCount `json:"fetch"`
	Parse StatsCount `json:"parse"`
}
