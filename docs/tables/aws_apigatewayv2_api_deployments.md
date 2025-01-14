
# Table: aws_apigatewayv2_api_deployments
An immutable representation of an API that can be called by users.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|auto_deployed|boolean|Specifies whether a deployment was automatically released.|
|created_date|timestamp without time zone|The date and time when the Deployment resource was created.|
|deployment_id|text|The identifier for the deployment.|
|deployment_status|text|The status of the deployment: PENDING, FAILED, or SUCCEEDED.|
|deployment_status_message|text|May contain additional feedback on the status of an API deployment.|
|description|text|The description for the deployment.|
