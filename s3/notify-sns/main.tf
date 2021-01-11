provider "aws" {
  region = var.aws_region
}

# Create a S3 bucket
resource "aws_s3_bucket" "bucket" {
  bucket = var.bucket_name
  acl = "private"
}

# Sends a notification to SNS on new object created
resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.bucket.id

  queue {
    queue_arn     = aws_sqs_queue.queue.arn
    events        = ["s3:ObjectCreated:*"]
  }
}

resource "aws_sqs_queue" "queue" {
  name = "sqs-for-${var.bucket_name}"

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "arn:aws:sqs:*:*:sqs-for-${var.bucket_name}",
      "Condition": {
        "ArnEquals": { "aws:SourceArn": "${aws_s3_bucket.bucket.arn}" }
      }
    }
  ]
}
POLICY
}