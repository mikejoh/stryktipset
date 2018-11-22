resource "aws_api_gateway_rest_api" "stryktipset" {
  name        = "stryktipset"
  description = "Stryktipset convert REST API endpoint demo"
}

resource "aws_lambda_permission" "lambda_permission" {
  statement_id  = "AllowStryktipsetInvoke"
  action        = "lambda:InvokeFunction"
  function_name = "${aws_lambda_function.stryktipset.function_name}"
  principal     = "apigateway.amazonaws.com"
  source_arn = "${aws_api_gateway_rest_api.stryktipset.execution_arn}/*/*/*"
}

resource "aws_api_gateway_resource" "stryktipset" {
  rest_api_id = "${aws_api_gateway_rest_api.stryktipset.id}"
  parent_id   = "${aws_api_gateway_rest_api.stryktipset.root_resource_id}"
  path_part   = "convert"
}

resource "aws_api_gateway_method" "stryktipset" {
  rest_api_id   = "${aws_api_gateway_rest_api.stryktipset.id}"
  resource_id   = "${aws_api_gateway_resource.stryktipset.id}"
  http_method   = "ANY"
    authorization        = "NONE"
}

resource "aws_api_gateway_integration" "stryktipset" {
  rest_api_id          = "${aws_api_gateway_rest_api.stryktipset.id}"
  resource_id          = "${aws_api_gateway_resource.stryktipset.id}"
  http_method          = "${aws_api_gateway_method.stryktipset.http_method}"
  type                 = "AWS_PROXY"
  integration_http_method = "POST"
  uri = "arn:aws:apigateway:${var.myregion}:lambda:path/2015-03-31/functions/${aws_lambda_function.stryktipset.arn}/invocations"
}

resource "aws_api_gateway_deployment" "stryktipset" {
    depends_on = [
        "aws_api_gateway_integration.stryktipset"
    ]
    rest_api_id = "${aws_api_gateway_rest_api.stryktipset.id}"
    stage_name  = "staging"
}

output "base_url" {
  value = "${aws_api_gateway_deployment.stryktipset.invoke_url}"
}