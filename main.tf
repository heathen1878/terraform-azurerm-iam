resource "azurerm_role_assignment" "this" {
  description          = var.description
  principal_id         = var.principal_id
  principal_type       = var.principal_type
  role_definition_name = var.role_definition_name
  scope                = var.scope
}