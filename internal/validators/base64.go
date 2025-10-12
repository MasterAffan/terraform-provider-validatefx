package validators

import (
	"context"
	"encoding/base64"
	"fmt"

	frameworkvalidator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ frameworkvalidator.String = Base64Validator()

func Base64Validator() frameworkvalidator.String {
	return base64Validator{}
}

type base64Validator struct{}

func (base64Validator) Description(_ context.Context) string {
	return "value must be a base 64 string"
}

func (base64Validator) MarkdownDescription(_ context.Context) string {
	return "value must be a base 64 string"
}

func (base64Validator) ValidateString(_ context.Context, req frameworkvalidator.StringRequest, resp *frameworkvalidator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() || req.ConfigValue.ValueString() == "" {
		return
	}

	if _, err := base64.StdEncoding.DecodeString(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid base64 string",
			fmt.Sprintf("Value %q is not a valid base64 string", req.ConfigValue.ValueString()),
		)
	}

}
