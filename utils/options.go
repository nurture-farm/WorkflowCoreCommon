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

package utils

import (
	"time"

	"github.com/spf13/viper"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"go.uber.org/zap"
)

var LOG = getLogger()

func getLogger() *zap.Logger {

	res, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return res
}

const (
	activityTimeoutScheduleToStartKey  = "ACTIVITY_TIMEOUT_SCHEDULE_TO_START"
	activityTimeoutScheduleToCloseKey  = "ACTIVITY_TIMEOUT_SCHEDULE_TO_CLOSE"
	activityTimeoutStartToCloseKey     = "ACTIVITY_TIMEOUT_START_TO_CLOSE"
	activityHeartbeatTimeoutKey        = "ACTIVITY_HEARTBEAT_TIMEOUT"
	activityWaitForCancellationKey     = "ACTIVITY_WAIT_FOR_CANCELLATION"
	activityRetryInitialIntervalKey    = "ACTIVITY_RETRY_INITIAL_INTERVAL"
	activityRetryBackoffCoefficientKey = "ACTIVITY_RETRY_BACKOFF_COEFFICIENT"
	activityRetryMaxAttemptsKey        = "ACTIVITY_RETRY_MAX_ATTEMPTS"
)

var (
	activityTimeoutScheduleToStart        = 360 * time.Second
	activityTimeoutScheduleToClose        = 360 * time.Second
	activityTimeoutStartToClose           = 360 * time.Second
	activityHeartbeatTimeout              = 0 * time.Second
	activityWaitForCancellation           = false
	activityRetryInitialInterval          = 5 * time.Second
	activityRetryBackoffCoefficient       = 2.0
	activityRetryMaxAttempts        int32 = 6
)

var _ = readConfig()

func readConfig() error {
	viper.AutomaticEnv()
	viper.SetDefault(activityTimeoutScheduleToCloseKey, activityTimeoutScheduleToClose)
	viper.SetDefault(activityTimeoutStartToCloseKey, activityTimeoutStartToClose)
	viper.SetDefault(activityTimeoutScheduleToStartKey, activityTimeoutScheduleToStart)
	viper.SetDefault(activityHeartbeatTimeoutKey, activityHeartbeatTimeout)
	viper.SetDefault(activityWaitForCancellationKey, activityWaitForCancellation)
	viper.SetDefault(activityRetryInitialIntervalKey, activityRetryInitialInterval)
	viper.SetDefault(activityRetryBackoffCoefficientKey, activityRetryBackoffCoefficient)
	viper.SetDefault(activityRetryMaxAttemptsKey, activityRetryMaxAttempts)

	activityTimeoutScheduleToStart = viper.GetDuration(activityTimeoutScheduleToStartKey)
	activityTimeoutScheduleToClose = viper.GetDuration(activityTimeoutScheduleToCloseKey)
	activityTimeoutStartToClose = viper.GetDuration(activityTimeoutStartToCloseKey)
	activityHeartbeatTimeout = viper.GetDuration(activityHeartbeatTimeoutKey)
	activityWaitForCancellation = viper.GetBool(activityWaitForCancellationKey)
	activityRetryInitialInterval = viper.GetDuration(activityRetryInitialIntervalKey)
	activityRetryBackoffCoefficient = viper.GetFloat64(activityRetryBackoffCoefficientKey)
	activityRetryMaxAttempts = viper.GetInt32(activityRetryMaxAttemptsKey)

	return nil
}

func GetActivityOptions() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		TaskQueue:              "",
		ScheduleToCloseTimeout: activityTimeoutScheduleToClose,
		ScheduleToStartTimeout: activityTimeoutScheduleToStart,
		StartToCloseTimeout:    activityTimeoutStartToClose,
		HeartbeatTimeout:       activityHeartbeatTimeout,
		WaitForCancellation:    activityWaitForCancellation,
		ActivityID:             "",
		RetryPolicy:            nil,
	}
}

func GetProdActivityOptions() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		ScheduleToCloseTimeout: activityTimeoutScheduleToClose,
		ScheduleToStartTimeout: activityTimeoutScheduleToStart,
		StartToCloseTimeout:    activityTimeoutStartToClose,
		HeartbeatTimeout:       activityHeartbeatTimeout,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    activityRetryInitialInterval,
			BackoffCoefficient: activityRetryBackoffCoefficient,

			MaximumAttempts: activityRetryMaxAttempts,
		},
		WaitForCancellation: false,
	}
}

func GetProdWorkflowOptions(id string, taskQueue string) client.StartWorkflowOptions {
	return client.StartWorkflowOptions{
		ID:                    id,
		TaskQueue:             taskQueue,
		WorkflowTaskTimeout:   10 * time.Minute,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}
}

func GetWorkflowOptionsWithWorkflowRunTimeout(id string, taskQueue string) client.StartWorkflowOptions {
	return client.StartWorkflowOptions{
		ID:                    id,
		TaskQueue:             taskQueue,
		WorkflowTaskTimeout:   5 * time.Minute,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
		WorkflowRunTimeout:    5 * time.Minute,
	}
}

func WithChildWorkflowOptions(ctx workflow.Context, id string) workflow.Context {
	options := workflow.ChildWorkflowOptions{
		WorkflowID: id,
		//TaskQueue:  WorkflowConfig["worker_name"],
	}

	childCtx := workflow.WithChildOptions(ctx, options)
	return childCtx
}

func WithChildWorkflowOptionsTaskQueue(ctx workflow.Context, taskQueue string, id string) workflow.Context {
	options := workflow.ChildWorkflowOptions{
		WorkflowID: id,
		TaskQueue:  taskQueue,
	}

	childCtx := workflow.WithChildOptions(ctx, options)
	return childCtx
}

func WithChildWorkflowOptionsDuplicates(ctx workflow.Context, id string) workflow.Context {
	options := workflow.ChildWorkflowOptions{
		WorkflowID: id,
		//TaskQueue:  WorkflowConfig["worker_name"],
		WorkflowTaskTimeout:   1 * time.Minute,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}

	childCtx := workflow.WithChildOptions(ctx, options)
	return childCtx
}

func WithIOTChildWorkflowOptionsDuplicates(ctx workflow.Context, id string) workflow.Context {
	options := workflow.ChildWorkflowOptions{
		WorkflowID: id,
		//TaskQueue:  WorkflowConfig["worker_name"],
		WorkflowTaskTimeout:   20 * time.Minute,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
		//ParentClosePolicy: enums.PARENT_CLOSE_POLICY_ABANDON,
		WorkflowExecutionTimeout: 20 * time.Hour,
	}

	childCtx := workflow.WithChildOptions(ctx, options)
	return childCtx
}

func WithShortLivedChildWorkflowOptionsDuplicates(ctx workflow.Context, id string) workflow.Context {
	options := workflow.ChildWorkflowOptions{
		WorkflowID: id,
		//TaskQueue:  WorkflowConfig["worker_name"],
		WorkflowTaskTimeout:      1 * time.Minute,
		WorkflowExecutionTimeout: 6 * time.Minute,
		WorkflowIDReusePolicy:    enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}

	childCtx := workflow.WithChildOptions(ctx, options)
	return childCtx
}

func WithChildWorkflowOptionsTaskQueueDuplicates(ctx workflow.Context, taskQueue string, id string) workflow.Context {
	options := workflow.ChildWorkflowOptions{
		WorkflowID:            id,
		TaskQueue:             taskQueue,
		WorkflowTaskTimeout:   1 * time.Minute,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE,
	}

	childCtx := workflow.WithChildOptions(ctx, options)
	return childCtx
}

func WithCEActivityOptions(ctx workflow.Context) workflow.Context {
	options := workflow.ActivityOptions{
		TaskQueue:              WorkflowConfig["ce_task_queue"],
		ScheduleToCloseTimeout: activityTimeoutScheduleToClose,
		ScheduleToStartTimeout: activityTimeoutScheduleToStart,
		StartToCloseTimeout:    activityTimeoutStartToClose,
		HeartbeatTimeout:       activityHeartbeatTimeout,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    activityRetryInitialInterval,
			BackoffCoefficient: activityRetryBackoffCoefficient,

			MaximumAttempts: activityRetryMaxAttempts,
		},
		WaitForCancellation: false,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	return ctx
}

func WithCEPActivityOptions(ctx workflow.Context) workflow.Context {
	options := workflow.ActivityOptions{
		TaskQueue:              WorkflowConfig["cep_task_queue"],
		ScheduleToCloseTimeout: activityTimeoutScheduleToClose,
		ScheduleToStartTimeout: activityTimeoutScheduleToStart,
		StartToCloseTimeout:    activityTimeoutStartToClose,
		HeartbeatTimeout:       activityHeartbeatTimeout,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    activityRetryInitialInterval,
			BackoffCoefficient: activityRetryBackoffCoefficient,
			MaximumAttempts:    activityRetryMaxAttempts,
		},
		WaitForCancellation: false,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	return ctx
}

func WithCMPSActivityOptions(ctx workflow.Context) workflow.Context {
	options := workflow.ActivityOptions{
		TaskQueue:              WorkflowConfig["cmps_task_queue"],
		ScheduleToCloseTimeout: 150 * time.Minute,
		ScheduleToStartTimeout: 6 * time.Minute,
		StartToCloseTimeout:    60 * time.Minute,
		HeartbeatTimeout:       0,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    10 * time.Second,
			BackoffCoefficient: 2,
			MaximumInterval:    6 * time.Minute,
			MaximumAttempts:    2,
		},
		WaitForCancellation: false,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	return ctx
}
