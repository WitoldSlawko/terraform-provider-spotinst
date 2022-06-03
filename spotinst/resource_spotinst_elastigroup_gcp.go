package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/gcp"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/commons"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_disk"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_gpu"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_instance_types"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_integrations"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_launch_configuration"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_network_interface"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_scaling_policies"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_scheduled_task"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_gcp_strategy"
)

func resourceSpotinstElastigroupGCP() *schema.Resource {
	setupElastigroupGCPResource()

	return &schema.Resource{
		CreateContext: resourceSpotinstElastigroupGCPCreate,
		ReadContext:   resourceSpotinstElastigroupGCPRead,
		UpdateContext: resourceSpotinstElastigroupGCPUpdate,
		DeleteContext: resourceSpotinstElastigroupGCPDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.ElastigroupGCPResource.GetSchemaMap(),
	}
}

// setupElastigroupGCPResource calls the setup function for each of the children blocks.
func setupElastigroupGCPResource() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	elastigroup_gcp.Setup(fieldsMap)
	elastigroup_gcp_disk.Setup(fieldsMap)
	elastigroup_gcp_gpu.Setup(fieldsMap)
	elastigroup_gcp_instance_types.Setup(fieldsMap)
	elastigroup_gcp_integrations.Setup(fieldsMap)
	elastigroup_gcp_launch_configuration.Setup(fieldsMap)
	elastigroup_gcp_network_interface.Setup(fieldsMap)
	elastigroup_gcp_scaling_policies.Setup(fieldsMap)
	elastigroup_gcp_scheduled_task.Setup(fieldsMap)
	elastigroup_gcp_strategy.Setup(fieldsMap)

	commons.ElastigroupGCPResource = commons.NewElastigroupGCPResource(fieldsMap)
}

// resourceSpotinstElastigroupGCPCreate begins the creation request and
// creates an object representing the newly created group or returns an error.
func resourceSpotinstElastigroupGCPCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupGCPResource.GetName())

	elastigroup, err := commons.ElastigroupGCPResource.OnCreate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	groupId, err := createGCPGroup(elastigroup, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(groupId))
	log.Printf("===> Elastigroup created successfully: %s <===", resourceData.Id())
	return resourceSpotinstElastigroupGCPRead(ctx, resourceData, meta)
}

// createGCPGroup makes the create request to the spotinst API and returns
// the group ID of created group or an error if the request fails. It will retry
// the request (default 1 min) when encountering a retryable error.
func createGCPGroup(elastigroup *gcp.Group, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(elastigroup); err != nil {
		return nil, err
	} else {
		log.Printf("===> Group create configuration: %s", json)
	}
	var resp *gcp.CreateGroupOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &gcp.CreateGroupInput{Group: elastigroup}
		r, err := spotinstClient.elastigroup.CloudProviderGCP().Create(context.Background(), input)
		if err != nil {
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("[ERROR] failed to create group: %s", err)
	}
	return resp.Group.ID, nil
}

// resourceSpotinstElastigroupGCPRead creates an object representing an existing elastigroup
// by making a get request using the Spotinst API or returns an error.
func resourceSpotinstElastigroupGCPRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	groupId := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead),
		commons.ElastigroupGCPResource.GetName(), groupId)

	input := &gcp.ReadGroupInput{GroupID: spotinst.String(groupId)}
	resp, err := meta.(*Client).elastigroup.CloudProviderGCP().Read(context.Background(), input)
	if err != nil {
		// if the group was not found, return nil to show the group doesn't exist
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// report any other error
		return diag.Errorf("failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	groupResponse := resp.Group
	if groupResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ElastigroupGCPResource.OnRead(groupResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Elastigroup read successfully: %s <===", groupId)
	if json, err := commons.ToJson(groupResponse); err != nil {
		return diag.FromErr(err)
	} else {
		log.Printf("===> Group read configuration: %s", json)
	}
	return nil
}

// resourceSpotinstElastigroupGCPUpdate updates an existing elastigroup
// and creates an object representing the updated group or returns an error.
func resourceSpotinstElastigroupGCPUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	groupId := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupGCPResource.GetName(), groupId)

	shouldUpdate, elastigroup, err := commons.ElastigroupGCPResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	if shouldUpdate {
		elastigroup.SetID(spotinst.String(groupId))
		if err := updateGCPGroup(elastigroup, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Elastigroup updated successfully: %s <===", groupId)
	return resourceSpotinstElastigroupGCPRead(ctx, resourceData, meta)
}

// updateGCPGroup sends the update request to the Spotinst API and returns an error if the request fails.
func updateGCPGroup(elastigroup *gcp.Group, resourceData *schema.ResourceData, meta interface{}) error {
	var input = &gcp.UpdateGroupInput{Group: elastigroup}
	groupId := resourceData.Id()

	if json, err := commons.ToJson(elastigroup); err != nil {
		return err
	} else {
		log.Printf("===> Group update configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderGCP().Update(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] Failed to update group [%v]: %v", groupId, err)
	}

	return nil
}

// resourceSpotinstElastigroupGCPDelete deletes a specific elastigroup or returns an error.
func resourceSpotinstElastigroupGCPDelete(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	groupId := resourceData.Id()
	log.Printf(string(commons.ResourceOnDelete),
		commons.ElastigroupGCPResource.GetName(), groupId)

	if err := deleteGCPGroup(resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Elastigroup deleted successfully: %s <===", resourceData.Id())
	resourceData.SetId("")
	return nil
}

// deleteGCPGroup sends the delete request to the Spotinst API or an error if the request fails.
func deleteGCPGroup(resourceData *schema.ResourceData, meta interface{}) error {
	groupId := resourceData.Id()
	input := &gcp.DeleteGroupInput{GroupID: spotinst.String(groupId)}

	if json, err := commons.ToJson(input); err != nil {
		return err
	} else {
		log.Printf("===> Group delete configuration: %s", json)
	}

	if _, err := meta.(*Client).elastigroup.CloudProviderGCP().Delete(context.Background(), input); err != nil {
		return fmt.Errorf("[ERROR] onDelete() -> Failed to delete group: %s", err)
	}
	return nil
}
