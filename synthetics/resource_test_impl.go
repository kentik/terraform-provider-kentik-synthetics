package synthetics

import (
	"context"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi/synthetics"
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

	req := *synthetics.NewV202101beta1CreateTestRequest()
	req.SetTest(*test)

	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
		TestCreate(ctx).
		Body(req).
		Execute()
	if err != nil {
		return detailedDiagError("failed to create test", err, httpResp)
	}

	err = d.Set(idKey, resp.Test.GetId())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Test.GetId())

	// read back the just-created resource to handle the case when server applies modifications to provided data
	return resourceTestRead(ctx, d, m)
}

func resourceTestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resp, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
		TestGet(ctx, d.Get(idKey).(string)).
		Execute()
	if err != nil {
		if httpResp.StatusCode == http.StatusNotFound {
			d.SetId("")
			return nil
		}
		return detailedDiagError("failed to read test", err, httpResp)
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

		req := *synthetics.NewV202101beta1PatchTestRequest()
		req.SetTest(*test)
		req.SetMask(strings.Join(patchTestFields(test), ","))

		_, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
			TestPatch(ctx, d.Get(idKey).(string)).
			Body(req).
			Execute()
		if err != nil {
			return detailedDiagError("failed to patch test", err, httpResp)
		}
	}

	// read back the just-updated resource to handle the case when server applies modifications to provided data
	return resourceTestRead(ctx, d, m)
}

// patchTestFields returns updatable fields for PATCH operation.
// TODO(dfurman): use JSON struct tags.
func patchTestFields(test *synthetics.V202101beta1Test) []string {
	return append(
		commonPatchTestFields(),
		optionalPatchTestFields(test)...,
	)
}

func commonPatchTestFields() []string {
	return []string{
		"test.name",
		"test.type",
		"test.status",
		"test.settings.agentIds",
		"test.settings.period",
		"test.settings.count",
		"test.settings.expiry",
		"test.settings.limit",
		"test.settings.tasks",
		"test.settings.healthSettings.latencyCritical",
		"test.settings.healthSettings.latencyWarning",
		"test.settings.healthSettings.packetLossCritical",
		"test.settings.healthSettings.packetLossWarning",
		"test.settings.healthSettings.jitterCritical",
		"test.settings.healthSettings.jitterWarning",
		"test.settings.healthSettings.httpLatencyCritical",
		"test.settings.healthSettings.httpLatencyWarning",
		"test.settings.healthSettings.httpValidCodes",
		"test.settings.healthSettings.dnsValidCodes",
		"test.settings.monitoringSettings.activationGracePeriod",
		"test.settings.monitoringSettings.activationTimeUnit",
		"test.settings.monitoringSettings.activationTimeWindow",
		"test.settings.monitoringSettings.activationTimes",
		// FIXME: cannot create or update notification channels - API returns HTTP 500:
		// 500 Internal Server Error, body: {{"code":13,"message":"internal error (name:TypeError type:object)","details":[]}}
		// "test.settings.monitoringSettings.notificationChannels",
		"test.settings.ping.period",
		"test.settings.ping.count",
		"test.settings.ping.expiry",
		"test.settings.trace.period",
		"test.settings.trace.count",
		"test.settings.trace.protocol",
		"test.settings.trace.port",
		"test.settings.trace.expiry",
		"test.settings.trace.limit",
		"test.settings.port",
		"test.settings.protocol",
		"test.settings.family",
		"test.settings.servers",
		"test.settings.useLocalIp",
		"test.settings.reciprocal",
		"test.settings.rollupLevel",
	}
}

func optionalPatchTestFields(test *synthetics.V202101beta1Test) []string {
	var fields []string

	if test.Settings.HasHostname() {
		fields = append(fields, "test.settings.hostname.target")
	}

	if test.Settings.HasIp() {
		fields = append(fields, "test.settings.ip.targets")
	}

	if test.Settings.HasAgent() {
		fields = append(fields, "test.settings.agent.targets")
	}

	if test.Settings.HasFlow() {
		fields = append(
			fields,
			"test.settings.flow.target",
			"test.settings.flow.targetRefreshIntervalMillis",
			"test.settings.flow.maxTasks",
			"test.settings.flow.type",
		)
	}

	if test.Settings.HasSite() {
		fields = append(fields, "test.settings.site.target")
	}

	if test.Settings.HasTag() {
		fields = append(fields, "test.settings.tag.target")
	}

	if test.Settings.HasDns() {
		fields = append(fields, "test.settings.dns.target")
	}

	if test.Settings.HasUrl() {
		fields = append(fields, "test.settings.url.target")
	}

	return fields
}

func resourceTestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	_, httpResp, err := m.(*kentikapi.Client).SyntheticsAdminServiceAPI.
		TestDelete(ctx, d.Get(idKey).(string)).
		Execute()
	if err != nil {
		return detailedDiagError("failed to delete test", err, httpResp)
	}

	return nil
}
