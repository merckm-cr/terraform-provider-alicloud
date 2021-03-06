package alicloud

import (
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ram"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
)

func resourceAlicloudRamGroupPolicyAtatchment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlicloudRamGroupPolicyAttachmentCreate,
		Read:   resourceAlicloudRamGroupPolicyAttachmentRead,
		Delete: resourceAlicloudRamGroupPolicyAttachmentDelete,

		Schema: map[string]*schema.Schema{
			"group_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policy_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validatePolicyType,
			},
		},
	}
}

func resourceAlicloudRamGroupPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)

	request := ram.CreateAttachPolicyToGroupRequest()
	request.PolicyType = d.Get("policy_type").(string)
	request.PolicyName = d.Get("policy_name").(string)
	request.GroupName = d.Get("group_name").(string)

	raw, err := client.WithRamClient(func(ramClient *ram.Client) (interface{}, error) {
		return ramClient.AttachPolicyToGroup(request)
	})
	if err != nil {
		return WrapErrorf(err, DefaultErrorMsg, "alicloud_ram_group_policy_attachment", request.GetActionName(), AlibabaCloudSdkGoERROR)
	}
	addDebug(request.GetActionName(), raw)
	d.SetId(strings.Join([]string{"group", request.PolicyName, string(request.PolicyType), request.GroupName}, COLON_SEPARATED))

	return resourceAlicloudRamGroupPolicyAttachmentRead(d, meta)
}

func resourceAlicloudRamGroupPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ramService := RamService{client}
	id := strings.Join([]string{"group", d.Get("policy_name").(string), d.Get("policy_type").(string), d.Get("group_name").(string)}, COLON_SEPARATED)
	if d.Id() != id {
		d.SetId(id)
	}
	object, err := ramService.DescribeRamGroupPolicyAttachment(d.Id())
	if err != nil {
		if NotFoundError(err) {
			d.SetId("")
			return nil
		}
		return WrapError(err)
	}

	parts, err := ParseResourceId(id, 4)
	if err != nil {
		return WrapError(err)
	}
	d.Set("group_name", parts[3])
	d.Set("policy_name", object.PolicyName)
	d.Set("policy_type", object.PolicyType)
	return nil
}

func resourceAlicloudRamGroupPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*connectivity.AliyunClient)
	ramService := RamService{client}

	id := strings.Join([]string{"group", d.Get("policy_name").(string), d.Get("policy_type").(string), d.Get("group_name").(string)}, COLON_SEPARATED)
	if d.Id() != id {
		d.SetId(id)
	}
	parts, err := ParseResourceId(id, 4)
	if err != nil {
		return WrapError(err)
	}
	request := ram.CreateDetachPolicyFromGroupRequest()
	request.PolicyName = parts[1]
	request.PolicyType = parts[2]
	request.GroupName = parts[3]

	raw, err := client.WithRamClient(func(ramClient *ram.Client) (interface{}, error) {
		return ramClient.DetachPolicyFromGroup(request)
	})
	if err != nil {
		if RamEntityNotExist(err) {
			return nil
		}
		return WrapErrorf(err, DefaultErrorMsg, d.Id(), request.GetActionName(), AlibabaCloudSdkGoERROR)
	}

	addDebug(request.GetActionName(), raw)

	return WrapError(ramService.WaitForRamGroupPolicyAttachment(d.Id(), Deleted, DefaultTimeout))

}
