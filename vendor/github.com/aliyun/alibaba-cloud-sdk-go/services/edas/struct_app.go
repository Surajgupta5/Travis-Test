package edas

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// App is a nested struct in edas response
type App struct {
	Name                   string  `json:"Name" xml:"Name"`
	CsClusterId            string  `json:"CsClusterId" xml:"CsClusterId"`
	Instances              int     `json:"Instances" xml:"Instances"`
	CreateTime             int64   `json:"CreateTime" xml:"CreateTime"`
	DeployType             string  `json:"DeployType" xml:"DeployType"`
	Dockerize              bool    `json:"Dockerize" xml:"Dockerize"`
	EdasContainerVersion   string  `json:"EdasContainerVersion" xml:"EdasContainerVersion"`
	RegionId               string  `json:"RegionId" xml:"RegionId"`
	SlbPort                int     `json:"SlbPort" xml:"SlbPort"`
	UserId                 string  `json:"UserId" xml:"UserId"`
	ApplicationType        string  `json:"ApplicationType" xml:"ApplicationType"`
	TomcatVersion          string  `json:"TomcatVersion" xml:"TomcatVersion"`
	Description            string  `json:"Description" xml:"Description"`
	ClusterId              string  `json:"ClusterId" xml:"ClusterId"`
	Port                   int     `json:"Port" xml:"Port"`
	ApplicationName        string  `json:"ApplicationName" xml:"ApplicationName"`
	ExtSlbIp               string  `json:"ExtSlbIp" xml:"ExtSlbIp"`
	BuildPackageId         int64   `json:"BuildPackageId" xml:"BuildPackageId"`
	Memory                 int     `json:"Memory" xml:"Memory"`
	BuildpackId            int     `json:"BuildpackId" xml:"BuildpackId"`
	ExtSlbId               string  `json:"ExtSlbId" xml:"ExtSlbId"`
	Owner                  string  `json:"Owner" xml:"Owner"`
	AppId                  string  `json:"AppId" xml:"AppId"`
	InstanceCount          int     `json:"InstanceCount" xml:"InstanceCount"`
	HealthCheckUrl         string  `json:"HealthCheckUrl" xml:"HealthCheckUrl"`
	InstancesBeforeScaling int     `json:"InstancesBeforeScaling" xml:"InstancesBeforeScaling"`
	SlbId                  string  `json:"SlbId" xml:"SlbId"`
	ClusterType            int     `json:"ClusterType" xml:"ClusterType"`
	Cpu                    int     `json:"Cpu" xml:"Cpu"`
	Cmd                    string  `json:"Cmd" xml:"Cmd"`
	RunningInstanceCount   int     `json:"RunningInstanceCount" xml:"RunningInstanceCount"`
	SlbIp                  string  `json:"SlbIp" xml:"SlbIp"`
	CmdArgs                CmdArgs `json:"CmdArgs" xml:"CmdArgs"`
	EnvList                EnvList `json:"EnvList" xml:"EnvList"`
}
