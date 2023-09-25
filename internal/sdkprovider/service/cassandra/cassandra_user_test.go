package cassandra_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/aiven/aiven-go-client/v2"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"
	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

func TestAccAivenCassandraUser_basic(t *testing.T) {
	resourceName := "aiven_cassandra_user.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acc.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acc.TestProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckAivenCassandraUserResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCassandraUserResource(rName),
				Check: resource.ComposeTestCheckFunc(
					schemautil.TestAccCheckAivenServiceUserAttributes("data.aiven_cassandra_user.user"),
					resource.TestCheckResourceAttr(resourceName, "service_name", fmt.Sprintf("test-acc-sr-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "project", os.Getenv("AIVEN_PROJECT_NAME")),
					resource.TestCheckResourceAttr(resourceName, "username", fmt.Sprintf("user-%s", rName)),
					resource.TestCheckResourceAttr(resourceName, "password", "Test$1234"),
				),
			},
		},
	})
}

func testAccCheckAivenCassandraUserResourceDestroy(s *terraform.State) error {
	c := acc.GetTestAivenClient()

	ctx := context.Background()

	// loop through the resources in state, verifying each aiven_cassandra_user is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aiven_cassandra_user" {
			continue
		}

		projectName, serviceName, username, err := schemautil.SplitResourceID3(rs.Primary.ID)
		if err != nil {
			return err
		}

		p, err := c.ServiceUsers.Get(ctx, projectName, serviceName, username)
		if err != nil {
			if err.(aiven.Error).Status != 404 {
				return err
			}
		}

		if p != nil {
			return fmt.Errorf("common user (%s) still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCassandraUserResource(name string) string {
	return fmt.Sprintf(`
data "aiven_project" "foo" {
  project = "%s"
}

resource "aiven_cassandra" "bar" {
  project                 = data.aiven_project.foo.project
  cloud_name              = "google-europe-west1"
  plan                    = "startup-4"
  service_name            = "test-acc-sr-%s"
  maintenance_window_dow  = "monday"
  maintenance_window_time = "10:00:00"
}

resource "aiven_cassandra_user" "foo" {
  service_name = aiven_cassandra.bar.service_name
  project      = data.aiven_project.foo.project
  username     = "user-%s"
  password     = "Test$1234"

  depends_on = [aiven_cassandra.bar]
}

data "aiven_cassandra_user" "user" {
  service_name = aiven_cassandra.bar.service_name
  project      = aiven_cassandra.bar.project
  username     = aiven_cassandra_user.foo.username

  depends_on = [aiven_cassandra_user.foo]
}`, os.Getenv("AIVEN_PROJECT_NAME"), name, name)
}
