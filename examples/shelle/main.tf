terraform {
  required_providers {
    shelle = {
      source = "matthias-herrmann/shelle"
      version = "0.2.2"
    }
  }
}

provider "shelle" {
  # Configuration options
}



data "hashicups_shelle" "edu" {
  command_text = "ping -t 10 google.de"
}

output "data" {
  value = data.hashicups_shelle.edu
}