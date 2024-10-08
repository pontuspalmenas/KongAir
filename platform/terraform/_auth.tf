# Configure the provider to use your Kong Konnect account
terraform {
  required_providers {
    konnect = {
      source  = "kong/konnect"
      version = "1.0.0"
    }
  }
}
provider "konnect" {
  personal_access_token = env("PERSONAL_ACCESS_TOKEN")
  server_url            = "https://eu.api.konghq.com"
}