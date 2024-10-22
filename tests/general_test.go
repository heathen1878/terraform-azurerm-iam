package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestItApplies(t *testing.T) {
	t.Parallel()

	opts := DefaultOptions().With(Options{
		"role_definition_name": "Contributor",
	})

	terraformOptions := Setup(t, "examples/standalone", opts)
	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	iam := map[string]any{}
	terraform.OutputStruct(t, terraformOptions, "iam", &iam)

	require.Equal(t, "Contributor", iam["role_definition_name"])
}
