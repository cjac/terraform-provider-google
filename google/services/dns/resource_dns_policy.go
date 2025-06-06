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
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/dns/Policy.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package dns

import (
	"bytes"
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

func ResourceDNSPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSPolicyCreate,
		Read:   resourceDNSPolicyRead,
		Update: resourceDNSPolicyUpdate,
		Delete: resourceDNSPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceDNSPolicyImport,
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
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `User assigned name for this policy.`,
			},
			"alternative_name_server_config": {
				Type:     schema.TypeList,
				Optional: true,
				Description: `Sets an alternative name server for the associated networks.
When specified, all DNS queries are forwarded to a name server that you choose.
Names such as .internal are not available when an alternative name server is specified.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_name_servers": {
							Type:     schema.TypeSet,
							Required: true,
							Description: `Sets an alternative name server for the associated networks. When specified,
all DNS queries are forwarded to a name server that you choose. Names such as .internal
are not available when an alternative name server is specified.`,
							Elem: dnsPolicyAlternativeNameServerConfigTargetNameServersSchema(),
							Set: func(v interface{}) int {
								raw := v.(map[string]interface{})
								if address, ok := raw["ipv4_address"]; ok {
									tpgresource.Hashcode(address.(string))
								}
								var buf bytes.Buffer
								schema.SerializeResourceForHash(&buf, raw, dnsPolicyAlternativeNameServerConfigTargetNameServersSchema())
								return tpgresource.Hashcode(buf.String())
							},
						},
					},
				},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A textual description field. Defaults to 'Managed by Terraform'.`,
				Default:     "Managed by Terraform",
			},
			"dns64_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				ForceNew:    true,
				Description: `Configurations related to DNS64 for this Policy.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"scope": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: `The scope to which DNS64 config will be applied to.`,
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"all_queries": {
										Type:        schema.TypeBool,
										Optional:    true,
										ForceNew:    true,
										Description: `Controls whether DNS64 is enabled globally at the network level.`,
									},
								},
							},
						},
					},
				},
			},
			"enable_inbound_forwarding": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Allows networks bound to this policy to receive DNS queries sent
by VMs or applications over VPN connections. When enabled, a
virtual IP address will be allocated from each of the sub-networks
that are bound to this policy.`,
			},
			"enable_logging": {
				Type:     schema.TypeBool,
				Optional: true,
				Description: `Controls whether logging is enabled for the networks bound to this policy.
Defaults to no logging if not set.`,
			},
			"networks": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: `List of network names specifying networks to which this policy is applied.`,
				Elem:        dnsPolicyNetworksSchema(),
				Set: func(v interface{}) int {
					raw := v.(map[string]interface{})
					if url, ok := raw["network_url"]; ok {
						return tpgresource.SelfLinkRelativePathHash(url)
					}
					var buf bytes.Buffer
					schema.SerializeResourceForHash(&buf, raw, dnsPolicyNetworksSchema())
					return tpgresource.Hashcode(buf.String())
				},
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

func dnsPolicyAlternativeNameServerConfigTargetNameServersSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"ipv4_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: `IPv4 address to forward to.`,
			},
			"forwarding_path": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"default", "private", ""}),
				Description: `Forwarding path for this TargetNameServer. If unset or 'default' Cloud DNS will make forwarding
decision based on address ranges, i.e. RFC1918 addresses go to the VPC, Non-RFC1918 addresses go
to the Internet. When set to 'private', Cloud DNS will always send queries through VPC for this target Possible values: ["default", "private"]`,
			},
		},
	}
}

func dnsPolicyNetworksSchema() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"network_url": {
				Type:             schema.TypeString,
				Required:         true,
				DiffSuppressFunc: tpgresource.CompareSelfLinkOrResourceName,
				Description: `The id or fully qualified URL of the VPC network to forward queries to.
This should be formatted like 'projects/{project}/global/networks/{network}' or
'https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}'`,
			},
		},
	}
}

func resourceDNSPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	alternativeNameServerConfigProp, err := expandDNSPolicyAlternativeNameServerConfig(d.Get("alternative_name_server_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("alternative_name_server_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(alternativeNameServerConfigProp)) && (ok || !reflect.DeepEqual(v, alternativeNameServerConfigProp)) {
		obj["alternativeNameServerConfig"] = alternativeNameServerConfigProp
	}
	descriptionProp, err := expandDNSPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	dns64ConfigProp, err := expandDNSPolicyDns64Config(d.Get("dns64_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("dns64_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(dns64ConfigProp)) && (ok || !reflect.DeepEqual(v, dns64ConfigProp)) {
		obj["dns64Config"] = dns64ConfigProp
	}
	enableInboundForwardingProp, err := expandDNSPolicyEnableInboundForwarding(d.Get("enable_inbound_forwarding"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_inbound_forwarding"); ok || !reflect.DeepEqual(v, enableInboundForwardingProp) {
		obj["enableInboundForwarding"] = enableInboundForwardingProp
	}
	enableLoggingProp, err := expandDNSPolicyEnableLogging(d.Get("enable_logging"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
		obj["enableLogging"] = enableLoggingProp
	}
	nameProp, err := expandDNSPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	networksProp, err := expandDNSPolicyNetworks(d.Get("networks"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("networks"); !tpgresource.IsEmptyValue(reflect.ValueOf(networksProp)) && (ok || !reflect.DeepEqual(v, networksProp)) {
		obj["networks"] = networksProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Policy: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
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
		return fmt.Errorf("Error creating Policy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/policies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Policy %q: %#v", d.Id(), res)

	return resourceDNSPolicyRead(d, meta)
}

func resourceDNSPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
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
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("DNSPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	if err := d.Set("alternative_name_server_config", flattenDNSPolicyAlternativeNameServerConfig(res["alternativeNameServerConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("description", flattenDNSPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("dns64_config", flattenDNSPolicyDns64Config(res["dns64Config"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("enable_inbound_forwarding", flattenDNSPolicyEnableInboundForwarding(res["enableInboundForwarding"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("enable_logging", flattenDNSPolicyEnableLogging(res["enableLogging"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("name", flattenDNSPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}
	if err := d.Set("networks", flattenDNSPolicyNetworks(res["networks"], d, config)); err != nil {
		return fmt.Errorf("Error reading Policy: %s", err)
	}

	return nil
}

func resourceDNSPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("alternative_name_server_config") || d.HasChange("description") || d.HasChange("enable_inbound_forwarding") || d.HasChange("enable_logging") || d.HasChange("networks") {
		obj := make(map[string]interface{})

		alternativeNameServerConfigProp, err := expandDNSPolicyAlternativeNameServerConfig(d.Get("alternative_name_server_config"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("alternative_name_server_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, alternativeNameServerConfigProp)) {
			obj["alternativeNameServerConfig"] = alternativeNameServerConfigProp
		}
		descriptionProp, err := expandDNSPolicyDescription(d.Get("description"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
			obj["description"] = descriptionProp
		}
		enableInboundForwardingProp, err := expandDNSPolicyEnableInboundForwarding(d.Get("enable_inbound_forwarding"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("enable_inbound_forwarding"); ok || !reflect.DeepEqual(v, enableInboundForwardingProp) {
			obj["enableInboundForwarding"] = enableInboundForwardingProp
		}
		enableLoggingProp, err := expandDNSPolicyEnableLogging(d.Get("enable_logging"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("enable_logging"); ok || !reflect.DeepEqual(v, enableLoggingProp) {
			obj["enableLogging"] = enableLoggingProp
		}
		networksProp, err := expandDNSPolicyNetworks(d.Get("networks"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("networks"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, networksProp)) {
			obj["networks"] = networksProp
		}

		url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
		if err != nil {
			return err
		}

		headers := make(http.Header)

		// err == nil indicates that the billing_project value was found
		if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
			billingProject = bp
		}

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
			return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Policy %q: %#v", d.Id(), res)
		}

	}

	d.Partial(false)

	return resourceDNSPolicyRead(d, meta)
}

func resourceDNSPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Policy: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	// if networks are attached, they need to be detached before the policy can be deleted
	if d.Get("networks.#").(int) > 0 {
		patched := make(map[string]interface{})
		patched["networks"] = nil

		url, err := tpgresource.ReplaceVars(d, config, "{{DNSBasePath}}projects/{{project}}/policies/{{name}}")
		if err != nil {
			return err
		}

		_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   project,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      patched,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
		})
		if err != nil {
			return fmt.Errorf("Error updating Policy %q: %s", d.Id(), err)
		}
	}

	log.Printf("[DEBUG] Deleting Policy %q", d.Id())
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
		return transport_tpg.HandleNotFoundError(err, d, "Policy")
	}

	log.Printf("[DEBUG] Finished deleting Policy %q: %#v", d.Id(), res)
	return nil
}

func resourceDNSPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/policies/(?P<name>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<name>[^/]+)$",
		"^(?P<name>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/policies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenDNSPolicyAlternativeNameServerConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["target_name_servers"] =
		flattenDNSPolicyAlternativeNameServerConfigTargetNameServers(original["targetNameServers"], d, config)
	return []interface{}{transformed}
}
func flattenDNSPolicyAlternativeNameServerConfigTargetNameServers(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(func(v interface{}) int {
		raw := v.(map[string]interface{})
		if address, ok := raw["ipv4_address"]; ok {
			tpgresource.Hashcode(address.(string))
		}
		var buf bytes.Buffer
		schema.SerializeResourceForHash(&buf, raw, dnsPolicyAlternativeNameServerConfigTargetNameServersSchema())
		return tpgresource.Hashcode(buf.String())
	}, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"ipv4_address":    flattenDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(original["ipv4Address"], d, config),
			"forwarding_path": flattenDNSPolicyAlternativeNameServerConfigTargetNameServersForwardingPath(original["forwardingPath"], d, config),
		})
	}
	return transformed
}
func flattenDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSPolicyAlternativeNameServerConfigTargetNameServersForwardingPath(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSPolicyDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSPolicyDns64Config(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["scope"] =
		flattenDNSPolicyDns64ConfigScope(original["scope"], d, config)
	return []interface{}{transformed}
}
func flattenDNSPolicyDns64ConfigScope(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["all_queries"] =
		flattenDNSPolicyDns64ConfigScopeAllQueries(original["allQueries"], d, config)
	return []interface{}{transformed}
}
func flattenDNSPolicyDns64ConfigScopeAllQueries(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSPolicyEnableInboundForwarding(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSPolicyEnableLogging(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenDNSPolicyNetworks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(func(v interface{}) int {
		raw := v.(map[string]interface{})
		if url, ok := raw["network_url"]; ok {
			return tpgresource.SelfLinkRelativePathHash(url)
		}
		var buf bytes.Buffer
		schema.SerializeResourceForHash(&buf, raw, dnsPolicyNetworksSchema())
		return tpgresource.Hashcode(buf.String())
	}, []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"network_url": flattenDNSPolicyNetworksNetworkUrl(original["networkUrl"], d, config),
		})
	}
	return transformed
}
func flattenDNSPolicyNetworksNetworkUrl(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandDNSPolicyAlternativeNameServerConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedTargetNameServers, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServers(original["target_name_servers"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTargetNameServers); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["targetNameServers"] = transformedTargetNameServers
	}

	return transformed, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServers(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedIpv4Address, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(original["ipv4_address"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedIpv4Address); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["ipv4Address"] = transformedIpv4Address
		}

		transformedForwardingPath, err := expandDNSPolicyAlternativeNameServerConfigTargetNameServersForwardingPath(original["forwarding_path"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedForwardingPath); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["forwardingPath"] = transformedForwardingPath
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServersIpv4Address(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyAlternativeNameServerConfigTargetNameServersForwardingPath(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyDns64Config(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedScope, err := expandDNSPolicyDns64ConfigScope(original["scope"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedScope); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["scope"] = transformedScope
	}

	return transformed, nil
}

func expandDNSPolicyDns64ConfigScope(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedAllQueries, err := expandDNSPolicyDns64ConfigScopeAllQueries(original["all_queries"], d, config)
	if err != nil {
		return nil, err
	} else {
		transformed["allQueries"] = transformedAllQueries
	}

	return transformed, nil
}

func expandDNSPolicyDns64ConfigScopeAllQueries(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyEnableInboundForwarding(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyEnableLogging(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandDNSPolicyNetworks(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedNetworkUrl, err := expandDNSPolicyNetworksNetworkUrl(original["network_url"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedNetworkUrl); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["networkUrl"] = transformedNetworkUrl
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandDNSPolicyNetworksNetworkUrl(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	if v == nil || v.(string) == "" {
		return "", nil
	} else if strings.HasPrefix(v.(string), "https://") {
		return v, nil
	}
	url, err := tpgresource.ReplaceVars(d, config, "{{ComputeBasePath}}"+v.(string))
	if err != nil {
		return "", err
	}
	return tpgresource.ConvertSelfLinkToV1(url), nil
}
