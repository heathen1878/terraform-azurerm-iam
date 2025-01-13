package test

import (
	
)

func TestItApplies_Regression_Part_1(t *testing.T) {
	t.Parallel()

	opts := GetTestConfig(t)

	terraformOptions := Setup(t, "examples/standalone", opts)

	_, err := terraform.InitAndApplyE(t, terraformOptions)
	require.Nil(t, err)
}

func TestItApplies_Regression_Part_2(t *testing.T) {
	t.Parallel()

	opts := GetTestConfig(t)

	terraformOptions := Setup(t, "examples/standalone", opts)

	defer terraform.Destroy(t, terraformOptions)
	applyOutput, err := terraform.ApplyAndIdempotent(t, terraformOptions)

	// Check if apply output contains indications of resource changes
	hasChanges := strings.Contains(applyOutput, "Apply complete!")
	require.Nil(t, err)
	//require.Equal(t, "", "")
	assert.True(t, hasChanges, "Terraform apply made no changes.")
}