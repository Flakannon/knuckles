module "lambda-function" {
  source        = "./lambda_module"
  function_name = "battlelambda"
  handler       = "battlelambda"
  filename      = data.archive_file.battlelambda_zip.output_path
  role          = aws_iam_role.lambda_role.arn
  app_version   = "0.0.2" //TODO: have it TFVAR'd in by pipeline when its set up
  sns_topic_arn = aws_sns_topic.game_starter.arn
}
