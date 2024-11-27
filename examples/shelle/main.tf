terraform {
  required_providers {
    hashicups = {
      source = "hashicorp.com/edu/hashicups"
    }
  }
}

provider "hashicups" {
}

data "hashicups_shelle" "edu" {
  command_text = "ping -t 10 google.de"
}

output "data" {
  value = data.hashicups_shelle.edu
}