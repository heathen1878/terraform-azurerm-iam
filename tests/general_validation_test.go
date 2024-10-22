package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

func TestItErrorsWhenScopeIsEmpty(t *testing.T) {
	t.Parallel()

	//opts := DefaultOptions().Without("scope")
	opts := DefaultOptions()
	terraformOptions := Setup(t, "examples/standalone", opts)

	_, err := terraform.InitAndPlanE(t, terraformOptions)
	require.NotNil(t, err)
}

func TestItErrorsWhenPrincipalIdIsEmpty(t *testing.T) {
	t.Parallel()

	opts := DefaultOptions().Without("principal_id")

	terraformOptions := Setup(t, "examples/standalone", opts)

	_, err := terraform.InitAndPlanE(t, terraformOptions)
	require.NotNil(t, err)
}

func TestItErrorsWhenRoleDefintionNameIsEmpty(t *testing.T) {
	t.Parallel()

	opts := DefaultOptions().Without("role_definition_name")

	terraformOptions := Setup(t, "examples/standalone", opts)

	_, err := terraform.InitAndPlanE(t, terraformOptions)
	require.NotNil(t, err)
}
