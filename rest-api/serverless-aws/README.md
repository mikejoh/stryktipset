# Running serverless with AWS Lambda

This shows how to serve the `convert` (see `helpers.go`) functionality of the `github.com/mikejoh/stryktipset` package through a AWS Lambda function and an AWS API Gateway.

When we're done this will be the resulting URL you would use to invoke the conversion function:
```
https://REST_API_ID.execute-api.eu-west-1.amazonaws.com/staging/convert?sek=192
{"sek":192,"full":1,"half":6}
```

Heavily inspired by [this](https://www.alexedwards.net/blog/serverless-api-with-go-and-aws-lambda) blog post!

## Prerequisites

* Create a new IAM user, for the purpose of just testing Lambda and API gateway i attached the `AdministratorAccess` policy to the user. Not to be used in production!

## Terraform step-by-step guide

Coming soon!

## Semi-manual step-by-step guide

_Some of these steps require you to have [`jq`](https://stedolan.github.io/jq/) installed. It's a fantastic tool to for slicing and filter and map and transform JSON data._

_At the moment each step will be copy-pasted into the CLI you're using. Not great and when i have time to re-create everything multiple lines could fit into one or i'll just create a bash script._

_If you for want to change the Lambda function for any reason, use this command to re-upload the function code: `aws lambda update-function-code --function-name stryktipset --zip-file fileb://${ZIP_PATH}`_

1. Run `make aws-lambda-build` to cross compile for Linux (required to be used by Lambda) and zip up the resulting binary. This zip-file will be used when creating the Lambda function.

2. Create the IAM role (permissions) that the Lambda function will use when running. the trust policy document will allow Lambda services to assume the `lambda-stryktipset-executor` role:
```
ROLE_NAME="lambda-stryktipset-executor"
TRUST_POLICY_PATH="./trust-policy.json"
aws iam create-role --role-name $ROLE_NAME --assume-role-policy-document file://${TRUST_POLICY_PATH}
ROLE_ARN=$(aws iam list-roles | jq -r --arg ROLE_NAME "$ROLE_NAME" '.Roles[] | select(.RoleName == $ROLE_NAME) | .Arn')
```

3. Specify the permissions that the new role has. The policy name is somewhat self-explanatory.
```
aws iam attach-role-policy --role-name $ROLE_NAME --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
```

4. Create the Lambda function
```
FUNC_NAME="stryktipset"
ZIP_PATH="./stryktipset.zip"
aws lambda create-function --function-name $FUNC_NAME --runtime go1.x --role $ROLE_ARN --handler stryktipset --zip-file fileb://${ZIP_PATH}
```

5. Create the API Gateway
```
API_NAME="stryktipset"
aws apigateway create-rest-api --name $API_NAME
REST_API_ID=$(aws apigateway get-rest-apis | jq -r '.items[] | select(.name == "stryktipset") | .id')
PARENT_RESOURCE_ID=$(aws apigateway get-resources --rest-api-id $REST_API_ID | jq -r '.items[] | select(.path == "/") | .id')
```

6. Create the API gateway resource in this case the endpoint `/convert`
```
aws apigateway create-resource --rest-api-id $REST_API_ID --parent-id $PARENT_RESOURCE_ID --path-part convert
CONVERT_RESOURCE_ID=$(aws apigateway get-resources --rest-api-id $REST_API_ID | jq -r '.items[] | select(.path == "/convert") | .id')
```
_Note that it's possible to include placeholders within your path by wrapping part of the path in curly braces. For example, a --path-part parameter of books/{id} would match requests to /books/foo and /books/bar_

7. Add `ANY` method to the `/convert` endpoint
```
aws apigateway put-method --rest-api-id $REST_API_ID --resource-id $CONVERT_RESOURCE_ID --http-method ANY --authorization-type NONE
```

8. Configure the method integration, we're using the Lambda proxy integration type (AWS_PROXY)

_The `--uri` flag below will have the following form: arn:aws:apigateway:{region}:{subdomain.service|service}:path|action/{service_api}_

```
FUNC_ARN=$(aws lambda list-functions | jq -r '.Functions[] | select(.FunctionName == "stryktipset") | .FunctionArn')
aws apigateway put-integration --rest-api-id $REST_API_ID --resource-id $CONVERT_RESOURCE_ID --http-method ANY --type AWS_PROXY --integration-http-method POST --uri arn:aws:apigateway:eu-west-1:lambda:path/2015-03-31/functions/${FUNC_ARN}/invocations
```

9. Add permission to the resource policy with the specified Lambda function. **Replace the ACCOUNT_ID below!**
```
aws lambda add-permission --function-name stryktipset --statement-id stryktipset2018 --action lambda:InvokeFunction --principal apigateway.amazonaws.com --source-arn arn:aws:execute-api:eu-west-1:ACCOUNT_ID:${REST_API_ID}/*/*/*
```

10. Deploy the API, to be able to access it externally. Use any name you want as the `--stage-name`

```
aws apigateway create-deployment --rest-api-id $REST_API_ID \
--stage-name staging
```

11. We should be done by now, let's use `awscli` to run a test invokation

```
aws apigateway test-invoke-method --rest-api-id $REST_API_ID --resource-id $CONVERT_RESOURCE_ID --http-method "GET" --path-with-query-string "/convert?sek=192"
```

You should also be able to access the following URL in a browser with the path and query parameters as above:

```
curl https://${REST_API_ID}.execute-api.eu-west-1.amazonaws.com/staging/convert?sek=192
```
The response (JSON) should look like this:
```
{"sek":192,"full":1,"half":6}
```
