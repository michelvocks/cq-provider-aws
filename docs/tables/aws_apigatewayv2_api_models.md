
# Table: aws_apigatewayv2_api_models
Represents a data model for an API.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|api_id|uuid|Unique ID of aws_apigatewayv2_apis table (FK)|
|model_template|text||
|name|text|The name of the model. Must be alphanumeric.|
|content_type|text|The content-type for the model, for example, "application/json".|
|description|text|The description of the model.|
|model_id|text|The model identifier.|
|schema|text|The schema for the model. For application/json models, this should be JSON schema draft 4 model.|
