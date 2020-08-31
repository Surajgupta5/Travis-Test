package smartag

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

// ModifyCloudConnectNetwork invokes the smartag.ModifyCloudConnectNetwork API synchronously
// api document: https://help.aliyun.com/api/smartag/modifycloudconnectnetwork.html
func (client *Client) ModifyCloudConnectNetwork(request *ModifyCloudConnectNetworkRequest) (response *ModifyCloudConnectNetworkResponse, err error) {
	response = CreateModifyCloudConnectNetworkResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCloudConnectNetworkWithChan invokes the smartag.ModifyCloudConnectNetwork API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifycloudconnectnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCloudConnectNetworkWithChan(request *ModifyCloudConnectNetworkRequest) (<-chan *ModifyCloudConnectNetworkResponse, <-chan error) {
	responseChan := make(chan *ModifyCloudConnectNetworkResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCloudConnectNetwork(request)
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

// ModifyCloudConnectNetworkWithCallback invokes the smartag.ModifyCloudConnectNetwork API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifycloudconnectnetwork.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCloudConnectNetworkWithCallback(request *ModifyCloudConnectNetworkRequest, callback func(response *ModifyCloudConnectNetworkResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCloudConnectNetworkResponse
		var err error
		defer close(result)
		response, err = client.ModifyCloudConnectNetwork(request)
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

// ModifyCloudConnectNetworkRequest is the request struct for api ModifyCloudConnectNetwork
type ModifyCloudConnectNetworkRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CcnId                string           `position:"Query" name:"CcnId"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	CidrBlock            string           `position:"Query" name:"CidrBlock"`
	InterworkingStatus   string           `position:"Query" name:"InterworkingStatus"`
}

// ModifyCloudConnectNetworkResponse is the response struct for api ModifyCloudConnectNetwork
type ModifyCloudConnectNetworkResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCloudConnectNetworkRequest creates a request to invoke ModifyCloudConnectNetwork API
func CreateModifyCloudConnectNetworkRequest() (request *ModifyCloudConnectNetworkRequest) {
	request = &ModifyCloudConnectNetworkRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifyCloudConnectNetwork", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCloudConnectNetworkResponse creates a response to parse from ModifyCloudConnectNetwork response
func CreateModifyCloudConnectNetworkResponse() (response *ModifyCloudConnectNetworkResponse) {
	response = &ModifyCloudConnectNetworkResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
