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
	ce "github.com/nurture-farm/Contracts/CommunicationEngine/Gen/GoCommunicationEngine"
	"github.com/nurture-farm/WorkflowCoreCommon/utils"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

func CSSendCommunication(ctx workflow.Context, flowName string, event *ce.CommunicationEvent) error {
	ctx = utils.WithCEActivityOptions(ctx)
	commFuture := workflow.ExecuteActivity(ctx, ce.CommunicationEngineServer.SendCommunication, event)
	commResponse := &ce.CommunicationResponse{}
	err := commFuture.Get(ctx, &commResponse)
	if err != nil {
		workflow.GetLogger(ctx).Error("CSSendCommunication:::FAILED::CommunicationServiceCall",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("SendCommunication", event))
		return fmt.Errorf("CSSendCommunication:::FAILED::CommunicationServiceCall, cause: %v, event: %v",
			err, event)

	} else {
		workflow.GetLogger(ctx).Info("CSSendCommunication:::SUCCESS::CommunicationServiceCall",
			zap.String("flowName", flowName), zap.Any("SendCommunication", event),
			zap.Any("response", commResponse))
	}
	return nil
}

func CSSendBulkCommunication(ctx workflow.Context, flowName string, event *ce.BulkCommunicationEvent) error {
	ctx = utils.WithCEActivityOptions(ctx)
	commFuture := workflow.ExecuteActivity(ctx, ce.CommunicationEngineServer.SendBulkCommunication, event)
	commResponse := &ce.BulkCommunicationResponse{}
	err := commFuture.Get(ctx, &commResponse)
	if err != nil {
		workflow.GetLogger(ctx).Error("CSSendBulkCommunication:::FAILED::CommunicationServiceBulkCall",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("SendBulkCommunication", event))
		return fmt.Errorf("CSSendBulkCommunication:::FAILED::CommunicationServiceBulkCall, cause: %v, event: %v",
			err, event)

	} else {
		workflow.GetLogger(ctx).Info("CSSendBulkCommunication:::SUCCESS::CommunicationServiceBulkCall",
			zap.String("flowName", flowName), zap.Any("SendBulkCommunication", event),
			zap.Any("response", commResponse))
	}
	return nil
}
