terraform {

  extra_arguments "plan_with_vars" {
    commands = [
      "plan"]

    arguments = [
      "-var-file=var.tfvars"
    ]
  }

  extra_arguments "apply_autoapprouve" {
    commands = [
      "apply",
      "destroy"]

    arguments = [
      "-var-file=var.tfvars",
      "-auto-approve"
    ]
  }

}
remote_state {
  backend = "local"
  config = {
    path = ".terraform/runner.tfstate"
  }
}
