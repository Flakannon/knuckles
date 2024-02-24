module "lambda-function" {
  source        = "./lambda_module"
  function_name = "knuckles"
  handler       = "knuckles"
  filename      = data.archive_file.knuckles_zip.output_path
  role          = aws_iam_role.lambda_role.arn
  app_version   = "0.0.2" //TODO: have it TFVAR'd in by pipeline when its set up
}
