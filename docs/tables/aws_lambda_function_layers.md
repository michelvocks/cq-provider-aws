
# Table: aws_lambda_function_layers
An AWS Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html). 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_id|uuid|Unique ID of aws_lambda_functions table (FK)|
|arn|text|The Amazon Resource Name (ARN) of the function layer.|
|code_size|bigint|The size of the layer archive in bytes.|
|signing_job_arn|text|The Amazon Resource Name (ARN) of a signing job.|
|signing_profile_version_arn|text|The Amazon Resource Name (ARN) for a signing profile version.|
