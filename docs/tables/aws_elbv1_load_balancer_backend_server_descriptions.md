
# Table: aws_elbv1_load_balancer_backend_server_descriptions
Information about the configuration of an EC2 instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|load_balancer_id|uuid|Unique ID of aws_elbv1_load_balancers table (FK)|
|instance_port|integer|The port on which the EC2 instance is listening.|
|policy_names|text[]|The names of the policies enabled for the EC2 instance.|
