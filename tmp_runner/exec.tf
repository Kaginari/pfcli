
resource "null_resource" "register_runner" {


  triggers = {
    volume = docker_volume.runner_volume.name
    base_image = var.base_image
    tmp_runner_token = var.tmp_runner_token
    runner_dsc = var.runner_dsc

  }

  provisioner local-exec {
    when = create
    command =<<EOT
        sleep 5 && docker run --rm -v ${self.triggers.volume}:/etc/gitlab-runner gitlab/gitlab-runner register   --non-interactive     --executor "docker"   --docker-image ${self.triggers.base_image}      --url "https://gitlab.com/"      --registration-token ${self.triggers.tmp_runner_token}    --description ${self.triggers.runner_dsc}   --tag-list "golang-tmp"    --run-untagged="true"     --locked="false"   --access-level="not_protected"
EOT
  }
}
resource "null_resource" "restart_runner" {
  triggers = {
    always = uuid()
    name = docker_container.tmp_runner.name
  }
  provisioner "local-exec" {
    command ="docker restart ${self.triggers.name}"
  }
}
