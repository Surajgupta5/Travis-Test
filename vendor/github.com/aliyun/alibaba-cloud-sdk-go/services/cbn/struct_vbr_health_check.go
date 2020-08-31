package cbn

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

// VbrHealthCheck is a nested struct in cbn response
type VbrHealthCheck struct {
	CenId               string `json:"CenId" xml:"CenId"`
	VbrInstanceId       string `json:"VbrInstanceId" xml:"VbrInstanceId"`
	LinkStatus          string `json:"LinkStatus" xml:"LinkStatus"`
	PacketLoss          int64  `json:"PacketLoss" xml:"PacketLoss"`
	HealthCheckSourceIp string `json:"HealthCheckSourceIp" xml:"HealthCheckSourceIp"`
	HealthCheckTargetIp string `json:"HealthCheckTargetIp" xml:"HealthCheckTargetIp"`
	Delay               int64  `json:"Delay" xml:"Delay"`
	HealthCheckInterval int    `json:"HealthCheckInterval" xml:"HealthCheckInterval"`
	HealthyThreshold    int    `json:"HealthyThreshold" xml:"HealthyThreshold"`
}
