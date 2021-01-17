provider "aws" {
  region = var.aws_region
}

# Create a S3 bucket which only allows zip files
resource "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name
  acl = "private"
  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "Policy1464968545158",
  "Statement": [
    {
      "Sid": "Stmt1464968483619",
      "Effect": "Deny",
      "Principal": "*",
      "Action": "s3:PutObject",
      "NotResource": [
        "arn:aws:s3:::${var.bucket_name}/*.zip"
      ]
    }
  ]
}
POLICY
}