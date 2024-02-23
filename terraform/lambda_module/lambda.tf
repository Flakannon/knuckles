resource "aws_lambda_function" "lambda-function" {
  filename         = var.filename
  function_name    = var.function_name
  role             = var.role
  source_code_hash = filebase64sha256(var.filename)
  runtime          = "go1.x"
  timeout          = 30
  memory_size      = 1024
  handler          = var.handler // handler needs to be the name of the executable file that contains the relevant main function

  environment {
    variables = {
    }
  }
}
