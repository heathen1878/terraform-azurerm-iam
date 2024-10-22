locals {
  location = "uksouth"
}

resource "azurerm_resource_group" "this" {
  name     = format("rg-%s", random_id.this.hex)
  location = local.location
}

module "iam" {

  source = "../.."

  principal_id         = data.azurerm_client_config.this.object_id
  role_definition_name = var.role_definition_name
  scope                = azurerm_resource_group.this.id
}