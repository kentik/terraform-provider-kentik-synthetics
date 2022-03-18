package synthetics

import (
	"context"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/kentik/community_sdk_golang/kentikapi/synthetics"
)

// Note the file is named "resource_test_impl.go" instead of "resource_test.go", so that
// Go toolchain does not treat it as test file (because of _test.go suffix)

func resourceTest() *schema.Resource {
	return &schema.Resource{
		Description:   "Resource representing synthetic test",
		CreateContext: resourceTestCreate,
		ReadContext:   resourceTestRead,
		UpdateContext: resourceTestUpdate,
		DeleteContext: resourceTestDelete,
		Schema:        makeTestSchema(create),
	}
}

func resourceTestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	test, err := resourceDataToTest(d)
	if err != nil {
		return diag.FromErr(err)
	}

	setRequiredInternalTestAttributes(test)

	req := *synthetics.NewV202101beta1CreateTestRequest()
	req.Test = test
	tflog.Debug(ctx, "Create Test Kentik API request", map[string]interface{}{"request": req})
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
		TestCreate(ctx).
		Body(req).
		Execute()
	tflog.Debug(ctx, "Create Test Kentik API response", map[string]interface{}{"response": resp})
	if err != nil {
		return detailedDiagError("Failed to create test", err, httpResp)
	}

	err = d.Set(idKey, resp.Test.GetId())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Test.GetId()) // create the resource in TF state

	// read back the just-created resource to handle the case when server applies modifications to provided data
	return resourceTestRead(ctx, d, m)
}

func resourceTestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Get Test Kentik API request", map[string]interface{}{"ID": d.Get(idKey).(string)})
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
		TestGet(ctx, d.Get(idKey).(string)).
		Execute()
	tflog.Debug(ctx, "Get Test Kentik API response", map[string]interface{}{"response": resp})
	if err != nil {
		if httpResp.StatusCode == http.StatusNotFound {
			d.SetId("") // delete the resource in TF state
			return nil
		}
		return detailedDiagError("Failed to read test", err, httpResp)
	}

	for k, v := range testToMap(resp.Test) {
		if err = d.Set(k, v); err != nil {
			return diag.FromErr(err)
		}
	}

	return nil
}

func resourceTestUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// check if any attribute has changed
	if d.HasChange("") {
		test, err := resourceDataToTest(d)
		if err != nil {
			return diag.FromErr(err)
		}

		setRequiredInternalTestAttributes(test)

		req := *synthetics.NewV202101beta1PatchTestRequest()
		req.Test = test
		req.SetMask(strings.Join(patchTestFields(test), ","))
		tflog.Debug(ctx, "Update Test Kentik API request", map[string]interface{}{"request": req})

		_, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
			TestPatch(ctx, d.Get(idKey).(string)).
			Body(req).
			Execute()
		tflog.Debug(ctx, "Update Test Kentik API response", map[string]interface{}{"response": httpResp.Body})
		if err != nil {
			return detailedDiagError("Failed to update test", err, httpResp)
		}
	}

	// read back the just-updated resource to handle the case when server applies modifications to provided data
	return resourceTestRead(ctx, d, m)
}

func setRequiredInternalTestAttributes(test *synthetics.V202101beta1Test) {
	test.SetDeviceId("0")
	if test.Settings.MonitoringSettings == nil {
		test.Settings.MonitoringSettings = synthetics.NewV202101beta1TestMonitoringSettings()
	}
	test.Settings.MonitoringSettings.SetActivationTimeUnit("m")
	test.Settings.MonitoringSettings.SetActivationTimeWindow("5")
	test.Settings.MonitoringSettings.SetActivationTimes("3")
}

// patchTestFields returns all fields that are: non-nil, non-internal and updatable.
// TODO(dfurman): use JSON struct tags.
func patchTestFields(test *synthetics.V202101beta1Test) []string {
	var fields []string

	if test.HasName() {
		fields = append(fields, "test.name")
	}

	if test.HasStatus() {
		fields = append(fields, "test.status")
	}

	return append(fields, testSettingsFields(test.Settings)...)
}

//nolint: gocyclo
func testSettingsFields(ts *synthetics.V202101beta1TestSettings) []string {
	var fields []string

	if ts.HasHostname() {
		fields = append(fields, "test.settings.hostname.target")
	}

	if ts.HasIp() {
		fields = append(fields, "test.settings.ip.targets")
	}

	if ts.HasAgent() {
		fields = append(fields, "test.settings.agent.target")
	}

	if ts.HasFlow() {
		fields = append(
			fields,
			"test.settings.flow.target",
			"test.settings.flow.targetRefreshIntervalMillis",
			"test.settings.flow.maxTasks",
			"test.settings.flow.type",
			"test.settings.flow.inetDirection",
			"test.settings.flow.direction",
		)
	}

	if ts.HasSite() {
		fields = append(fields, "test.settings.site.target")
	}

	if ts.HasTag() {
		fields = append(fields, "test.settings.tag.target")
	}

	if ts.HasDns() {
		fields = append(fields, "test.settings.dns.target", "test.settings.dns.type")
	}

	if ts.HasUrl() {
		fields = append(fields, "test.settings.url.target")
	}

	if ts.HasNetworkGrid() {
		fields = append(fields, "test.settings.networkGrid.targets")
	}

	if ts.HasPageLoad() {
		fields = append(fields, "test.settings.pageLoad.target")
	}

	if ts.HasDnsGrid() {
		fields = append(fields, "test.settings.dnsGrid.targets", "test.settings.dnsGrid.type")
	}

	if ts.HasAgentIds() {
		fields = append(fields, "test.settings.agentIds")
	}

	if ts.HasPeriod() {
		fields = append(fields, "test.settings.period")
	}

	if ts.HasCount() {
		fields = append(fields, "test.settings.count")
	}

	if ts.HasExpiry() {
		fields = append(fields, "test.settings.expiry")
	}

	if ts.HasLimit() {
		fields = append(fields, "test.settings.limit")
	}

	if ts.HasTasks() {
		fields = append(fields, "test.settings.tasks")
	}

	fields = append(fields, healthSettingsFields(ts.HealthSettings)...)
	fields = append(fields, monitoringSettingsFields(ts.MonitoringSettings)...)
	fields = append(fields, pingSettingsFields(ts.Ping)...)
	fields = append(fields, traceSettingsFields(ts.Trace)...)

	if ts.HasPort() {
		fields = append(fields, "test.settings.port")
	}

	if ts.HasProtocol() {
		fields = append(fields, "test.settings.protocol")
	}

	if ts.HasFamily() {
		fields = append(fields, "test.settings.family")
	}

	// TODO(dfurman): remove list length check when API accepts such payload
	if ts.HasServers() && len(ts.GetServers()) != 0 {
		fields = append(fields, "test.settings.servers")
	}

	if ts.HasRollupLevel() {
		fields = append(fields, "test.settings.rollupLevel")
	}

	if ts.HasHttp() {
		fields = append(fields, httpSettingsFields(ts.Http)...)
	}

	return fields
}

//nolint: gocyclo
func healthSettingsFields(hs *synthetics.V202101beta1HealthSettings) []string {
	var fields []string

	if hs.HasLatencyCritical() {
		fields = append(fields, "test.settings.healthSettings.latencyCritical")
	}

	if hs.HasLatencyWarning() {
		fields = append(fields, "test.settings.healthSettings.latencyWarning")
	}

	if hs.HasPacketLossCritical() {
		fields = append(fields, "test.settings.healthSettings.packetLossCritical")
	}

	if hs.HasPacketLossWarning() {
		fields = append(fields, "test.settings.healthSettings.packetLossWarning")
	}

	if hs.HasJitterCritical() {
		fields = append(fields, "test.settings.healthSettings.jitterCritical")
	}

	if hs.HasJitterWarning() {
		fields = append(fields, "test.settings.healthSettings.jitterWarning")
	}

	if hs.HasHttpLatencyCritical() {
		fields = append(fields, "test.settings.healthSettings.httpLatencyCritical")
	}

	if hs.HasHttpLatencyWarning() {
		fields = append(fields, "test.settings.healthSettings.httpLatencyWarning")
	}

	// TODO(dfurman): remove list length check when API accepts such payload
	if hs.HasHttpValidCodes() && len(hs.GetHttpValidCodes()) != 0 {
		fields = append(fields, "test.settings.healthSettings.httpValidCodes")
	}

	// TODO(dfurman): remove list length check when API accepts such payload
	if hs.HasDnsValidCodes() && len(hs.GetDnsValidCodes()) != 0 {
		fields = append(fields, "test.settings.healthSettings.dnsValidCodes")
	}

	if hs.HasLatencyCriticalStddev() {
		fields = append(fields, "test.settings.healthSettings.latencyCriticalStddev")
	}

	if hs.HasLatencyWarningStddev() {
		fields = append(fields, "test.settings.healthSettings.latencyWarningStddev")
	}

	if hs.HasJitterCriticalStddev() {
		fields = append(fields, "test.settings.healthSettings.jitterCriticalStddev")
	}

	if hs.HasJitterWarningStddev() {
		fields = append(fields, "test.settings.healthSettings.jitterWarningStddev")
	}

	if hs.HasHttpLatencyCriticalStddev() {
		fields = append(fields, "test.settings.healthSettings.httpLatencyCriticalStddev")
	}

	if hs.HasHttpLatencyWarningStddev() {
		fields = append(fields, "test.settings.healthSettings.httpLatencyWarningStddev")
	}

	return fields
}

func monitoringSettingsFields(ms *synthetics.V202101beta1TestMonitoringSettings) []string {
	var fields []string

	// TODO(dfurman): remove list length check when API accepts such payload
	if ms.HasNotificationChannels() &&
		len(ms.GetNotificationChannels()) != 0 {
		fields = append(fields, "test.settings.monitoringSettings.notificationChannels")
	}

	return fields
}

func pingSettingsFields(ps *synthetics.V202101beta1TestPingSettings) []string {
	var fields []string

	if ps.HasPeriod() {
		fields = append(fields, "test.settings.ping.period")
	}

	if ps.HasCount() {
		fields = append(fields, "test.settings.ping.count")
	}

	if ps.HasExpiry() {
		fields = append(fields, "test.settings.ping.expiry")
	}

	if ps.HasDelay() {
		fields = append(fields, "test.settings.ping.delay")
	}

	return fields
}

func traceSettingsFields(ts *synthetics.V202101beta1TestTraceSettings) []string {
	var fields []string

	if ts.HasPeriod() {
		fields = append(fields, "test.settings.trace.period")
	}

	if ts.HasCount() {
		fields = append(fields, "test.settings.trace.count")
	}

	if ts.HasProtocol() {
		fields = append(fields, "test.settings.trace.protocol")
	}

	if ts.HasPort() {
		fields = append(fields, "test.settings.trace.port")
	}

	if ts.HasExpiry() {
		fields = append(fields, "test.settings.trace.expiry")
	}

	if ts.HasLimit() {
		fields = append(fields, "test.settings.trace.limit")
	}

	if ts.HasDelay() {
		fields = append(fields, "test.settings.trace.delay")
	}

	return fields
}

func httpSettingsFields(h *synthetics.V202101beta1HTTPConfig) []string {
	var fields []string

	if h.HasPeriod() {
		fields = append(fields, "test.settings.http.period")
	}

	if h.HasExpiry() {
		fields = append(fields, "test.settings.http.expiry")
	}

	if h.HasMethod() {
		fields = append(fields, "test.settings.http.method")
	}

	if h.HasHeaders() {
		fields = append(fields, "test.settings.http.headers")
	}

	if h.HasBody() {
		fields = append(fields, "test.settings.http.body")
	}

	if h.HasIgnoreTlsErrors() {
		fields = append(fields, "test.settings.http.ignoreTlsErrors")
	}

	if h.HasCssSelectors() {
		fields = append(fields, "test.settings.http.cssSelectors")
	}

	return fields
}

func resourceTestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Debug(ctx, "Delete Test Kentik API request", map[string]interface{}{"ID": d.Get(idKey).(string)})
	_, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
		TestDelete(ctx, d.Get(idKey).(string)).
		Execute()
	tflog.Debug(ctx, "Delete Test Kentik API response", map[string]interface{}{"response": httpResp.Body})
	if err != nil {
		return detailedDiagError("Failed to delete test", err, httpResp)
	}

	return nil
}
