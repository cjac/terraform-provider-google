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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/gemini/CodeRepositoryIndex.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package gemini

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

func ResourceGeminiCodeRepositoryIndex() *schema.Resource {
	return &schema.Resource{
		Create: resourceGeminiCodeRepositoryIndexCreate,
		Read:   resourceGeminiCodeRepositoryIndexRead,
		Update: resourceGeminiCodeRepositoryIndexUpdate,
		Delete: resourceGeminiCodeRepositoryIndexDelete,

		Importer: &schema.ResourceImporter{
			State: resourceGeminiCodeRepositoryIndexImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
			Update: schema.DefaultTimeout(90 * time.Minute),
			Delete: schema.DefaultTimeout(90 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"code_repository_index_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. Id of the Code Repository Index.`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location of the Code Repository Index, for example 'us-central1'.`,
			},
			"kms_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Optional. Immutable. Customer-managed encryption key name, in the format
'projects/*/locations/*/keyRings/*/cryptoKeys/*'.`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Optional. Labels as key value pairs.

**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Create time stamp.`,
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Immutable. Identifier. Name of Code Repository Index.`,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Code Repository Index instance State.
Possible values are: 'STATE_UNSPECIFIED', 'CREATING', 'ACTIVE', 'DELETING', 'SUSPENDED'.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. Update time stamp.`,
			},
			"force_destroy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `If set to true, will allow deletion of the CodeRepositoryIndex even if there are existing RepositoryGroups for the resource. These RepositoryGroups will also be deleted.`,
				Default:     false,
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

func resourceGeminiCodeRepositoryIndexCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	kmsKeyProp, err := expandGeminiCodeRepositoryIndexKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key"); !tpgresource.IsEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}
	labelsProp, err := expandGeminiCodeRepositoryIndexEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/codeRepositoryIndexes?codeRepositoryIndexId={{code_repository_index_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new CodeRepositoryIndex: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CodeRepositoryIndex: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "POST",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutCreate),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsCodeRepositoryIndexUnreadyError, transport_tpg.IsRepositoryGroupQueueError},
	})
	if err != nil {
		return fmt.Errorf("Error creating CodeRepositoryIndex: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = GeminiOperationWaitTime(
		config, res, project, "Creating CodeRepositoryIndex", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create CodeRepositoryIndex: %s", err)
	}

	log.Printf("[DEBUG] Finished creating CodeRepositoryIndex %q: %#v", d.Id(), res)

	return resourceGeminiCodeRepositoryIndexRead(d, meta)
}

func resourceGeminiCodeRepositoryIndexRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CodeRepositoryIndex: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "GET",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsCodeRepositoryIndexUnreadyError, transport_tpg.IsRepositoryGroupQueueError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("GeminiCodeRepositoryIndex %q", d.Id()))
	}

	// Explicitly set virtual fields to default values if unset
	if _, ok := d.GetOkExists("force_destroy"); !ok {
		if err := d.Set("force_destroy", false); err != nil {
			return fmt.Errorf("Error setting force_destroy: %s", err)
		}
	}
	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}

	if err := d.Set("update_time", flattenGeminiCodeRepositoryIndexUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}
	if err := d.Set("state", flattenGeminiCodeRepositoryIndexState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}
	if err := d.Set("labels", flattenGeminiCodeRepositoryIndexLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}
	if err := d.Set("kms_key", flattenGeminiCodeRepositoryIndexKmsKey(res["kmsKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}
	if err := d.Set("name", flattenGeminiCodeRepositoryIndexName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}
	if err := d.Set("create_time", flattenGeminiCodeRepositoryIndexCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}
	if err := d.Set("terraform_labels", flattenGeminiCodeRepositoryIndexTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}
	if err := d.Set("effective_labels", flattenGeminiCodeRepositoryIndexEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading CodeRepositoryIndex: %s", err)
	}

	return nil
}

func resourceGeminiCodeRepositoryIndexUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CodeRepositoryIndex: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	labelsProp, err := expandGeminiCodeRepositoryIndexEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating CodeRepositoryIndex %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
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
			Config:               config,
			Method:               "PATCH",
			Project:              billingProject,
			RawURL:               url,
			UserAgent:            userAgent,
			Body:                 obj,
			Timeout:              d.Timeout(schema.TimeoutUpdate),
			Headers:              headers,
			ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsCodeRepositoryIndexUnreadyError, transport_tpg.IsRepositoryGroupQueueError},
		})

		if err != nil {
			return fmt.Errorf("Error updating CodeRepositoryIndex %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating CodeRepositoryIndex %q: %#v", d.Id(), res)
		}

		err = GeminiOperationWaitTime(
			config, res, project, "Updating CodeRepositoryIndex", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceGeminiCodeRepositoryIndexRead(d, meta)
}

func resourceGeminiCodeRepositoryIndexDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for CodeRepositoryIndex: %s", err)
	}
	billingProject = project

	lockName, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}")
	if err != nil {
		return err
	}
	transport_tpg.MutexStore.Lock(lockName)
	defer transport_tpg.MutexStore.Unlock(lockName)

	url, err := tpgresource.ReplaceVars(d, config, "{{GeminiBasePath}}projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}?force={{force_destroy}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting CodeRepositoryIndex %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:               config,
		Method:               "DELETE",
		Project:              billingProject,
		RawURL:               url,
		UserAgent:            userAgent,
		Body:                 obj,
		Timeout:              d.Timeout(schema.TimeoutDelete),
		Headers:              headers,
		ErrorRetryPredicates: []transport_tpg.RetryErrorPredicateFunc{transport_tpg.IsCodeRepositoryIndexUnreadyError, transport_tpg.IsRepositoryGroupQueueError},
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "CodeRepositoryIndex")
	}

	err = GeminiOperationWaitTime(
		config, res, project, "Deleting CodeRepositoryIndex", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting CodeRepositoryIndex %q: %#v", d.Id(), res)
	return nil
}

func resourceGeminiCodeRepositoryIndexImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/codeRepositoryIndexes/(?P<code_repository_index_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<code_repository_index_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<code_repository_index_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/codeRepositoryIndexes/{{code_repository_index_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Explicitly set virtual fields to default values on import
	if err := d.Set("force_destroy", false); err != nil {
		return nil, fmt.Errorf("Error setting force_destroy: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenGeminiCodeRepositoryIndexUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiCodeRepositoryIndexState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiCodeRepositoryIndexLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGeminiCodeRepositoryIndexKmsKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiCodeRepositoryIndexName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiCodeRepositoryIndexCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenGeminiCodeRepositoryIndexTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenGeminiCodeRepositoryIndexEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandGeminiCodeRepositoryIndexKmsKey(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandGeminiCodeRepositoryIndexEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
