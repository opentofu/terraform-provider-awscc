// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53ResolverResolverConfigDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverConfig", "awscc_route53resolver_resolver_config", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSRoute53ResolverResolverConfigDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53Resolver::ResolverConfig", "awscc_route53resolver_resolver_config", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}