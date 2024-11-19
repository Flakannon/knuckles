resource "aws_sqs_queue" "queue" {
  name = "${var.queue_name}_sqs"

  redrive_policy = jsonencode({
    "deadLetterTargetArn" = aws_sqs_queue.dlq.arn,
    "maxReceiveCount"     = 3 //number of times a message can be retried before moving to DLQ
  })
}

resource "aws_sqs_queue" "dlq" {
  name                      = "${var.queue_name}_sqs-dlq"
  message_retention_seconds = 1209600 //max retention for DLQ : 14 days
}
