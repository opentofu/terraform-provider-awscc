// Code generated by generators/resource/main.go; DO NOT EDIT.

package refactorspaces_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRefactorSpacesEnvironment_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RefactorSpaces::Environment", "awscc_refactorspaces_environment", "test")

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

func TestAccAWSRefactorSpacesEnvironment_disappears(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::RefactorSpaces::Environment", "awscc_refactorspaces_environment", "test")

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