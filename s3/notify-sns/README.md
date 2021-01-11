# Send notifications on fileupload

An example on how to set up a S3 bucket to send a notification event to a SQS queue when a file was uploaded.

## Getting Started

These instructions will get you a copy of the project up and running on AWS.

### Prerequisites

What things you need to install the software and how to install them

- [AWS Free Tier account](https://aws.amazon.com/free)
- Terraform
- AWS CLI

```
brew install terraform awscli
```
### Configure your AWS CLI

Run the following command and complete the wizard with your info.
```
aws configure
```

### Creating AWS resources with Terraform

Run the following command replacing `bucket_name` with your bucket name:

```
terraform apply -var bucket_name=REPLACE-ME
```

### Checking notifications on file upload

1. Upload a new file to the bucket
    ```
    aws s3 cp upload.me s3://REPLACE-ME
    ```
1. Run the following command to consume the notification from SQS:
    ```
    go run ReceiveMessage.go -q sqs-for-REPLACE-ME-WITH-BUCKET_NAME
    ```

## Tearing down

After you are done, remove all AWS resources by running the following command:

```
terraform destroy -var bucket_name=REPLACE-ME
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

aws s3 cp upload.me.log s3://isaacasensio-s3-notification

    