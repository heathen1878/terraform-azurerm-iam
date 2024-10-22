locals {
  location = "uksouth"
}

resource "azurerm_resource_group" "this" {
  name     = format("rg-%s", random_id.this.hex)
  location = local.location
}

module "iam" {

  source = "../.."

  principal_id         = var.principal_id
  role_definition_name = var.role_definition_name
  scope                = var.scope
}