# Running serverless with AWS Lambda

**Work in progress!!!**

This shows how to serve the `convert` (see `helpers.go`) functionality of the `github.com/mikejoh/stryktipset` package through a AWS Lambda function and an AWS API Gateway.

Overview of what we'll do in this example:

1. Install and configure the `awscli` tool
2. Configure `awscli`
2. Create an AWS IAM user
3. Install needed Go packages
4. Create an simple Go function that serves a `convert` function
5. Create an AWS Lambda function
6. Create an API Gateway and send events to the Lambda function
7. Test the whole setup externally

## Step-by-step

