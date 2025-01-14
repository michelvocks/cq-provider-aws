package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ConfigConformancePack() *schema.Table {
	return &schema.Table{
		Name:         "aws_config_conformance_packs",
		Description:  "Returns details of a conformance pack.",
		Resolver:     fetchConfigConformancePacks,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "conformance_pack_arn",
				Description: "Amazon Resource Name (ARN) of the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "conformance_pack_id",
				Description: "ID of the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "conformance_pack_name",
				Description: "Name of the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "conformance_pack_input_parameters",
				Description: "A list of ConformancePackInputParameter objects.",
				Type:        schema.TypeJSON,
				Resolver:    resolveConfigConformancePackConformancePackInputParameters,
			},
			{
				Name:        "created_by",
				Description: "AWS service that created the conformance pack.",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_s3_bucket",
				Description: "Amazon S3 bucket where AWS Config stores conformance pack templates.",
				Type:        schema.TypeString,
			},
			{
				Name:        "delivery_s3_key_prefix",
				Description: "The prefix for the Amazon S3 bucket.",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_update_requested_time",
				Description: "Last time when conformation pack update was requested.",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchConfigConformancePacks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	config := configservice.DescribeConformancePacksInput{}
	for {
		resp, err := c.Services().ConfigService.DescribeConformancePacks(ctx, &config, func(options *configservice.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- resp.ConformancePackDetails
		if resp.NextToken == nil {
			break
		}
		config.NextToken = resp.NextToken
	}
	return nil
}
func resolveConfigConformancePackConformancePackInputParameters(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	conformancePack := resource.Item.(types.ConformancePackDetail)
	params := make(map[string]*string, len(conformancePack.ConformancePackInputParameters))
	for _, p := range conformancePack.ConformancePackInputParameters {
		params[*p.ParameterName] = p.ParameterValue
	}
	return resource.Set(c.Name, params)
}
