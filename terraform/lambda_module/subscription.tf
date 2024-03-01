resource "aws_sns_topic_subscription" "user_updates_lampda_target" {
  topic_arn = var.sns_topic_arn
  protocol  = "lambda"
  endpoint  = aws_lambda_function.lambda-function.arn
}
