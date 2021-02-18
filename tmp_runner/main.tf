terraform {
  backend "local" {}
  required_providers {
    docker = {
      source = "terraform-providers/docker"
    }
  }
}
provider "docker" {
  host = "tcp://127.0.0.1:2375"

    ## either 1 - provide cert dir path
  //  cert_path = "${VAULT_OUTPUT}/tls"
    ## 2 - provide client cert key and root ca
  //  ca_material   = file("${VAULT_OUTPUT}/tls/docker_ca.crt")
  //  cert_material = file("${VAULT_OUTPUT}/tls/client.crt")
  //  key_material  = file("${VAULT_OUTPUT}/tls/client.key")
}
