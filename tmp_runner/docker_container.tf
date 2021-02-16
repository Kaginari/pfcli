


data docker_registry_image tmp_runner_image {
name = "gitlab/gitlab-runner:latest"
}

resource docker_image runner_base_image {
  name = data.docker_registry_image.tmp_runner_image.name
  keep_locally = true
  pull_triggers = [data.docker_registry_image.tmp_runner_image.sha256_digest]
}
resource docker_volume runner_volume {
  name = "config_runner"
}
resource "docker_container" "tmp_runner" {
  image = docker_image.runner_base_image.latest
  name = "tmp_runner"
  volumes {
    volume_name = docker_volume.runner_volume.name
    container_path = "/etc/gitlab-runner"
  }
  volumes {
     host_path = "/var/run/docker.sock"
    container_path = "/var/run/docker.sock"
  }

  network_mode = "host"
}

