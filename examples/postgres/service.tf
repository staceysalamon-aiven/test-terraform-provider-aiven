# European Postgres Service
resource "aiven_pg" "example_postgres" {
  project      = var.project_name
  service_name = "example-postgres-service"
  cloud_name   = "aws-eu-west-2" 
  plan         = "startup-4"
}
