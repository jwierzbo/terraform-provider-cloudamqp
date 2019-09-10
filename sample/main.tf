module "test_instance_bunny" {
  source = "cloudamqp-instance"

  instance_name   = "test-instance"
  instance_plan   = "bunny"
  instance_region = "amazon-web-services::eu-west-1"
}

resource "cloudamqp_alarm_recipient" "test_recipient" {
  type            = "email"
  value           = "some_email@domain.com"
  console_api_key = "${module.test_instance_bunny.console_api_key}"
}

resource "cloudamqp_alarm_configuration" "test_alarm_cpu" {
  type            = "cpu"
  value_threshold = 90
  time_threshold = 120
  console_api_key = "${module.test_instance_bunny.console_api_key}"
}

resource "cloudamqp_alarm_configuration" "test_alarm_memory" {
  type            = "memory"
  value_threshold = 90
  time_threshold = 120
  console_api_key = "${module.test_instance_bunny.console_api_key}"
}

resource "cloudamqp_alarm_configuration" "test_alarm_disk" {
  type            = "disk"
  value_threshold = 5
  time_threshold = 600
  console_api_key = "${module.test_instance_bunny.console_api_key}"
}

resource "cloudamqp_alarm_configuration" "test_alarm_connection" {
  type            = "connection"
  value_threshold = 100
  time_threshold = 600
  console_api_key = "${module.test_instance_bunny.console_api_key}"
}

resource "cloudamqp_alarm_configuration" "test_alarm_netsplit" {
  type            = "netsplit"
  console_api_key = "${module.test_instance_bunny.console_api_key}"
}

resource "cloudamqp_alarm_configuration" "test_alarm_queue" {
  type            = "queue"
  value_threshold = 100
  time_threshold = 600
  vhost_regex = ".*"
  queue_regex = ".*"
  console_api_key = "${module.test_instance_bunny.console_api_key}"
}
