package synthetics

import "testing"

func TestProvider(t *testing.T) {
	if err := NewProvider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}
