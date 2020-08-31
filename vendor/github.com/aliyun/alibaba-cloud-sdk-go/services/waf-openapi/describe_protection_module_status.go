package waf_openapi

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeProtectionModuleStatus invokes the waf_openapi.DescribeProtectionModuleStatus API synchronously
// api document: https://help.aliyun.com/api/waf-openapi/describeprotectionmodulestatus.html
func (client *Client) DescribeProtectionModuleStatus(request *DescribeProtectionModuleStatusRequest) (response *DescribeProtectionModuleStatusResponse, err error) {
	response = CreateDescribeProtectionModuleStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeProtectionModuleStatusWithChan invokes the waf_openapi.DescribeProtectionModuleStatus API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/describeprotectionmodulestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProtectionModuleStatusWithChan(request *DescribeProtectionModuleStatusRequest) (<-chan *DescribeProtectionModuleStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeProtectionModuleStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeProtectionModuleStatus(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeProtectionModuleStatusWithCallback invokes the waf_openapi.DescribeProtectionModuleStatus API asynchronously
// api document: https://help.aliyun.com/api/waf-openapi/describeprotectionmodulestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeProtectionModuleStatusWithCallback(request *DescribeProtectionModuleStatusRequest, callback func(response *DescribeProtectionModuleStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeProtectionModuleStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeProtectionModuleStatus(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeProtectionModuleStatusRequest is the request struct for api DescribeProtectionModuleStatus
type DescribeProtectionModuleStatusRequest struct {
	*requests.RpcRequest
	DefenseType string `position:"Query" name:"DefenseType"`
	InstanceId  string `position:"Query" name:"InstanceId"`
	SourceIp    string `position:"Query" name:"SourceIp"`
	Domain      string `position:"Query" name:"Domain"`
	Lang        string `position:"Query" name:"Lang"`
}

// DescribeProtectionModuleStatusResponse is the response struct for api DescribeProtectionModuleStatus
type DescribeProtectionModuleStatusResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ModuleStatus int    `json:"ModuleStatus" xml:"ModuleStatus"`
}

// CreateDescribeProtectionModuleStatusRequest creates a request to invoke DescribeProtectionModuleStatus API
func CreateDescribeProtectionModuleStatusRequest() (request *DescribeProtectionModuleStatusRequest) {
	request = &DescribeProtectionModuleStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DescribeProtectionModuleStatus", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeProtectionModuleStatusResponse creates a response to parse from DescribeProtectionModuleStatus response
func CreateDescribeProtectionModuleStatusResponse() (response *DescribeProtectionModuleStatusResponse) {
	response = &DescribeProtectionModuleStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
