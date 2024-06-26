resource "awscc_wafv2_ip_set" "example" {
  name               = "example"
  description        = "Example IP set"
  scope              = "REGIONAL"
  ip_address_version = "IPV4"
  addresses          = ["1.2.3.4/32", "5.6.7.8/32"]

  tags = [
    {
      key   = "ModifiedBy"
      value = "AWSCC"
    }
  ]
}
