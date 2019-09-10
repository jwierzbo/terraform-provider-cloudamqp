output "console_api_key" {
  value = "${cloudamqp_instance.generic_instance.apikey}"
  description = "API_KEY to access created Instance through Console API"
  sensitive   = true
}
