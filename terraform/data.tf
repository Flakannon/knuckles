data "archive_file" "battlelambda_zip" {
  type        = "zip"
  source_file = "../build/battlelambda"
  output_path = "battlelambda.zip"
}
