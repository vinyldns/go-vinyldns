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

import (
	"fmt"
	"strings"
)

func zonesEP(c *Client) string {
	return concatStrs("", c.Host, "/zones")
}

func pingEP(c *Client) string {
	return concatStrs("", c.Host, "/ping")
}

func healthEP(c *Client) string {
	return concatStrs("", c.Host, "/health")
}

func colorEP(c *Client) string {
	return concatStrs("", c.Host, "/color")
}

func prometheusMetricsEP(c *Client, names []string) string {
	base := concatStrs("", c.Host, "/metrics/prometheus")
	query := buildPrometheusQuery(names)
	return concatStrs("", base, query)
}

func statusEP(c *Client) string {
	return concatStrs("", c.Host, "/status")
}

func statusUpdateEP(c *Client, processingDisabled bool) string {
	return concatStrs("", statusEP(c), fmt.Sprintf("?processingDisabled=%t", processingDisabled))
}

func zonesListEP(c *Client, f ListFilter) string {
	query := buildQuery(f, "nameFilter")

	return concatStrs("", zonesEP(c), query)
}

func zoneEP(c *Client, id string) string {
	return concatStrs("", zonesEP(c), "/", id)
}

func zoneDetailsEP(c *Client, id string) string {
	return concatStrs("", zonesEP(c), "/", id, "/details")
}

func zoneBackendIDsEP(c *Client) string {
	return concatStrs("", zonesEP(c), "/backendids")
}

func zoneNameEP(c *Client, name string) string {
	return concatStrs("", zonesEP(c), "/name/", name)
}

func zoneChangesEP(c *Client, id string, f ListFilter) string {
	query := buildQuery(f, "nameFilter")

	return concatStrs("", zoneEP(c, id), "/changes", query)
}

func zoneDeletedChangesEP(c *Client, f DeletedZonesFilter) string {
	query := buildDeletedZonesQuery(f)
	return concatStrs("", zonesEP(c), "/deleted/changes", query)
}

func zoneACLRulesEP(c *Client, id string) string {
	return concatStrs("", zoneEP(c, id), "/acl/rules")
}

func zoneSyncEP(c *Client, id string) string {
	return concatStrs("", zoneEP(c, id), "/sync")
}

func recordSetsEP(c *Client, zoneID string) string {
	return concatStrs("", zoneEP(c, zoneID), "/recordsets")
}

func recordSetsListEP(c *Client, zoneID string, f ListFilter) string {
	query := buildQuery(f, "recordNameFilter")

	return concatStrs("", recordSetsEP(c, zoneID), query)
}

func recordSetsGlobalListEP(c *Client, f GlobalListFilter) string {
	query := buildGlobalListQuery(f)
	base := concatStrs("", c.Host, "/recordsets")

	return concatStrs("", base, query)
}

func recordSetEP(c *Client, zoneID, recordSetID string) string {
	return concatStrs("", recordSetsEP(c, zoneID), "/", recordSetID)
}

func recordSetCountEP(c *Client, zoneID string) string {
	return concatStrs("", zoneEP(c, zoneID), "/recordsetcount")
}

func recordSetChangesEP(c *Client, zoneID string, f ListFilterRecordSetChanges) string {
	query := buildRecordSetChangesQuery(f)

	return concatStrs("", zoneEP(c, zoneID), "/recordsetchanges", query)
}

func recordSetChangeEP(c *Client, zoneID, recordSetID, changeID string) string {
	return concatStrs("", recordSetEP(c, zoneID, recordSetID), "/changes/", changeID)
}

func recordSetChangeHistoryEP(c *Client, f RecordSetChangeHistoryFilter) string {
	query := buildRecordSetChangeHistoryQuery(f)
	return concatStrs("", c.Host, "/recordsetchange/history", query)
}

func groupsEP(c *Client) string {
	return concatStrs("", c.Host, "/groups")
}

func groupsListEP(c *Client, f ListFilter) string {
	query := buildQuery(f, "groupNameFilter")

	return concatStrs("", groupsEP(c), query)
}

func groupEP(c *Client, groupID string) string {
	return concatStrs("", groupsEP(c), "/", groupID)
}

func groupAdminsEP(c *Client, groupID string) string {
	return concatStrs("", groupEP(c, groupID), "/admins")
}

func groupMembersEP(c *Client, groupID string) string {
	return concatStrs("", groupEP(c, groupID), "/members")
}

func groupActivityEP(c *Client, groupID string) string {
	return concatStrs("", groupEP(c, groupID), "/activity")
}

func groupChangeEP(c *Client, groupChangeID string) string {
	return concatStrs("", groupsEP(c), "/change/", groupChangeID)
}

func groupValidDomainsEP(c *Client) string {
	return concatStrs("", groupsEP(c), "/valid/domains")
}

func usersEP(c *Client) string {
	return concatStrs("", c.Host, "/users")
}

func userEP(c *Client, userIdentifier string) string {
	return concatStrs("", usersEP(c), "/", userIdentifier)
}

func userLockEP(c *Client, userID string) string {
	return concatStrs("", userEP(c, userID), "/lock")
}

func userUnlockEP(c *Client, userID string) string {
	return concatStrs("", userEP(c, userID), "/unlock")
}

func batchRecordChangesEP(c *Client) string {
	return concatStrs("", zonesEP(c), "/batchrecordchanges")
}

func batchRecordChangeEP(c *Client, changeID string) string {
	return concatStrs("", batchRecordChangesEP(c), "/", changeID)
}

func batchRecordChangeApproveEP(c *Client, changeID string) string {
	return concatStrs("", batchRecordChangeEP(c, changeID), "/approve")
}

func batchRecordChangeRejectEP(c *Client, changeID string) string {
	return concatStrs("", batchRecordChangeEP(c, changeID), "/reject")
}

func batchRecordChangeCancelEP(c *Client, changeID string) string {
	return concatStrs("", batchRecordChangeEP(c, changeID), "/cancel")
}

func zoneChangesFailureEP(c *Client, f ListFilter) string {
	query := buildStartMaxQuery(f)
	return concatStrs("", c.Host, "/metrics/health/zonechangesfailure", query)
}

func recordSetChangesFailureEP(c *Client, zoneID string, f ListFilter) string {
	query := buildStartMaxQuery(f)
	return concatStrs("", c.Host, "/metrics/health/zones/", zoneID, "/recordsetchangesfailure", query)
}

func buildQuery(f ListFilter, nameFilterName string) string {
	params := []string{}
	query := "?"

	if f.NameFilter != "" {
		params = append(params, fmt.Sprintf("%s=%s", nameFilterName, f.NameFilter))
	}

	if f.StartFrom != "" {
		params = append(params, fmt.Sprintf("startFrom=%s", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	if len(params) == 0 {
		query = ""
	}

	return query + strings.Join(params, "&")
}

func buildStartMaxQuery(f ListFilter) string {
	params := []string{}
	query := "?"

	if f.StartFrom != "" {
		params = append(params, fmt.Sprintf("startFrom=%s", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	if len(params) == 0 {
		query = ""
	}

	return query + strings.Join(params, "&")
}

func buildDeletedZonesQuery(f DeletedZonesFilter) string {
	params := []string{}
	query := "?"

	if f.NameFilter != "" {
		params = append(params, fmt.Sprintf("nameFilter=%s", f.NameFilter))
	}

	if f.StartFrom != "" {
		params = append(params, fmt.Sprintf("startFrom=%s", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	if f.IgnoreAccess != nil {
		params = append(params, fmt.Sprintf("ignoreAccess=%t", *f.IgnoreAccess))
	}

	if len(params) == 0 {
		query = ""
	}

	return query + strings.Join(params, "&")
}

func buildRecordSetChangesQuery(f ListFilterRecordSetChanges) string {
	params := []string{}
	query := "?"

	if f.StartFrom != 0 {
		params = append(params, fmt.Sprintf("startFrom=%d", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	if len(params) == 0 {
		query = ""
	}

	return query + strings.Join(params, "&")
}

func buildRecordSetChangeHistoryQuery(f RecordSetChangeHistoryFilter) string {
	params := []string{}
	query := "?"

	params = append(params, fmt.Sprintf("zoneId=%s", f.ZoneID))
	params = append(params, fmt.Sprintf("fqdn=%s", f.FQDN))
	params = append(params, fmt.Sprintf("recordType=%s", f.RecordType))

	if f.StartFrom != "" {
		params = append(params, fmt.Sprintf("startFrom=%s", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	return query + strings.Join(params, "&")
}

func buildPrometheusQuery(names []string) string {
	params := []string{}
	query := "?"

	for _, name := range names {
		if name != "" {
			params = append(params, fmt.Sprintf("name=%s", name))
		}
	}

	if len(params) == 0 {
		query = ""
	}

	return query + strings.Join(params, "&")
}

func buildGlobalListQuery(f GlobalListFilter) string {
	params := []string{}
	query := "?"

	if f.RecordNameFilter != "" {
		params = append(params, fmt.Sprintf("%s=%s", "recordNameFilter", f.RecordNameFilter))
	}

	if f.RecordTypeFilter != "" {
		params = append(params, fmt.Sprintf("%s=%s", "recordTypeFilter", f.RecordTypeFilter))
	}

	if f.RecordOwnerGroupFilter != "" {
		params = append(params, fmt.Sprintf("%s=%s", "recordOwnerGroupFilter", f.RecordOwnerGroupFilter))
	}

	if f.NameSort != "" {
		params = append(params, fmt.Sprintf("%s=%s", "nameSort", f.NameSort))
	}

	if f.StartFrom != "" {
		params = append(params, fmt.Sprintf("startFrom=%s", f.StartFrom))
	}

	if f.MaxItems != 0 {
		params = append(params, fmt.Sprintf("maxItems=%d", f.MaxItems))
	}

	if len(params) == 0 {
		query = ""
	}

	return query + strings.Join(params, "&")
}
