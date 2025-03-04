// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/stream"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestStreamModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*stream.StreamModel)(nil)
	schema := stream.ResourceSchema(context.TODO())
	errs := test_helpers.ValidateResourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
