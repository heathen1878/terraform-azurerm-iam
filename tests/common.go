package test

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
)

type Options map[string]any
type config struct {
	principal_id         string
	role_definition_name string
	scope                string
}

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

func GetTestConfig(t *testing.T) Options {
	t.Helper()

	return Options{
		"principal_id":         os.Getenv("ARM_CLIENT_OBJECT_ID"),
		"role_definition_name": "Reader",
		"scope":                os.Getenv("ARM_SUBSCRIPTION_ID"),
	}
}
