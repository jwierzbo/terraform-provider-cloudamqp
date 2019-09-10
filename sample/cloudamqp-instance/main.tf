resource "cloudamqp_instance" "generic_instance" {
  name   = "${var.instance_name}"
  plan   = "${var.instance_plan}"
  region = "${var.instance_region}"

  lifecycle {
    ignore_changes = [
      "nodes"
    ]
  }
}
