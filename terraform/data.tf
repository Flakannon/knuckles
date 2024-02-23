data "archive_file" "knuckles_zip" {
  type        = "zip"
  source_file = "../build/knuckles"
  output_path = "knuckles.zip"
}
