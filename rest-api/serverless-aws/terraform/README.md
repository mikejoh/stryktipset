# Use Terraform to setup the serverless REST API demo

## Step-by-step

1. Create a new IAM User and generate access keys. Take a note of the `AccessKeyId` and `SecretAccessKey` fields when you generate the access keys below, you'll need them in step 2:
```
aws iam create-user --user-name terraform
aws iam create-access-keys --user-name terraform
```
2. Attach a policy to the newly created user, i used the AdministratorAccess policy (!).
3. Create a new awscli profile and reference the generated access keys:
```
aws configure --profile terraform
```
4. The `provider.tf` file points at the created profile above and also a default region configured in `vars.tf`, change this as needed!
5. Run `make aws-lambda-build` in the `serverless-aws/` directory to compile a binary and create the needed zip-file.
6. You should now be able to run terraform to create what you need:
```
terraform init
terraform plan
terraform apply
```

Take a note of the `base_url` output that terraform generates!

The Lambda function and API Gateway with the method, integration and deployment will be created for you. Now you should be able to use the `base_url` output and copy-paste that into a browser, make sure you add `convert?sek=192` in the end of the URL.

If you want to remove everything created with terraform you can run `terraform destroy`.

Have fun!