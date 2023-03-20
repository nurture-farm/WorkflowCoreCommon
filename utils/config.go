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
	"fmt"
	"os"
	"path"
	"time"

	Common "github.com/nurture-farm/Contracts/Common/Gen/GoCommon"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/client"
	"go.uber.org/zap"
)

var SMSTemplatesConfig map[string]string
var TemplatesConfig map[string]map[string]string
var ServiceSpecSpecificTemplatesConfig map[string]map[string]map[string]string
var WorkflowConfig map[string]string
var GRPCClientConfig map[string]string
var DefaultTimezone, _ = time.LoadLocation("Asia/Kolkata")

var (
	TemporalHostPort  = "localhost:7233"
	TemporalNamespace = "default"
	SupportedChannels = []Common.CommunicationChannel{Common.CommunicationChannel_SMS,
		Common.CommunicationChannel_WHATSAPP}
)

var _ = getConfig()

// TODO: Check these flags +build !testing
func getConfig() error {
	basePath := os.Getenv("CONFIG_DIR")
	if len(basePath) == 0 {
		LOG.Warn("No config is read since CONFIG_DIR environment variable is not set")
		return nil
	}
	LOG.Info("Config read:", zap.String("basePath", basePath))
	configPath := basePath
	if len(os.Getenv("ENV")) > 0 {
		configPath = path.Join(basePath, os.Getenv("ENV"))
		LOG.Info("Overriding config path, since ENV variable found: ",
			zap.String("ENV", os.Getenv("ENV")), zap.String("configPath", configPath))
	}
	viper.AddConfigPath(configPath)
	viper.SetConfigFile(configPath + "/config.json")
	viper.SetConfigType("json")
	//viper.SetEnvPrefix("CONFIG")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		LOG.Panic("VIPER config read error", zap.Error(err))
	}
	configureCommunicationConfigs()
	WorkflowConfig = viper.GetStringMapString("workflow_config")
	GRPCClientConfig = viper.GetStringMapString("grpc_config")
	TemporalHostPort = viper.GetString("workflow_config.temporal_host_port")
	TemporalNamespace = viper.GetString("workflow_config.namespace")
	WorkflowConfig["temporal_host_port"] = TemporalHostPort
	WorkflowConfig["namespace"] = TemporalNamespace
	WorkflowConfig["worker_name"] = viper.GetString("workflow_config.worker_name")
	WorkflowConfig["ce_task_queue"] = viper.GetString("workflow_config.ce_task_queue")
	LOG.Info("CONFIG_READ:::", zap.Any("SMSTemplatesConfig", SMSTemplatesConfig),
		zap.Any("TemplatesConfig", TemplatesConfig),
		zap.Any("ServiceSpecSpecificTemplatesConfig", ServiceSpecSpecificTemplatesConfig),
		zap.Any("GRPCClientConfig", GRPCClientConfig), zap.Any("WorkflowConfig", WorkflowConfig))
	return nil
}

func configureCommunicationConfigs() {
	// Getting default templates
	TemplatesConfig = make(map[string]map[string]string)
	for _, channel := range SupportedChannels {
		TemplatesConfig[channel.String()] = viper.GetStringMapString(fmt.Sprintf("templates.%s", channel.String()))
		if channel == Common.CommunicationChannel_SMS {
			SMSTemplatesConfig = TemplatesConfig[channel.String()]
		}
	}

	// Fetching service spec specific templates
	templates := viper.GetStringMap("service_spec_id_specific_templates")
	ServiceSpecSpecificTemplatesConfig = make(map[string]map[string]map[string]string)
	for k, _ := range templates {
		ServiceSpecSpecificTemplatesConfig[k] = make(map[string]map[string]string)
		for _, channel := range SupportedChannels {
			ServiceSpecSpecificTemplatesConfig[k][channel.String()] =
				viper.GetStringMapString(fmt.Sprintf("service_spec_id_specific_templates.%s.%s", k, channel.String()))
		}
	}
}

var WorkflowClient client.Client = getClient()

func getClient() client.Client {
	return GetWFClient(TemporalNamespace, TemporalHostPort)
}

func GetWFClient(namespace, hostPort string) client.Client {
	LOG.Info("Creating workflow client", zap.String("namespace", namespace),
		zap.String("hostPort", hostPort))
	c, err := client.NewClient(client.Options{
		Namespace: namespace,
		HostPort:  hostPort,
	})
	if err != nil {
		LOG.Error("Unable to create client", zap.Error(err))
	}
	return c
}
