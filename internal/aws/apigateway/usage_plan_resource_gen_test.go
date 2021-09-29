// Code generated by generators/resource/main.go; DO NOT EDIT.

package apigateway_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSApiGatewayUsagePlan_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGateway::UsagePlan", "awscc_apigateway_usage_plan", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
			),
		},
		{
			ResourceName:      td.ResourceName,
			ImportState:       true,
			ImportStateVerify: true,
		},
	})
}

func TestAccAWSApiGatewayUsagePlan_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::ApiGateway::UsagePlan", "awscc_apigateway_usage_plan", "test")

	td.ResourceTest(t, []resource.TestStep{
		{
			Config: td.EmptyConfig(),
			Check: resource.ComposeTestCheckFunc(
				td.CheckExistsInAWS(),
				td.DeleteResource(),
			),
			ExpectNonEmptyPlan: true,
		},
	})
}