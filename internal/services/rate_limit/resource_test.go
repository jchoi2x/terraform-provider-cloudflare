package rate_limit_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"testing"

	cloudflare "github.com/cloudflare/cloudflare-go"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/consts"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccCloudflareRateLimit_Basic(t *testing.T) {
	// multiple instances of this config would conflict but we only use it once
	t.Parallel()
	var rateLimit cloudflare.RateLimit
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_rate_limit." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRateLimitConfigBasic(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zoneID),
					// dont check that specified values are set, this will be evident by lack of plan diff
					// some values will get empty values
					resource.TestCheckResourceAttr(name, "action.0.response.#", "0"),
					resource.TestCheckResourceAttr(name, "bypass_url_patterns.#", "0"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.statuses.#", "0"),
					resource.TestCheckResourceAttr(name, "disabled", "false"),
					// also expect api to generate some values
					resource.TestCheckResourceAttr(name, "match.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.request.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.request.0.schemes.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.request.0.url_pattern", "*"),
					resource.TestCheckResourceAttr(name, "match.0.response.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.origin_traffic", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRateLimitChallenge_Basic(t *testing.T) {
	t.Parallel()
	var rateLimit cloudflare.RateLimit
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_rate_limit." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRateLimitChallengeConfigBasic(zoneID, rnd),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zoneID),
					// check that the action challenge mode has been set
					resource.TestCheckResourceAttr(name, "action.0.mode", "challenge"),
					resource.TestCheckResourceAttr(name, "action.0.response.#", "0"),
					resource.TestCheckResourceAttr(name, "bypass_url_patterns.#", "0"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.statuses.#", "0"),
					resource.TestCheckResourceAttr(name, "disabled", "false"),
					resource.TestCheckResourceAttr(name, "match.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.request.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.request.0.schemes.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.request.0.url_pattern", "*"),
					resource.TestCheckResourceAttr(name, "match.0.response.#", "1"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.origin_traffic", "true"),
				),
			},
		},
	})
}

func TestAccCloudflareRateLimit_FullySpecified(t *testing.T) {
	t.Parallel()
	var rateLimit cloudflare.RateLimit
	zoneName := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_rate_limit." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRateLimitConfigFullySpecified(zoneID, rnd, zoneName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zoneID),
					// checking our overrides of default values worked
					resource.TestCheckResourceAttr(name, "action.0.response.#", "1"),
					resource.TestCheckResourceAttr(name, "action.0.response.0.content_type", "text/plain"),
					resource.TestCheckResourceAttr(name, "action.0.response.0.body", "my response body"),
					resource.TestCheckResourceAttr(name, "bypass_url_patterns.#", "2"),
					resource.TestCheckResourceAttr(name, "match.0.request.0.methods.#", "6"),
					resource.TestCheckResourceAttr(name, "match.0.request.0.schemes.#", "2"),
					resource.TestMatchResourceAttr(name, "match.0.request.0.url_pattern", regexp.MustCompile("tfacc-full")),
					resource.TestCheckResourceAttr(name, "match.0.response.0.origin_traffic", "false"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.statuses.#", "5"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.headers.#", "2"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.headers.0.name", "Test"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.headers.0.op", "ne"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.headers.0.value", "test"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.headers.1.name", "Host"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.headers.1.op", "eq"),
					resource.TestCheckResourceAttr(name, "match.0.response.0.headers.1.value", "localhost"),
					resource.TestCheckResourceAttr(name, "correlate.0.by", "nat"),
					resource.TestCheckResourceAttr(name, "disabled", "true"),
					resource.TestCheckResourceAttr(name, "description", "my fully specified rate limit for a zone"),
				),
			},
		},
	})
}

func TestAccCloudflareRateLimit_Update(t *testing.T) {
	t.Parallel()
	var rateLimit cloudflare.RateLimit
	var initialRateLimitId string
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_rate_limit." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRateLimitConfigMatchingUrl(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zoneID),
				),
			},
			{
				PreConfig: func() {
					initialRateLimitId = rateLimit.ID
				},
				Config: testAccCheckCloudflareRateLimitConfigFullySpecified(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zoneID),
					func(state *terraform.State) error {
						if initialRateLimitId != rateLimit.ID {
							// rate limit change shows resource was recreated, we want in place update
							return fmt.Errorf("rate limit id is different after second config applied ( %s -> %s )", initialRateLimitId, rateLimit.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudflareRateLimit_CreateAfterManualDestroy(t *testing.T) {
	t.Parallel()
	var rateLimit cloudflare.RateLimit
	var initialRateLimitId string
	domain := os.Getenv("CLOUDFLARE_DOMAIN")
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()
	name := "cloudflare_rate_limit." + rnd

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudflareRateLimitConfigMatchingUrl(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zoneID),
					testAccManuallyDeleteRateLimit(name, &rateLimit, &initialRateLimitId),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: testAccCheckCloudflareRateLimitConfigMatchingUrl(zoneID, rnd, domain),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCloudflareRateLimitExists(name, &rateLimit),
					testAccCheckCloudflareRateLimitIDIsValid(name, zoneID),
					func(state *terraform.State) error {
						if initialRateLimitId == rateLimit.ID {
							return fmt.Errorf("rate limit id is unchanged even after we thought we deleted it ( %s )", rateLimit.ID)
						}
						return nil
					},
				),
			},
		},
	})
}

func TestAccCloudflareRateLimit_WithoutTimeout(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareRateLimitConfigWithoutTimeout(zoneID, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("rate limit timeout must be set if the 'mode' is simulate or ban")),
			},
		},
	})
}

func TestAccCloudflareRateLimit_ChallengeWithTimeout(t *testing.T) {
	t.Parallel()
	zoneID := os.Getenv("CLOUDFLARE_ZONE_ID")
	rnd := utils.GenerateRandomResourceName()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareRateLimitDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckCloudflareRateLimitChallengeConfigWithTimeout(zoneID, rnd),
				ExpectError: regexp.MustCompile(regexp.QuoteMeta("rate limit timeout must not be set if the 'mode' is challenge or js_challenge")),
			},
		},
	})
}

func testAccCheckCloudflareRateLimitDestroy(s *terraform.State) error {
	client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
	if clientErr != nil {
		tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_rate_limit" {
			continue
		}

		_, err := client.RateLimit(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Rate limit still exists")
		}
	}

	return nil
}

func testAccCheckCloudflareRateLimitExists(n string, rateLimit *cloudflare.RateLimit) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Rate Limit ID is set")
		}

		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		foundRateLimit, err := client.RateLimit(context.Background(), rs.Primary.Attributes[consts.ZoneIDSchemaKey], rs.Primary.ID)
		if err != nil {
			return err
		}

		if foundRateLimit.ID != rs.Primary.ID {
			return fmt.Errorf("Rate limit not found")
		}

		*rateLimit = foundRateLimit

		return nil
	}
}

func testAccCheckCloudflareRateLimitIDIsValid(n, expectedZoneID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No rate limit ID is set")
		}

		if len(rs.Primary.ID) != 32 {
			return fmt.Errorf("invalid id %q, should be a string with 32 characters", rs.Primary.ID)
		}

		if zoneID, ok := rs.Primary.Attributes[consts.ZoneIDSchemaKey]; !ok || len(zoneID) < 1 {
			return errors.New("zone_id is unset, should always be set with id")
		}

		if zoneID, _ := rs.Primary.Attributes[consts.ZoneIDSchemaKey]; zoneID != expectedZoneID {
			return fmt.Errorf("found zone_id value %q, expected %q", zoneID, expectedZoneID)
		}

		return nil
	}
}

func testAccManuallyDeleteRateLimit(name string, rateLimit *cloudflare.RateLimit, initialRateLimitId *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, clientErr := acctest.SharedV1Client() // TODO(terraform): replace with SharedV2Clent
		if clientErr != nil {
			tflog.Error(context.TODO(), fmt.Sprintf("failed to create Cloudflare client: %s", clientErr))
		}
		*initialRateLimitId = rateLimit.ID
		err := client.DeleteRateLimit(context.Background(), s.RootModule().Resources[name].Primary.Attributes[consts.ZoneIDSchemaKey], rateLimit.ID)
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccCheckCloudflareRateLimitConfigBasic(zoneID, id string) string {
	return acctest.LoadTestCase("ratelimitconfigbasic.tf", id, zoneID)
}

func testAccCheckCloudflareRateLimitConfigMatchingUrl(zoneID, id, zoneName string) string {
	return acctest.LoadTestCase("ratelimitconfigmatchingurl.tf", id, zoneID, zoneName)
}

func testAccCheckCloudflareRateLimitConfigFullySpecified(zoneID, id, zoneName string) string {
	return acctest.LoadTestCase("ratelimitconfigfullyspecified.tf", id, zoneID, zoneName)
}

func testAccCheckCloudflareRateLimitChallengeConfigBasic(zoneID, id string) string {
	return acctest.LoadTestCase("ratelimitchallengeconfigbasic.tf", id, zoneID)
}

func testAccCheckCloudflareRateLimitConfigWithoutTimeout(zoneID, id string) string {
	return acctest.LoadTestCase("ratelimitconfigwithouttimeout.tf", id, zoneID)
}

func testAccCheckCloudflareRateLimitChallengeConfigWithTimeout(zoneID, id string) string {
	return acctest.LoadTestCase("ratelimitchallengeconfigwithtimeout.tf", id, zoneID)
}
