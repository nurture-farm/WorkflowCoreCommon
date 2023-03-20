/*
 *  Copyright 2023 NURTURE AGTECH PVT LTD
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

/*
 *  Copyright 2023 NURTURE AGTECH PVT LTD
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package activity_calls

import (
	"fmt"
	cmps "github.com/nurture-farm/Contracts/CampaignService/Gen/GoCampaignService"
	common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"github.com/nurture-farm/WorkflowCoreCommon/utils"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

func CMPSCampaign(ctx workflow.Context, flowName string, request *cmps.CampaignRequest) (*cmps.CampaignResponse, error) {
	ctx = utils.WithCMPSActivityOptions(ctx)
	cmpsFuture := workflow.ExecuteActivity(ctx, cmps.CampaignServiceServer.ExecuteCampaign, request)
	cmpsResponse := &cmps.CampaignResponse{}
	err := cmpsFuture.Get(ctx, &cmpsResponse)
	if err != nil {
		workflow.GetLogger(ctx).Error("CMPSCampaign:::FAILED::ExecuteCampaign",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("CampaignRequest", request))
		return nil, fmt.Errorf("CMPSCampaign:::FAILED::ExecuteCampaign, cause: %v, reqest: %v",
			err, request)

	}
	if cmpsResponse.Status.Status != common.RequestStatus_SUCCESS {
		workflow.GetLogger(ctx).Error("CMPSCampaign:::FAILED::ExecuteCampaign::InvalidResponse",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("CampaignRequest", request),
			zap.Any("cmpsResponse", cmpsResponse))
		return nil, fmt.Errorf("CMPSCampaign:::FAILED::ExecuteCampaign::InvalidResponse, cause: %v, reqest: %v",
			err, request)
	}
	workflow.GetLogger(ctx).Info("CMPSCampaign:::SUCCESS::ExecuteCampaign",
		zap.String("flowName", flowName), zap.Any("CampaignRequest", request),
		zap.Any("response", cmpsResponse))
	return cmpsResponse, nil
}

func CMPSUserJourneyCampaign(ctx workflow.Context, flowName string, request *cmps.UserJourneyCampaignRequest) (*cmps.UserJourneyCampaignResponse, error) {
	ctx = utils.WithCMPSActivityOptions(ctx)
	cmpsFuture := workflow.ExecuteActivity(ctx, cmps.CampaignServiceServer.ExecuteUserJourneyCampaign, request)
	cmpsResponse := &cmps.UserJourneyCampaignResponse{}
	err := cmpsFuture.Get(ctx, &cmpsResponse)
	if err != nil {
		workflow.GetLogger(ctx).Error("CMPSUserJourneyCampaign:::FAILED::ExecuteUserJourneyCampaign",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("UserJourneyCampaignRequest", request))
		return nil, fmt.Errorf("CMPSUserJourneyCampaign:::FAILED::ExecuteUserJourneyCampaign, cause: %v, reqest: %v",
			err, request)

	}
	if cmpsResponse.Status.Status != common.RequestStatus_SUCCESS {
		workflow.GetLogger(ctx).Error("CMPSUserJourneyCampaign:::FAILED::ExecuteUserJourneyCampaign::InvalidResponse",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("UserJourneyCampaignRequest", request),
			zap.Any("UserJourneycmpsResponse", cmpsResponse))
		return nil, fmt.Errorf("CMPSUserJourneyCampaign:::FAILED::ExecuteUserJourneyCampaign::InvalidResponse, cause: %v, reqest: %v",
			err, request)
	}
	workflow.GetLogger(ctx).Info("CMPSUserJourneyCampaign:::SUCCESS::ExecuteUserJourneyCampaign",
		zap.String("flowName", flowName), zap.Any("UserJourneyCampaignRequest", request),
		zap.Any("response", cmpsResponse))
	return cmpsResponse, nil
}

func CMPSUpdateCampaign(ctx workflow.Context, flowName string, request *cmps.UpdateCampaignRequest) (*cmps.UpdateCampaignResponse, error) {
	ctx = utils.WithCMPSActivityOptions(ctx)
	cmpsFuture := workflow.ExecuteActivity(ctx, cmps.CampaignServiceServer.ExecuteUpdateCampaign, request)
	cmpsResponse := &cmps.UpdateCampaignResponse{}
	err := cmpsFuture.Get(ctx, &cmpsResponse)
	if err != nil {
		workflow.GetLogger(ctx).Error("CMPSCampaign:::FAILED::ExecuteUpdateCampaign",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("CampaignRequest", request))
		return nil, fmt.Errorf("CMPSCampaign:::FAILED::ExecuteUpdateCampaign, cause: %v, reqest: %v",
			err, request)

	}
	if cmpsResponse.Status.Status != common.RequestStatus_SUCCESS {
		workflow.GetLogger(ctx).Error("CMPSCampaign:::FAILED::ExecuteUpdateCampaign::InvalidResponse",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("CampaignRequest", request),
			zap.Any("cmpsResponse", cmpsResponse))
		return nil, fmt.Errorf("CMPSCampaign:::FAILED::ExecuteUpdateCampaign::InvalidResponse, cause: %v, reqest: %v",
			err, request)
	}
	workflow.GetLogger(ctx).Info("CMPSCampaign:::SUCCESS::ExecuteUpdateCampaign",
		zap.String("flowName", flowName), zap.Any("CampaignRequest", request),
		zap.Any("response", cmpsResponse))
	return cmpsResponse, nil
}
