package synthetics_test

import (
	"testing"

	"github.com/kentik/terraform-provider-kentik-synthetics/synthetics"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	t.Parallel()
	err := synthetics.NewProvider().InternalValidate()
	assert.NoError(t, err)
}
