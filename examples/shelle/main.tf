terraform {
  required_providers {
    shelle = {
      source = "matthias-herrmann/shelle"
      version = "0.2.3"
    }
  }
}

provider "shelle" {
  # Configuration options
}

data "hashicups_shelle" "name" {
  
}