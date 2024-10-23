package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestItApplies(t *testing.T) {
	t.Parallel()

	opts := GetTestConfig(t)

	terraformOptions := Setup(t, "examples/standalone", opts)

	defer terraform.Destroy(t, terraformOptions)
	_, err := terraform.InitAndApplyE(t, terraformOptions)

	iam := map[string]any{}
	terraform.OutputStruct(t, terraformOptions, "iam", &iam)

	t.Log(iam["iam"].(map[string]any)["role_definition_name"]) // grab the value from a nested map...
	t.Log(opts["role_definition_name"])                        // grab the value passed to terraform plan...

	require.Nil(t, err)
	require.Equal(t, opts["role_definition_name"], iam["iam"].(map[string]any)["role_definition_name"])
}
