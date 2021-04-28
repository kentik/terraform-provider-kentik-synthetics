package synthetics

import "testing"

func TestProvider(t *testing.T) {
	t.Parallel()
	if err := NewProvider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}
