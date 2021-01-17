# Restrict file uploads based on file type

An example on how to set up a S3 bucket to restrict file uploading based on the file extension.

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

### Upload file

1. Upload a new file to the bucket (it will be rejected as it is not a ZIP file)
    ```
    aws s3 cp upload.me s3://REPLACE-ME
    ```
1. Upload a zip file to the bucket.
    ```
    aws s3 cp upload.me.zip s3://REPLACE-ME
    ```

## Tearing down

After you are done, remove all AWS resources by running the following command:

```
terraform destroy -var bucket_name=REPLACE-ME
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
    