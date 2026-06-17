# Generate a random password as an ephemeral value so it is never written to
# the Terraform state or plan files, then feed it into the write-only
# `password_wo` attribute of the user. Requires Terraform 1.11+ and the
# hashicorp/random provider 3.7.0+.

ephemeral "random_password" "user" {
  length  = 32
  special = true
}

resource "okta_user" "example" {
  first_name = "Example"
  last_name  = "User"
  login      = "example.user@example.com"
  email      = "example.user@example.com"

  # The generated password is applied to the user but never persisted in state.
  password_wo         = ephemeral.random_password.user.result
  password_wo_version = 1

  # expire_password_on_create works with either `password` or `password_wo`,
  # forcing the user to set their own password at first login.
  expire_password_on_create = true
}
