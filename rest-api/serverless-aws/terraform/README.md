# Use Terraform to setup the serverless REST API

**WORK IN PROGRESS!**

## Step-by-step

1. Create a new IAM User used to interact (create, read, delete) with the Lambda and API Gateway resources. We'll avoid the administrator access IAM policies.
2. Reference this user credentials in the `provider.tf` file
3. Run terraform to create:
* Lambda function
* API Gateway
4. Run the same tests as the manual `awscli` of setting the resources up

AWS Lambda create funtion IAM policy:
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "CreateFunctionPermissions",
            "Effect": "Allow",
            "Action": [
                "lambda:CreateFunction"
            ],
            "Resource": "*"
        }
    ]
}
```