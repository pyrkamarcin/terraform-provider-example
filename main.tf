terraform {
  required_providers {
    example = {
      version = "0.1"
      source = "github.com/pyrkamarcin/example"
    }
    azure = {
      source = "hashicorp/azurerm"
      version = ">2.41.0"
    }
  }
}

provider "example" {
}

provider "azure" {
}

resource "example_server" "test-server-1" {
  name = "test"
  vm_ip_address = "10.33.1.200"
  vm_user = "operations"
  vm_ssh_key = "./ssh/key-to-v-machine.pem"
}
