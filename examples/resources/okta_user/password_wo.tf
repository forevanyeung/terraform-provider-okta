resource "okta_user" "test" {
  first_name          = "TestAcc"
  last_name           = "Smith"
  login               = "testAcc-replace_with_uuid@example.com"
  email               = "testAcc-replace_with_uuid@example.com"
  password_wo         = "Abcd1234"
  password_wo_version = 1
}
