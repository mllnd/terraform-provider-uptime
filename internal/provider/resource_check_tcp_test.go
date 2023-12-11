package provider

import (
	"testing"

	petname "github.com/dustinkirkland/golang-petname"
	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccCheckTCPResource(t *testing.T) {
	names := [2]string{
		petname.Generate(3, "-"),
		petname.Generate(3, "-"),
	}
	resource.Test(t, testCaseFromSteps(t, []resource.TestStep{
		{
			ConfigVariables: config.Variables{
				"name":    config.StringVariable(names[0]),
				"address": config.StringVariable("example.com"),
				"port":    config.IntegerVariable(80),
			},
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/_basic"),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "name", names[0]),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "address", "example.com"),
			),
		},
		{
			ConfigVariables: config.Variables{
				"name":    config.StringVariable(names[1]),
				"address": config.StringVariable("example.net"),
				"port":    config.IntegerVariable(80),
			},
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/_basic"),
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "name", names[1]),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "address", "example.net"),
			),
		},
	}))
}

func TestAccCheckTCPResource_ContactGroups(t *testing.T) {
	name := petname.Generate(3, "-")
	resource.Test(t, testCaseFromSteps(t, []resource.TestStep{
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/contact_groups"),
			ConfigVariables: config.Variables{
				"name": config.StringVariable(name),
				"port": config.IntegerVariable(80),
				"contact_groups": config.ListVariable(
					config.StringVariable("nobody"),
				),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "contact_groups.#", "1"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "contact_groups.0", "nobody"),
			),
		},
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/contact_groups"),
			ConfigVariables: config.Variables{
				"name": config.StringVariable(name),
				"port": config.IntegerVariable(80),
				"contact_groups": config.ListVariable(
					config.StringVariable("nobody"),
					config.StringVariable("noone"),
				),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "contact_groups.#", "2"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "contact_groups.0", "nobody"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "contact_groups.1", "noone"),
			),
		},
	}))
}

func TestAccCheckTCPResource_Interval(t *testing.T) {
	name := petname.Generate(3, "-")
	resource.Test(t, testCaseFromSteps(t, []resource.TestStep{
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/interval"),
			ConfigVariables: config.Variables{
				"name":     config.StringVariable(name),
				"port":     config.IntegerVariable(80),
				"interval": config.IntegerVariable(5),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "interval", "5"),
			),
		},
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/interval"),
			ConfigVariables: config.Variables{
				"name":     config.StringVariable(name),
				"port":     config.IntegerVariable(80),
				"interval": config.IntegerVariable(10),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "interval", "10"),
			),
		},
	}))
}

func TestAccCheckTCPResource_Locations(t *testing.T) {
	name := petname.Generate(3, "-")
	resource.Test(t, testCaseFromSteps(t, []resource.TestStep{
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/locations"),
			ConfigVariables: config.Variables{
				"name": config.StringVariable(name),
				"port": config.IntegerVariable(80),
				"locations": config.ListVariable(
					config.StringVariable("US-CA-Los Angeles"),
					config.StringVariable("US-NY-New York"),
				),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "locations.#", "2"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "locations.0", "US-CA-Los Angeles"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "locations.1", "US-NY-New York"),
			),
		},
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/locations"),
			ConfigVariables: config.Variables{
				"name": config.StringVariable(name),
				"port": config.IntegerVariable(80),
				"locations": config.ListVariable(
					config.StringVariable("Finland-Helsinki"),
					config.StringVariable("Switzerland-Zurich"),
				),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "locations.#", "2"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "locations.0", "Finland-Helsinki"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "locations.1", "Switzerland-Zurich"),
			),
		},
	}))
}

func TestAccCheckTCPResource_NumRetries(t *testing.T) {
	name := petname.Generate(3, "-")
	resource.Test(t, testCaseFromSteps(t, []resource.TestStep{
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/num_retries"),
			ConfigVariables: config.Variables{
				"name":        config.StringVariable(name),
				"port":        config.IntegerVariable(80),
				"num_retries": config.IntegerVariable(3),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "num_retries", "3"),
			),
		},
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/num_retries"),
			ConfigVariables: config.Variables{
				"name":        config.StringVariable(name),
				"port":        config.IntegerVariable(80),
				"num_retries": config.IntegerVariable(2),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "num_retries", "2"),
			),
		},
	}))
}

func TestAccCheckTCPResource_SendExpectString(t *testing.T) {
	name := petname.Generate(3, "-")
	resource.Test(t, testCaseFromSteps(t, []resource.TestStep{
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/send_expect_string"),
			ConfigVariables: config.Variables{
				"name":          config.StringVariable(name),
				"send_string":   config.StringVariable("foo"),
				"expect_string": config.StringVariable("bar"),
				"port":          config.IntegerVariable(80),
				"num_retries":   config.IntegerVariable(3),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "send_string", "foo"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "expect_string", "bar"),
			),
		},
		{
			ConfigDirectory: config.StaticDirectory("testdata/resource_check_tcp/send_expect_string"),
			ConfigVariables: config.Variables{
				"name":          config.StringVariable(name),
				"send_string":   config.StringVariable("foobar"),
				"expect_string": config.StringVariable("baz"),
				"port":          config.IntegerVariable(80),
				"num_retries":   config.IntegerVariable(2),
			},
			Check: resource.ComposeAggregateTestCheckFunc(
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "send_string", "foobar"),
				resource.TestCheckResourceAttr("uptime_check_tcp.test", "expect_string", "baz"),
			),
		},
	}))
}
