package emr

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

// EmrMainVersion is a nested struct in emr response
type EmrMainVersion struct {
	EcmVersion          bool                                        `json:"EcmVersion" xml:"EcmVersion"`
	RegionId            string                                      `json:"RegionId" xml:"RegionId"`
	EmrVersion          string                                      `json:"EmrVersion" xml:"EmrVersion"`
	Display             bool                                        `json:"Display" xml:"Display"`
	ImageId             string                                      `json:"ImageId" xml:"ImageId"`
	MainVersionName     string                                      `json:"MainVersionName" xml:"MainVersionName"`
	StackName           string                                      `json:"StackName" xml:"StackName"`
	StackVersion        string                                      `json:"StackVersion" xml:"StackVersion"`
	WhiteUserList       WhiteUserListInDescribeEmrMainVersion       `json:"WhiteUserList" xml:"WhiteUserList"`
	ClusterTypeInfoList ClusterTypeInfoListInDescribeEmrMainVersion `json:"ClusterTypeInfoList" xml:"ClusterTypeInfoList"`
}
