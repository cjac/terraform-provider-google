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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dialogflowcx/SecuritySettings.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dialogflowcx

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
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceDialogflowCXSecuritySettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceDialogflowCXSecuritySettingsCreate,
		Read:   resourceDialogflowCXSecuritySettingsRead,
		Update: resourceDialogflowCXSecuritySettingsUpdate,
		Delete: resourceDialogflowCXSecuritySettingsDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDialogflowCXSecuritySettingsImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(40 * time.Minute),
			Update: schema.DefaultTimeout(40 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"display_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `The human-readable name of the security settings, unique within the location.`,
			},
			"location": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The location these settings are located in. Settings can only be applied to an agent in the same location.
See [Available Regions](https://cloud.google.com/dialogflow/cx/docs/concept/region#avail) for a list of supported locations.`,
			},
			"audio_export_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Controls audio export settings for post-conversation analytics when ingesting audio to conversations.
If retention_strategy is set to REMOVE_AFTER_CONVERSATION or gcs_bucket is empty, audio export is disabled.
If audio export is enabled, audio is recorded and saved to gcs_bucket, subject to retention policy of gcs_bucket.
This setting won't effect audio input for implicit sessions via [Sessions.DetectIntent](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.sessions/detectIntent#google.cloud.dialogflow.cx.v3.Sessions.DetectIntent).`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"audio_export_pattern": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Filename pattern for exported audio.`,
						},
						"audio_format": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: verify.ValidateEnum([]string{"MULAW", "MP3", "OGG", ""}),
							Description: `File format for exported audio file. Currently only in telephony recordings.
* MULAW: G.711 mu-law PCM with 8kHz sample rate.
* MP3: MP3 file format.
* OGG: OGG Vorbis. Possible values: ["MULAW", "MP3", "OGG"]`,
						},
						"enable_audio_redaction": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: `Enable audio redaction if it is true.`,
						},
						"gcs_bucket": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: `Cloud Storage bucket to export audio record to. Setting this field would grant the Storage Object Creator role to the Dialogflow Service Agent. API caller that tries to modify this field should have the permission of storage.buckets.setIamPolicy.`,
						},
					},
				},
			},
			"deidentify_template": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `[DLP](https://cloud.google.com/dlp/docs) deidentify template name. Use this template to define de-identification configuration for the content. If empty, Dialogflow replaces sensitive info with [redacted] text.
Note: deidentifyTemplate must be located in the same region as the SecuritySettings.
Format: projects/<Project ID>/locations/<Location ID>/deidentifyTemplates/<Template ID> OR organizations/<Organization ID>/locations/<Location ID>/deidentifyTemplates/<Template ID>`,
			},
			"insights_export_settings": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Controls conversation exporting settings to Insights after conversation is completed.
If retentionStrategy is set to REMOVE_AFTER_CONVERSATION, Insights export is disabled no matter what you configure here.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_insights_export": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: `If enabled, we will automatically exports conversations to Insights and Insights runs its analyzers.`,
						},
					},
				},
			},
			"inspect_template": {
				Type:     schema.TypeString,
				Optional: true,
				Description: `[DLP](https://cloud.google.com/dlp/docs) inspect template name. Use this template to define inspect base settings. If empty, we use the default DLP inspect config.
Note: inspectTemplate must be located in the same region as the SecuritySettings.
Format: projects/<Project ID>/locations/<Location ID>/inspectTemplates/<Template ID> OR organizations/<Organization ID>/locations/<Location ID>/inspectTemplates/<Template ID>`,
			},
			"purge_data_types": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `List of types of data to remove when retention settings triggers purge. Possible values: ["DIALOGFLOW_HISTORY"]`,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: verify.ValidateEnum([]string{"DIALOGFLOW_HISTORY"}),
				},
			},
			"redaction_scope": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"REDACT_DISK_STORAGE", ""}),
				Description: `Defines what types of data to redact. If not set, defaults to not redacting any kind of data.
* REDACT_DISK_STORAGE: On data to be written to disk or similar devices that are capable of holding data even if power is disconnected. This includes data that are temporarily saved on disk. Possible values: ["REDACT_DISK_STORAGE"]`,
			},
			"redaction_strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"REDACT_WITH_SERVICE", ""}),
				Description: `Defines how we redact data. If not set, defaults to not redacting.
* REDACT_WITH_SERVICE: Call redaction service to clean up the data to be persisted. Possible values: ["REDACT_WITH_SERVICE"]`,
			},
			"retention_strategy": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"REMOVE_AFTER_CONVERSATION", ""}),
				Description: `Defines how long we retain persisted data that contains sensitive info. Only one of 'retention_window_days' and 'retention_strategy' may be set.
* REMOVE_AFTER_CONVERSATION: Removes data when the conversation ends. If there is no conversation explicitly established, a default conversation ends when the corresponding Dialogflow session ends. Possible values: ["REMOVE_AFTER_CONVERSATION"]`,
				ConflictsWith: []string{"retention_window_days"},
			},
			"retention_window_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: `Retains the data for the specified number of days. User must set a value lower than Dialogflow's default 365d TTL (30 days for Agent Assist traffic), higher value will be ignored and use default. Setting a value higher than that has no effect. A missing value or setting to 0 also means we use default TTL.
Only one of 'retention_window_days' and 'retention_strategy' may be set.`,
				ConflictsWith: []string{"retention_strategy"},
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The unique identifier of the settings.
Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.`,
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

func resourceDialogflowCXSecuritySettingsCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXSecuritySettingsDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	redactionStrategyProp, err := expandDialogflowCXSecuritySettingsRedactionStrategy(d.Get("redaction_strategy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redaction_strategy"); !tpgresource.IsEmptyValue(reflect.ValueOf(redactionStrategyProp)) && (ok || !reflect.DeepEqual(v, redactionStrategyProp)) {
		obj["redactionStrategy"] = redactionStrategyProp
	}
	redactionScopeProp, err := expandDialogflowCXSecuritySettingsRedactionScope(d.Get("redaction_scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redaction_scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(redactionScopeProp)) && (ok || !reflect.DeepEqual(v, redactionScopeProp)) {
		obj["redactionScope"] = redactionScopeProp
	}
	inspectTemplateProp, err := expandDialogflowCXSecuritySettingsInspectTemplate(d.Get("inspect_template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("inspect_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(inspectTemplateProp)) && (ok || !reflect.DeepEqual(v, inspectTemplateProp)) {
		obj["inspectTemplate"] = inspectTemplateProp
	}
	deidentifyTemplateProp, err := expandDialogflowCXSecuritySettingsDeidentifyTemplate(d.Get("deidentify_template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("deidentify_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(deidentifyTemplateProp)) && (ok || !reflect.DeepEqual(v, deidentifyTemplateProp)) {
		obj["deidentifyTemplate"] = deidentifyTemplateProp
	}
	purgeDataTypesProp, err := expandDialogflowCXSecuritySettingsPurgeDataTypes(d.Get("purge_data_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("purge_data_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(purgeDataTypesProp)) && (ok || !reflect.DeepEqual(v, purgeDataTypesProp)) {
		obj["purgeDataTypes"] = purgeDataTypesProp
	}
	audioExportSettingsProp, err := expandDialogflowCXSecuritySettingsAudioExportSettings(d.Get("audio_export_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("audio_export_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(audioExportSettingsProp)) && (ok || !reflect.DeepEqual(v, audioExportSettingsProp)) {
		obj["audioExportSettings"] = audioExportSettingsProp
	}
	insightsExportSettingsProp, err := expandDialogflowCXSecuritySettingsInsightsExportSettings(d.Get("insights_export_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("insights_export_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(insightsExportSettingsProp)) && (ok || !reflect.DeepEqual(v, insightsExportSettingsProp)) {
		obj["insightsExportSettings"] = insightsExportSettingsProp
	}
	retentionWindowDaysProp, err := expandDialogflowCXSecuritySettingsRetentionWindowDays(d.Get("retention_window_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention_window_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionWindowDaysProp)) && (ok || !reflect.DeepEqual(v, retentionWindowDaysProp)) {
		obj["retentionWindowDays"] = retentionWindowDaysProp
	}
	retentionStrategyProp, err := expandDialogflowCXSecuritySettingsRetentionStrategy(d.Get("retention_strategy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention_strategy"); !tpgresource.IsEmptyValue(reflect.ValueOf(retentionStrategyProp)) && (ok || !reflect.DeepEqual(v, retentionStrategyProp)) {
		obj["retentionStrategy"] = retentionStrategyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/securitySettings")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SecuritySettings: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecuritySettings: %s", err)
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
		return fmt.Errorf("Error creating SecuritySettings: %s", err)
	}
	// Set computed resource properties from create API response so that they're available on the subsequent Read
	// call.
	err = resourceDialogflowCXSecuritySettingsPostCreateSetComputedFields(d, meta, res)
	if err != nil {
		return fmt.Errorf("setting computed ID format fields: %w", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/securitySettings/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)

	log.Printf("[DEBUG] Finished creating SecuritySettings %q: %#v", d.Id(), res)

	return resourceDialogflowCXSecuritySettingsRead(d, meta)
}

func resourceDialogflowCXSecuritySettingsRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/securitySettings/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecuritySettings: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DialogflowCXSecuritySettings %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}

	if err := d.Set("name", flattenDialogflowCXSecuritySettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("display_name", flattenDialogflowCXSecuritySettingsDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("redaction_strategy", flattenDialogflowCXSecuritySettingsRedactionStrategy(res["redactionStrategy"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("redaction_scope", flattenDialogflowCXSecuritySettingsRedactionScope(res["redactionScope"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("inspect_template", flattenDialogflowCXSecuritySettingsInspectTemplate(res["inspectTemplate"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("deidentify_template", flattenDialogflowCXSecuritySettingsDeidentifyTemplate(res["deidentifyTemplate"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("purge_data_types", flattenDialogflowCXSecuritySettingsPurgeDataTypes(res["purgeDataTypes"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("audio_export_settings", flattenDialogflowCXSecuritySettingsAudioExportSettings(res["audioExportSettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("insights_export_settings", flattenDialogflowCXSecuritySettingsInsightsExportSettings(res["insightsExportSettings"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("retention_window_days", flattenDialogflowCXSecuritySettingsRetentionWindowDays(res["retentionWindowDays"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}
	if err := d.Set("retention_strategy", flattenDialogflowCXSecuritySettingsRetentionStrategy(res["retentionStrategy"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecuritySettings: %s", err)
	}

	return nil
}

func resourceDialogflowCXSecuritySettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecuritySettings: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXSecuritySettingsDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	redactionStrategyProp, err := expandDialogflowCXSecuritySettingsRedactionStrategy(d.Get("redaction_strategy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redaction_strategy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, redactionStrategyProp)) {
		obj["redactionStrategy"] = redactionStrategyProp
	}
	redactionScopeProp, err := expandDialogflowCXSecuritySettingsRedactionScope(d.Get("redaction_scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("redaction_scope"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, redactionScopeProp)) {
		obj["redactionScope"] = redactionScopeProp
	}
	inspectTemplateProp, err := expandDialogflowCXSecuritySettingsInspectTemplate(d.Get("inspect_template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("inspect_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, inspectTemplateProp)) {
		obj["inspectTemplate"] = inspectTemplateProp
	}
	deidentifyTemplateProp, err := expandDialogflowCXSecuritySettingsDeidentifyTemplate(d.Get("deidentify_template"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("deidentify_template"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, deidentifyTemplateProp)) {
		obj["deidentifyTemplate"] = deidentifyTemplateProp
	}
	purgeDataTypesProp, err := expandDialogflowCXSecuritySettingsPurgeDataTypes(d.Get("purge_data_types"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("purge_data_types"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, purgeDataTypesProp)) {
		obj["purgeDataTypes"] = purgeDataTypesProp
	}
	audioExportSettingsProp, err := expandDialogflowCXSecuritySettingsAudioExportSettings(d.Get("audio_export_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("audio_export_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, audioExportSettingsProp)) {
		obj["audioExportSettings"] = audioExportSettingsProp
	}
	insightsExportSettingsProp, err := expandDialogflowCXSecuritySettingsInsightsExportSettings(d.Get("insights_export_settings"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("insights_export_settings"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, insightsExportSettingsProp)) {
		obj["insightsExportSettings"] = insightsExportSettingsProp
	}
	retentionWindowDaysProp, err := expandDialogflowCXSecuritySettingsRetentionWindowDays(d.Get("retention_window_days"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention_window_days"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retentionWindowDaysProp)) {
		obj["retentionWindowDays"] = retentionWindowDaysProp
	}
	retentionStrategyProp, err := expandDialogflowCXSecuritySettingsRetentionStrategy(d.Get("retention_strategy"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("retention_strategy"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, retentionStrategyProp)) {
		obj["retentionStrategy"] = retentionStrategyProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/securitySettings/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating SecuritySettings %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("redaction_strategy") {
		updateMask = append(updateMask, "redactionStrategy")
	}

	if d.HasChange("redaction_scope") {
		updateMask = append(updateMask, "redactionScope")
	}

	if d.HasChange("inspect_template") {
		updateMask = append(updateMask, "inspectTemplate")
	}

	if d.HasChange("deidentify_template") {
		updateMask = append(updateMask, "deidentifyTemplate")
	}

	if d.HasChange("purge_data_types") {
		updateMask = append(updateMask, "purgeDataTypes")
	}

	if d.HasChange("audio_export_settings") {
		updateMask = append(updateMask, "audioExportSettings")
	}

	if d.HasChange("insights_export_settings") {
		updateMask = append(updateMask, "insightsExportSettings")
	}

	if d.HasChange("retention_window_days") {
		updateMask = append(updateMask, "retentionWindowDays")
	}

	if d.HasChange("retention_strategy") {
		updateMask = append(updateMask, "retentionStrategy")
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
			return fmt.Errorf("Error updating SecuritySettings %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating SecuritySettings %q: %#v", d.Id(), res)
		}

	}

	// This is useful if the resource in question doesn't have a perfectly consistent API
	// That is, the Operation for Create might return before the Get operation shows the
	// completed state of the resource.
	time.Sleep(5 * time.Second)
	return resourceDialogflowCXSecuritySettingsRead(d, meta)
}

func resourceDialogflowCXSecuritySettingsDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for SecuritySettings: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DialogflowCXBasePath}}projects/{{project}}/locations/{{location}}/securitySettings/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting SecuritySettings %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "SecuritySettings")
	}

	log.Printf("[DEBUG] Finished deleting SecuritySettings %q: %#v", d.Id(), res)
	return nil
}

func resourceDialogflowCXSecuritySettingsImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/securitySettings/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/securitySettings/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDialogflowCXSecuritySettingsName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.GetResourceNameFromSelfLink(v.(string))
}

func flattenDialogflowCXSecuritySettingsDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsRedactionStrategy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsRedactionScope(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsInspectTemplate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsDeidentifyTemplate(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsPurgeDataTypes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsAudioExportSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["gcs_bucket"] =
		flattenDialogflowCXSecuritySettingsAudioExportSettingsGcsBucket(original["gcsBucket"], d, config)
	transformed["audio_export_pattern"] =
		flattenDialogflowCXSecuritySettingsAudioExportSettingsAudioExportPattern(original["audioExportPattern"], d, config)
	transformed["enable_audio_redaction"] =
		flattenDialogflowCXSecuritySettingsAudioExportSettingsEnableAudioRedaction(original["enableAudioRedaction"], d, config)
	transformed["audio_format"] =
		flattenDialogflowCXSecuritySettingsAudioExportSettingsAudioFormat(original["audioFormat"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXSecuritySettingsAudioExportSettingsGcsBucket(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsAudioExportSettingsAudioExportPattern(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsAudioExportSettingsEnableAudioRedaction(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsAudioExportSettingsAudioFormat(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsInsightsExportSettings(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["enable_insights_export"] =
		flattenDialogflowCXSecuritySettingsInsightsExportSettingsEnableInsightsExport(original["enableInsightsExport"], d, config)
	return []interface{}{transformed}
}
func flattenDialogflowCXSecuritySettingsInsightsExportSettingsEnableInsightsExport(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDialogflowCXSecuritySettingsRetentionWindowDays(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenDialogflowCXSecuritySettingsRetentionStrategy(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDialogflowCXSecuritySettingsDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRedactionStrategy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRedactionScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsInspectTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsDeidentifyTemplate(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsPurgeDataTypes(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedGcsBucket, err := expandDialogflowCXSecuritySettingsAudioExportSettingsGcsBucket(original["gcs_bucket"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedGcsBucket); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["gcsBucket"] = transformedGcsBucket
	}

	transformedAudioExportPattern, err := expandDialogflowCXSecuritySettingsAudioExportSettingsAudioExportPattern(original["audio_export_pattern"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudioExportPattern); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audioExportPattern"] = transformedAudioExportPattern
	}

	transformedEnableAudioRedaction, err := expandDialogflowCXSecuritySettingsAudioExportSettingsEnableAudioRedaction(original["enable_audio_redaction"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableAudioRedaction); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableAudioRedaction"] = transformedEnableAudioRedaction
	}

	transformedAudioFormat, err := expandDialogflowCXSecuritySettingsAudioExportSettingsAudioFormat(original["audio_format"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAudioFormat); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["audioFormat"] = transformedAudioFormat
	}

	return transformed, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsGcsBucket(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsAudioExportPattern(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsEnableAudioRedaction(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsAudioExportSettingsAudioFormat(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsInsightsExportSettings(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableInsightsExport, err := expandDialogflowCXSecuritySettingsInsightsExportSettingsEnableInsightsExport(original["enable_insights_export"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableInsightsExport); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enableInsightsExport"] = transformedEnableInsightsExport
	}

	return transformed, nil
}

func expandDialogflowCXSecuritySettingsInsightsExportSettingsEnableInsightsExport(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRetentionWindowDays(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXSecuritySettingsRetentionStrategy(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func resourceDialogflowCXSecuritySettingsPostCreateSetComputedFields(d *schema.ResourceData, meta interface{}, res map[string]interface{}) error {
	config := meta.(*transport_tpg.Config)
	if err := d.Set("name", flattenDialogflowCXSecuritySettingsName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}
	return nil
}
