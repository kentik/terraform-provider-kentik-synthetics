package synthetics

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func detailedDiagError(summary string, err error, r *http.Response) diag.Diagnostics {
	return diag.Diagnostics{diag.Diagnostic{
		Severity: diag.Error,
		Summary:  summary,
		Detail:   detail(r, err),
	}}
}

func detail(r *http.Response, err error) string {
	if r != nil {
		return fmt.Sprintf("error: %v, status: %v, body: %v", err.Error(), r.Status, r.Body)
	}
	return err.Error()
}
