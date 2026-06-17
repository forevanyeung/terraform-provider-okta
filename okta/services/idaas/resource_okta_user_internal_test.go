package idaas

import (
	"context"
	"testing"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestValidateExpirePasswordRequiresPassword(t *testing.T) {
	cases := map[string]struct {
		config    cty.Value
		expectErr bool
	}{
		"expire true, no password set": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.True,
				"password":                  cty.NullVal(cty.String),
				"password_wo":               cty.NullVal(cty.String),
			}),
			expectErr: true,
		},
		"expire true, empty password": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.True,
				"password":                  cty.StringVal(""),
				"password_wo":               cty.NullVal(cty.String),
			}),
			expectErr: true,
		},
		"expire true, password set": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.True,
				"password":                  cty.StringVal("Abcd1234"),
				"password_wo":               cty.NullVal(cty.String),
			}),
			expectErr: false,
		},
		"expire true, password_wo set": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.True,
				"password":                  cty.NullVal(cty.String),
				"password_wo":               cty.StringVal("Abcd1234"),
			}),
			expectErr: false,
		},
		"expire true, password_wo unknown (e.g. ephemeral)": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.True,
				"password":                  cty.NullVal(cty.String),
				"password_wo":               cty.UnknownVal(cty.String),
			}),
			expectErr: false,
		},
		"expire false, no password": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.False,
				"password":                  cty.NullVal(cty.String),
				"password_wo":               cty.NullVal(cty.String),
			}),
			expectErr: false,
		},
		"expire null, no password": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.NullVal(cty.Bool),
				"password":                  cty.NullVal(cty.String),
				"password_wo":               cty.NullVal(cty.String),
			}),
			expectErr: false,
		},
		"expire unknown, no password": {
			config: cty.ObjectVal(map[string]cty.Value{
				"expire_password_on_create": cty.UnknownVal(cty.Bool),
				"password":                  cty.NullVal(cty.String),
				"password_wo":               cty.NullVal(cty.String),
			}),
			expectErr: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			req := schema.ValidateResourceConfigFuncRequest{RawConfig: tc.config}
			resp := &schema.ValidateResourceConfigFuncResponse{}
			validateExpirePasswordRequiresPassword(context.Background(), req, resp)
			if resp.Diagnostics.HasError() != tc.expectErr {
				t.Fatalf("expected error=%v, got diagnostics=%v", tc.expectErr, resp.Diagnostics)
			}
		})
	}
}
