module "lambda-function" {
  source        = "./lambda_module"
  function_name = "knuckles"
  handler       = "main"
  filename      = data.archive_file.knuckles_zip.output_path
  role          = aws_iam_role.lambda_role.arn
}
