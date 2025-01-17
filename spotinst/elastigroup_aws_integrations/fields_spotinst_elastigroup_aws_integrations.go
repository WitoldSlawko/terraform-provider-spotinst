package elastigroup_aws_integrations

import (
	"github.com/spotinst/spotinst-sdk-go/service/elastigroup/providers/aws"
	"github.com/spotinst/spotinst-sdk-go/spotinst"
	"github.com/WitoldSlawko/terraform-provider-spotinst/spotinst/commons"
)

func Setup(fieldsMap map[commons.FieldName]*commons.GenericField) {
	SetupEcs(fieldsMap)
	SetupNomad(fieldsMap)
	SetupGitlab(fieldsMap)
	SetupRancher(fieldsMap)
	SetupKubernetes(fieldsMap)
	SetupMesosphere(fieldsMap)
	SetupCodeDeploy(fieldsMap)
	SetupMultaiRuntime(fieldsMap)
	SetupRoute53(fieldsMap)
	SetupDockerSwarm(fieldsMap)
	SetupElasticBeanstalk(fieldsMap)
}

func expandAWSGroupAutoScaleHeadroom(data interface{}) (*aws.AutoScaleHeadroom, error) {
	if list := data.([]interface{}); len(list) > 0 {
		headroom := &aws.AutoScaleHeadroom{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})

			if v, ok := m[string(CpuPerUnit)].(int); ok && v > 0 {
				headroom.SetCPUPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(MemoryPerUnit)].(int); ok && v > 0 {
				headroom.SetMemoryPerUnit(spotinst.Int(v))
			}

			if v, ok := m[string(NumOfUnits)].(int); ok && v > 0 {
				headroom.SetNumOfUnits(spotinst.Int(v))
			}
		}
		return headroom, nil
	}

	return nil, nil
}

func expandAWSGroupAutoScaleDown(data interface{}, isMaxScaleDownPercentageExist bool) (*aws.AutoScaleDown, error) {
	if list := data.([]interface{}); len(list) > 0 {
		autoScaleDown := &aws.AutoScaleDown{}
		if list != nil && list[0] != nil {
			m := list[0].(map[string]interface{})
			var maxScaleDownPercentage *float64 = nil

			if v, ok := m[string(EvaluationPeriods)].(int); ok && v > 0 {
				autoScaleDown.SetEvaluationPeriods(spotinst.Int(v))
			}

			if v, ok := m[string(MaxScaleDownPercentage)].(float64); ok && v > 0 {
				maxScaleDownPercentage = spotinst.Float64(v)
			}
			if isMaxScaleDownPercentageExist {
				autoScaleDown.SetMaxScaleDownPercentage(maxScaleDownPercentage)
			}
		}
		return autoScaleDown, nil
	}

	return nil, nil
}
