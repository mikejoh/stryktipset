resource "aws_iam_role" "iam_for_lambda" {
  name = "lambda-stryktipset-executor"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "stryktipset" {
  filename         = "../stryktipset.zip"
  function_name    = "stryktipset"
  role             = "${aws_iam_role.iam_for_lambda.arn}"
  handler          = "stryktipset"
  source_code_hash = "${base64sha256(file("../stryktipset.zip"))}"
  runtime          = "go1.x"
}