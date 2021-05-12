package synthetics

// Note the file is named "data_source_test_impl.go" instead of "data_source_test.go", so that
// Go toolchain does not treat it as test file (because of _test.go suffix)

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
)

func dataSourceTest() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing single synthetic test test",
		ReadContext: dataSourceTestRead,
		Schema:      makeTestSchema(readSingle),
	}
}

func dataSourceTestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	testID := d.Get("id").(string)
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceApi.TestGet(ctx, testID).Execute()
	if err != nil {
		return detailedDiagError("failed to read test", err, httpResp)
	}

	d.SetId(resp.Test.GetId())

	for k, v := range testToMap(resp.Test) {
		if err = d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func testToMap(obj *synthetics.V202101beta1Test) map[string]interface{} {
	if obj == nil {
		return nil
	}

	return map[string]interface{}{
		"id":              obj.Id,
		"name":            obj.Name,
		"type":            obj.Type,
		"device_id":       obj.DeviceId,
		"status":          obj.Status,
		"settings":        testSettingsToMapSlice(obj.Settings),
		"expires_on":      formatTime(obj.ExpiresOn),
		"cdate":           formatTime(obj.Cdate),
		"edate":           formatTime(obj.Edate),
		"created_by":      userInfoToMapSlice(obj.CreatedBy),
		"last_updated_by": userInfoToMapSlice(obj.LastUpdatedBy),
	}
}

func testSettingsToMapSlice(obj *synthetics.V202101beta1TestSettings) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"hostname":            testHostnameToMapSlice(obj.Hostname),
		"ip":                  testIPToMapSlice(obj.Ip),
		"agent":               testAgentToMapSlice(obj.Agent),
		"flow":                testFlowToMapSlice(obj.Flow),
		"site":                testSiteToMapSlice(obj.Site),
		"tag":                 testTagToMapSlice(obj.Tag),
		"dns":                 testDNSToMapSlice(obj.Dns),
		"url":                 testURLToMapSlice(obj.Url),
		"agent_ids":           obj.AgentIds,
		"period":              obj.Period,
		"count":               obj.Count,
		"expiry":              obj.Expiry,
		"limit":               obj.Limit,
		"tasks":               obj.Tasks,
		"health_settings":     testHealthSettingsToMapSlice(obj.HealthSettings),
		"monitoring_settings": testMonitoringSettingsToMapSlice(obj.MonitoringSettings),
		"ping":                testPingToMapSlice(obj.Ping),
		"trace":               testTraceToMapSlice(obj.Trace),
		"port":                obj.Port,
		"protocol":            obj.Protocol,
		"family":              obj.Family,
		"servers":             obj.Servers,
		"target_type":         obj.TargetType,
		"target_value":        obj.TargetValue,
		"use_local_ip":        obj.UseLocalIp,
		"reciprocal":          obj.Reciprocal,
		"rollup_level":        obj.RollupLevel,
	}}
}

func testHostnameToMapSlice(obj *synthetics.V202101beta1HostnameTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target": obj.Target,
	}}
}

func testIPToMapSlice(obj *synthetics.V202101beta1IpTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"targets": obj.Targets,
	}}
}

func testAgentToMapSlice(obj *synthetics.V202101beta1AgentTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target": obj.Target,
	}}
}

func testFlowToMapSlice(obj *synthetics.V202101beta1FlowTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target":                         obj.Target,
		"target_refresh_interval_millis": obj.TargetRefreshIntervalMillis,
		"max_tasks":                      obj.MaxTasks,
		"type":                           obj.Type,
	}}
}

func testSiteToMapSlice(obj *synthetics.V202101beta1SiteTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target": obj.Target,
	}}
}

func testTagToMapSlice(obj *synthetics.V202101beta1TagTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target": obj.Target,
	}}
}

func testDNSToMapSlice(obj *synthetics.V202101beta1DnsTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target": obj.Target,
	}}
}

func testURLToMapSlice(obj *synthetics.V202101beta1UrlTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target": obj.Target,
	}}
}

func testHealthSettingsToMapSlice(obj *synthetics.V202101beta1HealthSettings) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"latency_critical":      obj.LatencyCritical,
		"latency_warning":       obj.LatencyWarning,
		"packet_loss_critical":  obj.PacketLossCritical,
		"packet_loss_warning":   obj.PacketLossWarning,
		"jitter_critical":       obj.JitterCritical,
		"jitter_warning":        obj.JitterWarning,
		"http_latency_critical": obj.HttpLatencyCritical,
		"http_latency_warning":  obj.HttpLatencyWarning,
		"http_valid_codes":      obj.HttpValidCodes,
		"dns_valid_codes":       obj.DnsValidCodes,
	}}
}

func testMonitoringSettingsToMapSlice(obj *synthetics.V202101beta1TestMonitoringSettings) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"activation_grace_period": obj.ActivationGracePeriod,
		"activation_time_unit":    obj.ActivationTimeUnit,
		"activation_time_window":  obj.ActivationTimeWindow,
		"activation_times":        obj.ActivationTimes,
		"notification_channels":   obj.NotificationChannels,
	}}
}

func testPingToMapSlice(obj *synthetics.V202101beta1TestPingSettings) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"period": obj.Period,
		"count":  obj.Count,
		"expiry": obj.Expiry,
	}}
}

func testTraceToMapSlice(obj *synthetics.V202101beta1TestTraceSettings) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"period":   obj.Period,
		"count":    obj.Count,
		"protocol": obj.Protocol,
		"port":     obj.Port,
		"expiry":   obj.Expiry,
		"limit":    obj.Limit,
	}}
}

func userInfoToMapSlice(obj *synthetics.V202101beta1UserInfo) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"id":        obj.Id,
		"email":     obj.Email,
		"full_name": obj.FullName,
	}}
}
