// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
// ----------------------------------------------------------------------------
//
//	***     AUTO GENERATED CODE    ***    Type: Handwritten     ***
//
// ----------------------------------------------------------------------------
//
//	This code is generated by Magic Modules using the following:
//
//	Source file: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/third_party/terraform/services/dialogflowcx/resource_dialogflowcx_agent_test.go
//
//	DO NOT EDIT this file directly. Any changes made to this file will be
//	overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------
package dialogflowcx_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDialogflowCXAgent_update(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":          envvar.GetTestOrgFromEnv(t),
		"billing_account": envvar.GetTestBillingAccountFromEnv(t),
		"random_suffix":   acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXAgent_basic(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_agent.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"git_integration_settings.0.github_settings.0.access_token", "enable_stackdriver_logging", "advanced_settings.0.logging_settings"},
			},
			{
				Config: testAccDialogflowCXAgent_full(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_agent.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"git_integration_settings.0.github_settings.0.access_token", "enable_stackdriver_logging", "advanced_settings.0.logging_settings"},
			},
			{
				Config: testAccDialogflowCXAgent_removeSettings(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_agent.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"git_integration_settings.0.github_settings.0.access_token", "enable_stackdriver_logging", "advanced_settings.0.logging_settings"},
			},
		},
	})
}

func testAccDialogflowCXAgent_basic(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_dialogflow_cx_agent" "foobar" {
		display_name = "tf-test-%{random_suffix}"
		location = "global"
		default_language_code = "en"
		supported_language_codes = ["fr","de","es"]
		time_zone = "America/New_York"
		description = "Description 1."
		avatar_uri = "https://storage.cloud.google.com/dialogflow-test-host-image/cloud-logo.png"
	}
	`, context)
}

func testAccDialogflowCXAgent_full(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_storage_bucket" "bucket" {
		name                        = "tf-test-dialogflowcx-bucket%{random_suffix}"
		location                    = "US"
		uniform_bucket_level_access = true
	}

	resource "google_dialogflow_cx_agent" "foobar" {
		display_name = "tf-test-%{random_suffix}update"
		location = "global"
		default_language_code = "en"
		supported_language_codes = ["no"]
		time_zone = "Europe/London"
		description = "Description 2!"
		avatar_uri = "https://storage.cloud.google.com/dialogflow-test-host-image/cloud-logo-2.png"
		enable_stackdriver_logging = true
		enable_spell_correction    = true
		speech_to_text_settings {
			enable_speech_adaptation = true
		}
		advanced_settings {
			audio_export_gcs_destination {
				uri = "${google_storage_bucket.bucket.url}/prefix-"
			}
			speech_settings {
				endpointer_sensitivity        = 30
				no_speech_timeout             = "3.500s"
				use_timeout_based_endpointing = true
				models = {
				name : "wrench"
				mass : "1.3kg"
				count : "3"
				}
			}
			dtmf_settings {
				enabled      = true
				max_digits   = 1
				finish_digit = "#"
			}
			logging_settings {
				enable_stackdriver_logging     = true
				enable_interaction_logging     = true
				enable_consent_based_redaction = true
			}
		}
		git_integration_settings {
			github_settings {
				display_name = "Github Repo"
				repository_uri = "https://api.github.com/repos/githubtraining/hellogitworld"
				tracking_branch = "main"
				access_token = "secret-token"
				branches = ["main"]
			}
		}
		text_to_speech_settings {
			synthesize_speech_configs = jsonencode({
				en = {
					voice = {
						name = "en-US-Neural2-A"
					}
				}
			})
		}
	}
	  `, context)
}

func testAccDialogflowCXAgent_removeSettings(context map[string]interface{}) string {
	return acctest.Nprintf(`
	resource "google_dialogflow_cx_agent" "foobar" {
		display_name = "tf-test-%{random_suffix}"
		location = "global"
		default_language_code = "en"
		supported_language_codes = ["fr","de","es"]
		time_zone = "America/New_York"
		description = "Description 1."
		avatar_uri = "https://storage.cloud.google.com/dialogflow-test-host-image/cloud-logo.png"
		advanced_settings {}
		git_integration_settings {}
		text_to_speech_settings {}
	}
	  `, context)
}
