module "messaging" {
  source     = "./sqs_module"
  queue_name = "knuckles-orchestrator"
}
