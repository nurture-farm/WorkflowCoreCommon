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

func CEPSearchMessageAcknowledgements(ctx workflow.Context, flowName string, request *ce.MessageAcknowledgementRequest) (*ce.MessageAcknowledgementResponse, error) {

	ctx = utils.WithCEPActivityOptions(ctx)
	commFuture := workflow.ExecuteActivity(ctx, ce.CommunicationEnginePlatformServer.SearchMessageAcknowledgements, request)
	commResponse := &ce.MessageAcknowledgementResponse{}
	err := commFuture.Get(ctx, &commResponse)
	if err != nil {
		workflow.GetLogger(ctx).Error("CEPSearchMessageAcknowledgements:::FAILED::SearchMessageAcknowledgements",
			zap.Error(err), zap.String("flowName", flowName), zap.Any("MessageAcknowledgementRequest", request))
		return nil, fmt.Errorf("CEPSearchMessageAcknowledgements:::FAILED::SearchMessageAcknowledgements, cause: %v, request: %v", err, request)

	}
	workflow.GetLogger(ctx).Info("CEPSearchMessageAcknowledgements:::SUCCESS::SearchMessageAcknowledgements",
		zap.String("flowName", flowName), zap.Any("MessageAcknowledgementRequest", request), zap.Any("response", commResponse))
	return commResponse, nil
}
