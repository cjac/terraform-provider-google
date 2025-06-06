// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/iamworkforcepool/OauthClientCredential.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package iamworkforcepool

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIAMWorkforcePoolOauthClientCredential() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAMWorkforcePoolOauthClientCredentialCreate,
		Read:   resourceIAMWorkforcePoolOauthClientCredentialRead,
		Update: resourceIAMWorkforcePoolOauthClientCredentialUpdate,
		Delete: resourceIAMWorkforcePoolOauthClientCredentialDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIAMWorkforcePoolOauthClientCredentialImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.`,
			},
			"oauth_client_credential_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Required. The ID to use for the OauthClientCredential, which becomes the
final component of the resource name. This value should be 4-32 characters,
and may contain the characters [a-z0-9-]. The prefix 'gcp-' is
reserved for use by Google, and may not be specified.`,
			},
			"oauthclient": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in https://google.aip.dev/122.`,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Whether the OauthClientCredential is disabled. You cannot use a
disabled OauthClientCredential.`,
			},
			"display_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `A user-specified display name of the OauthClientCredential.

Cannot exceed 32 characters.`,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The system-generated OAuth client secret.

The client secret must be stored securely. If the client secret is
leaked, you must delete and re-create the client credential. To learn
more, see [OAuth client and credential security risks and
mitigations](https://cloud.google.com/iam/docs/workforce-oauth-app#security)`,
				Sensitive: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Immutable. Identifier. The resource name of the OauthClientCredential.

Format:
'projects/{project}/locations/{location}/oauthClients/{oauth_client}/credentials/{credential}'`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIAMWorkforcePoolOauthClientCredentialCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	disabledProp, err := expandIAMWorkforcePoolOauthClientCredentialDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	displayNameProp, err := expandIAMWorkforcePoolOauthClientCredentialDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}projects/{{project}}/locations/{{location}}/oauthClients/{{oauthclient}}/credentials?oauthClientCredentialId={{oauth_client_credential_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new OauthClientCredential: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for OauthClientCredential: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating OauthClientCredential: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/oauthClients/{{oauthclient}}/credentials/{{oauth_client_credential_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)

	log.Printf("[DEBUG] Finished creating OauthClientCredential %q: %#v", d.Id(), res)

	return resourceIAMWorkforcePoolOauthClientCredentialRead(d, meta)
}

func resourceIAMWorkforcePoolOauthClientCredentialRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}projects/{{project}}/locations/{{location}}/oauthClients/{{oauthclient}}/credentials/{{oauth_client_credential_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for OauthClientCredential: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IAMWorkforcePoolOauthClientCredential %q", d.Id()))
	}

	res, err = resourceIAMWorkforcePoolOauthClientCredentialDecoder(d, meta, res)
	if err != nil {
		return err
	}

	if res == nil {
		// Decoding the object has resulted in it being gone. It may be marked deleted
		log.Printf("[DEBUG] Removing IAMWorkforcePoolOauthClientCredential because it no longer exists.")
		d.SetId("")
		return nil
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading OauthClientCredential: %s", err)
	}

	if err := d.Set("disabled", flattenIAMWorkforcePoolOauthClientCredentialDisabled(res["disabled"], d, config)); err != nil {
		return fmt.Errorf("Error reading OauthClientCredential: %s", err)
	}
	if err := d.Set("client_secret", flattenIAMWorkforcePoolOauthClientCredentialClientSecret(res["clientSecret"], d, config)); err != nil {
		return fmt.Errorf("Error reading OauthClientCredential: %s", err)
	}
	if err := d.Set("display_name", flattenIAMWorkforcePoolOauthClientCredentialDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading OauthClientCredential: %s", err)
	}
	if err := d.Set("name", flattenIAMWorkforcePoolOauthClientCredentialName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading OauthClientCredential: %s", err)
	}

	return nil
}

func resourceIAMWorkforcePoolOauthClientCredentialUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for OauthClientCredential: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	disabledProp, err := expandIAMWorkforcePoolOauthClientCredentialDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disabled"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	displayNameProp, err := expandIAMWorkforcePoolOauthClientCredentialDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}projects/{{project}}/locations/{{location}}/oauthClients/{{oauthclient}}/credentials/{{oauth_client_credential_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating OauthClientCredential %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("disabled") {
		updateMask = append(updateMask, "disabled")
	}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating OauthClientCredential %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating OauthClientCredential %q: %#v", d.Id(), res)
		}

	}

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)
	return resourceIAMWorkforcePoolOauthClientCredentialRead(d, meta)
}

func resourceIAMWorkforcePoolOauthClientCredentialDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for OauthClientCredential: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{IAMWorkforcePoolBasePath}}projects/{{project}}/locations/{{location}}/oauthClients/{{oauthclient}}/credentials/{{oauth_client_credential_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting OauthClientCredential %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "OauthClientCredential")
	}

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)

	log.Printf("[DEBUG] Finished deleting OauthClientCredential %q: %#v", d.Id(), res)
	return nil
}

func resourceIAMWorkforcePoolOauthClientCredentialImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/oauthClients/(?P<oauthclient>[^/]+)/credentials/(?P<oauth_client_credential_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<oauthclient>[^/]+)/(?P<oauth_client_credential_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<oauthclient>[^/]+)/(?P<oauth_client_credential_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/oauthClients/{{oauthclient}}/credentials/{{oauth_client_credential_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIAMWorkforcePoolOauthClientCredentialDisabled(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolOauthClientCredentialClientSecret(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolOauthClientCredentialDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAMWorkforcePoolOauthClientCredentialName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIAMWorkforcePoolOauthClientCredentialDisabled(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAMWorkforcePoolOauthClientCredentialDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceIAMWorkforcePoolOauthClientCredentialDecoder(d *schema.ResourceData, meta interface{}, res map[string]interface{}) (map[string]interface{}, error) {
	if v := res["state"]; v == "DELETED" {
		return nil, nil
	}

	return res, nil
}
