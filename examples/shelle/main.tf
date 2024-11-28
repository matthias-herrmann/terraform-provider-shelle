terraform {
  required_providers {
    shelle = {
      source = "matthias-herrmann/shelle"
      version = "0.2.4"
    }
  }
}

provider "shelle" {
  # Configuration options
}

data "shelle_shelle" "name" {
  command_text = "echo Hello, World!"
}