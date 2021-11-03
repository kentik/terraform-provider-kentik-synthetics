package synthetics

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

func testToMap(obj *synthetics.V202101beta1Test) map[string]interface{} {
	if obj == nil {
		return nil
	}

	return map[string]interface{}{
		idKey:             obj.Id,
		"name":            obj.Name,
		"type":            obj.Type,
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
		"network_grid":        testNetworkGridToMapSlice(obj.NetworkGrid),
		"page_load":           testPageLoadToMapSlice(obj.PageLoad),
		"dns_grid":            testDNSGridToMapSlice(obj.DnsGrid),
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
		"rollup_level":        obj.RollupLevel,
		"http":                testHTTPToMapSlice(obj.Http),
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
		"inet_direction":                 obj.InetDirection,
		"direction":                      obj.Direction,
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
		"type":   obj.Type,
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

func testNetworkGridToMapSlice(obj *synthetics.V202101beta1GridTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"targets": obj.Targets,
	}}
}

func testPageLoadToMapSlice(obj *synthetics.V202101beta1UrlTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"target": obj.Target,
	}}
}

func testDNSGridToMapSlice(obj *synthetics.V202101beta1DnsGridTest) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		"targets": obj.Targets,
		"type":    obj.Type,
	}}
}

func testHealthSettingsToMapSlice(obj *synthetics.V202101beta1HealthSettings) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	// Kentik API sets 0 values for omitted optional fields.
	// Necessary conversion to nil, so Terraform configuration matches with actual state on the server.
	if areAllHSFieldsZeroValue(obj) {
		return nil
	}

	return []map[string]interface{}{{
		"latency_critical":             obj.LatencyCritical,
		"latency_warning":              obj.LatencyWarning,
		"packet_loss_critical":         obj.PacketLossCritical,
		"packet_loss_warning":          obj.PacketLossWarning,
		"jitter_critical":              obj.JitterCritical,
		"jitter_warning":               obj.JitterWarning,
		"http_latency_critical":        obj.HttpLatencyCritical,
		"http_latency_warning":         obj.HttpLatencyWarning,
		"http_valid_codes":             obj.HttpValidCodes,
		"dns_valid_codes":              obj.DnsValidCodes,
		"latency_critical_stddev":      obj.LatencyCriticalStddev,
		"latency_warning_stddev":       obj.LatencyWarningStddev,
		"jitter_critical_stddev":       obj.JitterCriticalStddev,
		"jitter_warning_stddev":        obj.JitterWarningStddev,
		"http_latency_critical_stddev": obj.HttpLatencyCriticalStddev,
		"http_latency_warning_stddev":  obj.HttpLatencyWarningStddev,
	}}
}

//nolint: gocyclo, gocognit
func areAllHSFieldsZeroValue(obj *synthetics.V202101beta1HealthSettings) bool {
	return obj.GetLatencyCritical() == 0 &&
		obj.GetLatencyWarning() == 0 &&
		obj.GetPacketLossCritical() == 0 &&
		obj.GetPacketLossWarning() == 0 &&
		obj.GetJitterCritical() == 0 &&
		obj.GetJitterWarning() == 0 &&
		obj.GetHttpLatencyCritical() == 0 &&
		obj.GetHttpLatencyWarning() == 0 &&
		len(obj.GetHttpValidCodes()) == 0 &&
		len(obj.GetDnsValidCodes()) == 0 &&
		obj.GetLatencyCriticalStddev() == 0 &&
		obj.GetLatencyWarningStddev() == 0 &&
		obj.GetJitterCriticalStddev() == 0 &&
		obj.GetJitterWarningStddev() == 0 &&
		obj.GetHttpLatencyCriticalStddev() == 0 &&
		obj.GetHttpLatencyWarningStddev() == 0
}

func testMonitoringSettingsToMapSlice(obj *synthetics.V202101beta1TestMonitoringSettings) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	// Kentik API sets 0 values for omitted optional fields.
	// Necessary conversion to nil, so Terraform configuration matches with actual state on the server.
	if len(obj.GetNotificationChannels()) == 0 {
		return nil
	}

	return []map[string]interface{}{{
		"notification_channels": obj.NotificationChannels,
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
		"delay":  obj.Delay,
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
		"delay":    obj.Delay,
	}}
}

func testHTTPToMapSlice(obj *synthetics.V202101beta1HTTPConfig) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	// Kentik API sets 0 values for omitted optional fields.
	// Necessary conversion to nil, so Terraform configuration matches with actual state on the server.
	if areAllHTTPFieldsZeroValue(obj) {
		return nil
	}

	return []map[string]interface{}{{
		"period":            obj.Period,
		"expiry":            obj.Expiry,
		"method":            obj.Method,
		"headers":           stringMapPtrToStringMap(obj.Headers),
		"body":              obj.Body,
		"ignore_tls_errors": obj.IgnoreTlsErrors,
		"css_selectors":     stringMapPtrToStringMap(obj.CssSelectors),
	}}
}

func areAllHTTPFieldsZeroValue(obj *synthetics.V202101beta1HTTPConfig) bool {
	return obj.GetPeriod() == 0 &&
		obj.GetExpiry() == 0 &&
		obj.GetMethod() == "" &&
		len(obj.GetHeaders()) == 0 &&
		obj.GetBody() == "" &&
		!obj.GetIgnoreTlsErrors() &&
		len(obj.GetCssSelectors()) == 0
}

func userInfoToMapSlice(obj *synthetics.V202101beta1UserInfo) []map[string]interface{} {
	if obj == nil {
		return nil
	}

	return []map[string]interface{}{{
		idKey:       obj.Id,
		"email":     obj.Email,
		"full_name": obj.FullName,
	}}
}

// resourceDataToTest converts TF resource data to Test object. Only user-writable attributes are set.
func resourceDataToTest(d *schema.ResourceData) (*synthetics.V202101beta1Test, error) {
	test := synthetics.NewV202101beta1Test()
	test.SetName(d.Get("name").(string))
	test.SetType(d.Get("type").(string))
	test.SetStatus(synthetics.V202101beta1TestStatus(d.Get("status").(string)))

	s, err := resourceDataToTestSettings(d.Get("settings"))
	if err != nil {
		return nil, err
	}
	test.Settings = s

	return test, nil
}

//nolint: gocyclo
func resourceDataToTestSettings(data interface{}) (*synthetics.V202101beta1TestSettings, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestSettings: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1TestSettings()

	h, err := resourceDataToTestHostname(m["hostname"])
	if err != nil {
		return nil, err
	}
	obj.Hostname = h

	ip, err := resourceDataToTestIP(m["ip"])
	if err != nil {
		return nil, err
	}
	obj.Ip = ip

	agent, err := resourceDataToTestAgent(m["agent"])
	if err != nil {
		return nil, err
	}
	obj.Agent = agent

	flow, err := resourceDataToTestFlow(m["flow"])
	if err != nil {
		return nil, err
	}
	obj.Flow = flow

	site, err := resourceDataToTestSite(m["site"])
	if err != nil {
		return nil, err
	}
	obj.Site = site

	tag, err := resourceDataToTestTag(m["tag"])
	if err != nil {
		return nil, err
	}
	obj.Tag = tag

	dns, err := resourceDataToTestDNS(m["dns"])
	if err != nil {
		return nil, err
	}
	obj.Dns = dns

	url, err := resourceDataToTestURL(m["url"])
	if err != nil {
		return nil, err
	}
	obj.Url = url

	ng, err := resourceDataToTestNetworkGrid(m["network_grid"])
	if err != nil {
		return nil, err
	}
	obj.NetworkGrid = ng

	pg, err := resourceDataToTestPageLoad(m["page_load"])
	if err != nil {
		return nil, err
	}
	obj.PageLoad = pg

	dg, err := resourceDataToTestDNSGrid(m["dns_grid"])
	if err != nil {
		return nil, err
	}
	obj.DnsGrid = dg

	obj.SetAgentIds(ifSliceToStringSlice(m["agent_ids"].([]interface{})))
	obj.SetPeriod(int64(m["period"].(int)))
	obj.SetCount(int64(m["count"].(int)))
	obj.SetExpiry(int64(m["expiry"].(int)))
	obj.SetLimit(int64(m["limit"].(int)))
	obj.SetTasks(ifSliceToStringSlice(m["tasks"].([]interface{})))

	healthSettings, err := resourceDataToTestHealthSettings(m["health_settings"])
	if err != nil {
		return nil, err
	}
	obj.HealthSettings = healthSettings

	monitoringSettings, err := resourceDataToTestMonitoringSettings(m["monitoring_settings"])
	if err != nil {
		return nil, err
	}
	obj.MonitoringSettings = monitoringSettings

	ping, err := resourceDataToTestPing(m["ping"])
	if err != nil {
		return nil, err
	}
	obj.Ping = ping

	trace, err := resourceDataToTestTrace(m["trace"])
	if err != nil {
		return nil, err
	}
	obj.Trace = trace

	obj.SetPort(int64(m["port"].(int)))
	obj.SetProtocol(m["protocol"].(string))

	if f, ok := m["family"]; ok && f != "" {
		obj.SetFamily(synthetics.V202101beta1IPFamily(f.(string)))
	}

	obj.SetServers(ifSliceToStringSlice(m["servers"].([]interface{})))
	obj.SetRollupLevel(int64(m["rollup_level"].(int)))

	http, err := resourceDataToTestHTTP(m["http"])
	if err != nil {
		return nil, err
	}
	obj.Http = http

	return obj, nil
}

func resourceDataToTestHostname(data interface{}) (*synthetics.V202101beta1HostnameTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestHostname: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1HostnameTest()
	obj.SetTarget(m["target"].(string))

	return obj, nil
}

func resourceDataToTestIP(data interface{}) (*synthetics.V202101beta1IpTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestIP: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1IpTest()
	obj.SetTargets(ifSliceToStringSlice(m["targets"].([]interface{})))

	return obj, nil
}

func resourceDataToTestAgent(data interface{}) (*synthetics.V202101beta1AgentTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestAgent: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1AgentTest()
	obj.SetTarget(m["target"].(string))

	return obj, nil
}

func resourceDataToTestFlow(data interface{}) (*synthetics.V202101beta1FlowTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestFlow: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1FlowTest()
	obj.SetTarget(m["target"].(string))
	obj.SetTargetRefreshIntervalMillis(int64(m["target_refresh_interval_millis"].(int)))
	obj.SetMaxTasks(int64(m["max_tasks"].(int)))
	obj.SetType(m["type"].(string))
	obj.SetInetDirection(m["inet_direction"].(string))
	obj.SetDirection(m["direction"].(string))

	return obj, nil
}

func resourceDataToTestSite(data interface{}) (*synthetics.V202101beta1SiteTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestSite: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1SiteTest()
	obj.SetTarget(m["target"].(string))

	return obj, nil
}

func resourceDataToTestTag(data interface{}) (*synthetics.V202101beta1TagTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestTag: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1TagTest()
	obj.SetTarget(m["target"].(string))

	return obj, nil
}

func resourceDataToTestDNS(data interface{}) (*synthetics.V202101beta1DnsTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestDNS: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1DnsTest()
	obj.SetTarget(m["target"].(string))
	obj.SetType(synthetics.V202101beta1DNSRecord(m["type"].(string)))

	return obj, nil
}

func resourceDataToTestURL(data interface{}) (*synthetics.V202101beta1UrlTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestURL: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1UrlTest()
	obj.SetTarget(m["target"].(string))

	return obj, nil
}

func resourceDataToTestNetworkGrid(data interface{}) (*synthetics.V202101beta1GridTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestNetworkGrid: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1GridTest()
	obj.SetTargets(ifSliceToStringSlice(m["targets"].([]interface{})))

	return obj, nil
}

func resourceDataToTestPageLoad(data interface{}) (*synthetics.V202101beta1UrlTest, error) {
	return resourceDataToTestURL(data)
}

func resourceDataToTestDNSGrid(data interface{}) (*synthetics.V202101beta1DnsGridTest, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestDNSGrid: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1DnsGridTest()
	obj.SetTargets(ifSliceToStringSlice(m["targets"].([]interface{})))
	obj.SetType(synthetics.V202101beta1DNSRecord(m["type"].(string)))

	return obj, nil
}

func resourceDataToTestHealthSettings(data interface{}) (*synthetics.V202101beta1HealthSettings, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestHealthSettings: %v", err)
	}

	obj := synthetics.NewV202101beta1HealthSettings()

	if m == nil {
		// Kentik API requires health_settings object
		return obj, nil
	}

	obj.SetLatencyCritical(float32(m["latency_critical"].(float64)))
	obj.SetLatencyWarning(float32(m["latency_warning"].(float64)))
	obj.SetPacketLossCritical(float32(m["packet_loss_critical"].(float64)))
	obj.SetPacketLossWarning(float32(m["packet_loss_warning"].(float64)))
	obj.SetJitterCritical(float32(m["jitter_critical"].(float64)))
	obj.SetJitterWarning(float32(m["jitter_warning"].(float64)))
	obj.SetHttpLatencyCritical(float32(m["http_latency_critical"].(float64)))
	obj.SetHttpLatencyWarning(float32(m["http_latency_warning"].(float64)))
	obj.SetHttpValidCodes(ifSliceToInt64Slice(m["http_valid_codes"].([]interface{})))
	obj.SetDnsValidCodes(ifSliceToInt64Slice(m["dns_valid_codes"].([]interface{})))
	obj.SetLatencyCriticalStddev(float32(m["latency_critical_stddev"].(float64)))
	obj.SetLatencyWarningStddev(float32(m["latency_warning_stddev"].(float64)))
	obj.SetJitterCriticalStddev(float32(m["jitter_critical_stddev"].(float64)))
	obj.SetJitterWarningStddev(float32(m["jitter_warning_stddev"].(float64)))
	obj.SetHttpLatencyCriticalStddev(float32(m["http_latency_critical_stddev"].(float64)))
	obj.SetHttpLatencyWarningStddev(float32(m["http_latency_warning_stddev"].(float64)))

	return obj, nil
}

func resourceDataToTestMonitoringSettings(data interface{}) (*synthetics.V202101beta1TestMonitoringSettings, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestMonitoringSettings: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1TestMonitoringSettings()
	obj.SetNotificationChannels(ifSliceToStringSlice(m["notification_channels"].([]interface{})))

	return obj, nil
}

func resourceDataToTestPing(data interface{}) (*synthetics.V202101beta1TestPingSettings, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestPing: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1TestPingSettings()
	obj.SetPeriod(float32(m["period"].(float64)))
	obj.SetCount(float32(m["count"].(float64)))
	obj.SetExpiry(float32(m["expiry"].(float64)))
	obj.SetDelay(float32(m["delay"].(float64)))

	return obj, nil
}

func resourceDataToTestTrace(data interface{}) (*synthetics.V202101beta1TestTraceSettings, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestTrace: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1TestTraceSettings()
	obj.SetPeriod(float32(m["period"].(float64)))
	obj.SetCount(float32(m["count"].(float64)))
	obj.SetProtocol(m["protocol"].(string))
	obj.SetPort(float32(m["port"].(float64)))
	obj.SetExpiry(float32(m["expiry"].(float64)))
	obj.SetLimit(float32(m["limit"].(float64)))
	obj.SetDelay(float32(m["delay"].(float64)))

	return obj, nil
}

func resourceDataToTestHTTP(data interface{}) (*synthetics.V202101beta1HTTPConfig, error) {
	m, err := getObjectFromNestedResourceData(data)
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestHTTP: %v", err)
	}
	if m == nil {
		return nil, nil
	}

	obj := synthetics.NewV202101beta1HTTPConfig()
	obj.SetPeriod(int64(m["period"].(int)))
	obj.SetExpiry(int64(m["expiry"].(int)))
	obj.SetMethod(m["method"].(string))

	h, err := stringInterfaceMapToStringMap(m["headers"].(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestHTTP (headers): %v", err)
	}
	obj.SetHeaders(h)

	obj.SetBody(m["body"].(string))
	obj.SetIgnoreTlsErrors(m["ignore_tls_errors"].(bool))

	css, err := stringInterfaceMapToStringMap(m["css_selectors"].(map[string]interface{}))
	if err != nil {
		return nil, fmt.Errorf("resourceDataToTestHTTP (css_selectors): %v", err)
	}
	obj.SetCssSelectors(css)

	return obj, nil
}
