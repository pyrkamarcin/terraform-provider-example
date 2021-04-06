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

resource "example_cluster" "test-server-1" {
  name = "test"

  specification {
    type = "test"
  }

  components {
    postgresql {
      count = 1
      machines = [
        "super-machine-1"
      ]
    }
  }
}
