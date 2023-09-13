package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCheckHTTPResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { _ = testAccAPIClient(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories(),
		Steps: []resource.TestStep{
			{
				Config: testRenderSnippet(t, "resource_check_http.tf", 0, nil),
			},
			{
				Config: testRenderSnippet(t, "resource_check_http.tf", 1, nil),
			},
			{
				Config: testRenderSnippet(t, "resource_check_http.tf", 2, nil),
			},
		},
	})
}