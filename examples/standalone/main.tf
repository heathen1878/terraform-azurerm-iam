module "iam" {

  source = "../.."

  principal_id         = var.principal_id
  role_definition_name = var.role_definition_name
  scope                = var.scope
}