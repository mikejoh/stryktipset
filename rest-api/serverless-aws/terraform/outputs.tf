output "stryktipset_rest_api_endpoint" {
  value = "${aws_api_gateway_deployment.stryktipset.invoke_url}/convert"
}