
# Table: aws_lambda_function_version_file_system_configs
Details about the connection between a Lambda function and an Amazon EFS file system. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|function_version_id|uuid|Unique ID of aws_lambda_function_versions table (FK)|
|arn|text|The Amazon Resource Name (ARN) of the Amazon EFS access point that provides access to the file system.|
|local_mount_path|text|The path where the function can access the file system, starting with /mnt/.|
