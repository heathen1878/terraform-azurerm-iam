package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type Options map[string]any
type config struct{}

func DefaultOptions() Options {
	return Options{
		"principal_id":         "00000000-0000-0000-0000-000000000000",
		"role_definition_name": "Reader",
		"scope":                "/subscriptions/00000000-0000-0000-0000-000000000000",
	}
}

func (o Options) With(with Options) Options {
	options := o
	for k, v := range with {
		options[k] = v
	}
	return options
}

func (o Options) Without(key string) Options {
	option := o
	delete(option, key)
	return option
}

func Setup(t *testing.T, e string, opts Options) *terraform.Options {
	tempFolder := test_structure.CopyTerraformFolderToTemp(t, "../", e)
	return &terraform.Options{
		TerraformDir: tempFolder,
		Vars:         opts,
	}
}

func GetTestConfig(t *testing.T) *config {
	t.Helper()

	config := config{}

	return &config
}
