
# Table: aws_ec2_security_group_ip_permission_ip_ranges
Describes an IPv4 range.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|security_group_ip_permission_id|uuid|Unique ID of aws_ec2_security_group_ip_permissions table (FK)|
|cidr_ip|text|The IPv4 CIDR range.|
|description|text|A description for the security group rule that references this IPv4 address range.|
