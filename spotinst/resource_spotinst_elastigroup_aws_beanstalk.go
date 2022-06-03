package spotinst

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/tree/main/spotinst/commons"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_aws_beanstalk"
	"https://github.com/WitoldSlawko/terraform-provider-spotinst/elastigroup_aws_beanstalk_scheduled_task"
)

func resourceSpotinstElastigroupAWSBeanstalk() *schema.Resource {
	setupElastigroupAWSBeanstalk()
	return &schema.Resource{
		CreateContext: resourceSpotinstAWSBeanstalkGroupCreate,
		ReadContext:   resourceSpotinstAWSBeanstalkGroupRead,
		UpdateContext: resourceSpotinstAWSBeanstalkGroupUpdate,
		DeleteContext: resourceSpotinstAWSBeanstalkGroupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: commons.ElastigroupAWSBeanstalkResource.GetSchemaMap(),
	}
}

func setupElastigroupAWSBeanstalk() {
	fieldsMap := make(map[commons.FieldName]*commons.GenericField)

	elastigroup_aws_beanstalk.Setup(fieldsMap)
	elastigroup_aws_beanstalk_scheduled_task.Setup(fieldsMap)

	commons.ElastigroupAWSBeanstalkResource = commons.NewElastigroupAWSBeanstalkResource(fieldsMap)
}

func importBeanstalkGroup(resourceData *schema.ResourceData, meta interface{}) (*aws.Group, error) {
	var input *aws.ImportBeanstalkInput

	if environmentId, ok := resourceData.GetOk("beanstalk_environment_id"); ok {
		input = &aws.ImportBeanstalkInput{
			EnvironmentId: spotinst.String(environmentId.(string)),
			Region:        spotinst.String(resourceData.Get("region").(string))}

	} else if environmentName, ok := resourceData.GetOk("beanstalk_environment_name"); ok {
		input = &aws.ImportBeanstalkInput{
			EnvironmentName: spotinst.String(environmentName.(string)),
			Region:          spotinst.String(resourceData.Get("region").(string))}
	}

	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().ImportBeanstalkEnv(context.Background(), input)

	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil, err
				}
			}
		}
		// Some other error, report it.
		return nil, fmt.Errorf("BEANSTALK:IMPORT failed to read group: %s", err)
	}

	return resp.Group, err
}

func toggleMaintenanceMode(resourceData *schema.ResourceData, meta interface{}, op string) diag.Diagnostics {
	id := resourceData.Id()

	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &aws.BeanstalkMaintenanceInput{GroupID: spotinst.String(id)}
		if status, err := meta.(*Client).elastigroup.CloudProviderAWS().GetBeanstalkMaintenanceStatus(context.Background(), input); err == nil {
			if op == "START" {
				if *status == "AWAIT_USER_UPDATE" {
					err = fmt.Errorf("===> Unable to start maintenance, already in maintenance mode")
					return resource.NonRetryableError(err)
				} else if *status == "ACTIVE" {
					_, err := meta.(*Client).elastigroup.CloudProviderAWS().StartBeanstalkMaintenance(context.Background(), input)
					if err != nil {
						return resource.NonRetryableError(err)
					}
					log.Printf("===> Sending request to begin Beanstalk Maintenance Mode <===")
				} else {
					err = fmt.Errorf("===> Unable to start maintenance, group status is: %s <===", *status)
					return resource.RetryableError(err)
				}
				return nil
			} else if op == "END" {
				if *status == "ACTIVE" {
					err = fmt.Errorf("===> Unable to end maintenance, your beanstalk elastigroup is already active")
					return resource.NonRetryableError(err)
				} else if *status == "AWAIT_USER_UPDATE" {
					_, err := meta.(*Client).elastigroup.CloudProviderAWS().FinishBeanstalkMaintenance(context.Background(), input)
					if err != nil {
						return resource.NonRetryableError(err)
					}
					log.Printf("===> Sending request to end Beanstalk Maintenance Mode <===")
				} else {
					err = fmt.Errorf("===> Unable to end Maintenance state, group status is: %s <===", *status)
					return resource.RetryableError(err)
				}
				return nil
			} else if op == "STATUS" {
				log.Printf("===> Beanstalk Maintenance Status: %s <===", *status)
				return nil
			}
		}
		return nil
	})

	if err != nil {
		return diag.Errorf("BEANSTALK:MaintenanceMode failed to resolve Maintenance Mode %s", err)
	}
	return nil
}

func resourceSpotinstAWSBeanstalkGroupCreate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf(string(commons.ResourceOnCreate),
		commons.ElastigroupAWSBeanstalkResource.GetName())

	beanstalkGroup, err := importBeanstalkGroup(resourceData, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	if beanstalkGroup == nil {
		return diag.Errorf("[ERROR] Failed to import group. Does the Beanstalk environment exist?")
	}

	tempGroup, err := commons.ElastigroupAWSBeanstalkResource.OnCreate(beanstalkGroup, resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	groupId, err := createBeanstalkGroup(tempGroup, meta.(*Client))
	if err != nil {
		return diag.FromErr(err)
	}

	resourceData.SetId(spotinst.StringValue(groupId))
	log.Printf("===> AWSBeanstalkGroup created successfully: %s <===", resourceData.Id())
	return resourceSpotinstAWSBeanstalkGroupRead(ctx, resourceData, meta)
}

func createBeanstalkGroup(beanstalkGroup *aws.Group, spotinstClient *Client) (*string, error) {
	if json, err := commons.ToJson(beanstalkGroup); err != nil {
		return nil, err
	} else {
		log.Printf("===> Group create configuration: %s", json)
	}

	var resp *aws.CreateGroupOutput = nil
	err := resource.RetryContext(context.Background(), time.Minute, func() *resource.RetryError {
		input := &aws.CreateGroupInput{Group: beanstalkGroup}
		r, err := spotinstClient.elastigroup.CloudProviderAWS().Create(context.Background(), input)
		if err != nil {
			// Checks whether we should retry the group creation.
			if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
				for _, err := range errs {
					if err.Code == "InvalidParameterValue" &&
						strings.Contains(err.Message, "Invalid IAM Instance Profile") {
						return resource.RetryableError(err)
					}
				}
			}

			// If there's some other error, report it.
			return resource.NonRetryableError(err)
		}
		resp = r
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("BEANSTALK:Create failed to create group: %s", err)
	}
	return resp.Group.ID, nil
}

func resourceSpotinstAWSBeanstalkGroupRead(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnRead), commons.ElastigroupAWSBeanstalkResource.GetName(), id)

	input := &aws.ReadGroupInput{GroupID: spotinst.String(id)}
	resp, err := meta.(*Client).elastigroup.CloudProviderAWS().Read(context.Background(), input)
	if err != nil {
		// If the group was not found, return nil so that we can show
		// that the group is gone.
		if errs, ok := err.(client.Errors); ok && len(errs) > 0 {
			for _, err := range errs {
				if err.Code == ErrCodeGroupNotFound {
					resourceData.SetId("")
					return nil
				}
			}
		}

		// Some other error, report it.
		return diag.Errorf("BEANSTALK:READ failed to read group: %s", err)
	}

	// If nothing was found, then return no state.
	groupResponse := resp.Group
	if groupResponse == nil {
		resourceData.SetId("")
		return nil
	}

	if err := commons.ElastigroupAWSBeanstalkResource.OnRead(groupResponse, resourceData, meta); err != nil {
		return diag.FromErr(err)
	}

	log.Printf("===> Elastigroup read successfully: %s <===", id)
	return nil
}

func resourceSpotinstAWSBeanstalkGroupUpdate(ctx context.Context, resourceData *schema.ResourceData, meta interface{}) diag.Diagnostics {
	id := resourceData.Id()
	log.Printf(string(commons.ResourceOnUpdate),
		commons.ElastigroupAWSBeanstalkResource.GetName(), id)

	shouldUpdate, elastigroupBeanstalk, err := commons.ElastigroupAWSBeanstalkResource.OnUpdate(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	maint, err := commons.ElastigroupAWSBeanstalkResource.MaintenanceState(resourceData, meta)
	if err != nil {
		return diag.FromErr(err)
	}

	maintErr := toggleMaintenanceMode(resourceData, meta, maint)
	if maintErr != nil {
		return maintErr
	}
	if shouldUpdate {
		elastigroupBeanstalk.SetId(spotinst.String(id))
		if err := updateGroup(elastigroupBeanstalk, resourceData, meta); err != nil {
			return diag.FromErr(err)
		}
	}

	log.Printf("===> Beanstalk Elastigroup updated successfully: %s <===", id)
	return resourceSpotinstAWSBeanstalkGroupRead(ctx, resourceData, meta)
}

func resourceSpotinstAWSBeanstalkGroupDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("[INFO] Deleting group: %s", d.Id())
	input := &aws.DeleteGroupInput{GroupID: spotinst.String(d.Id())}

	if _, err := meta.(*Client).elastigroup.CloudProviderAWS().Delete(context.Background(), input); err != nil {
		return diag.Errorf("failed to delete group: %s", err)
	}
	d.SetId("")
	return nil
}
