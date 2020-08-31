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

// ModifySagExpressConnectInterface invokes the smartag.ModifySagExpressConnectInterface API synchronously
// api document: https://help.aliyun.com/api/smartag/modifysagexpressconnectinterface.html
func (client *Client) ModifySagExpressConnectInterface(request *ModifySagExpressConnectInterfaceRequest) (response *ModifySagExpressConnectInterfaceResponse, err error) {
	response = CreateModifySagExpressConnectInterfaceResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySagExpressConnectInterfaceWithChan invokes the smartag.ModifySagExpressConnectInterface API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifysagexpressconnectinterface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySagExpressConnectInterfaceWithChan(request *ModifySagExpressConnectInterfaceRequest) (<-chan *ModifySagExpressConnectInterfaceResponse, <-chan error) {
	responseChan := make(chan *ModifySagExpressConnectInterfaceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySagExpressConnectInterface(request)
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

// ModifySagExpressConnectInterfaceWithCallback invokes the smartag.ModifySagExpressConnectInterface API asynchronously
// api document: https://help.aliyun.com/api/smartag/modifysagexpressconnectinterface.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySagExpressConnectInterfaceWithCallback(request *ModifySagExpressConnectInterfaceRequest, callback func(response *ModifySagExpressConnectInterfaceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySagExpressConnectInterfaceResponse
		var err error
		defer close(result)
		response, err = client.ModifySagExpressConnectInterface(request)
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

// ModifySagExpressConnectInterfaceRequest is the request struct for api ModifySagExpressConnectInterface
type ModifySagExpressConnectInterfaceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Vlan                 string           `position:"Query" name:"Vlan"`
	Mask                 string           `position:"Query" name:"Mask"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	IP                   string           `position:"Query" name:"IP"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SmartAGId            string           `position:"Query" name:"SmartAGId"`
	SmartAGSn            string           `position:"Query" name:"SmartAGSn"`
	PortName             string           `position:"Query" name:"PortName"`
}

// ModifySagExpressConnectInterfaceResponse is the response struct for api ModifySagExpressConnectInterface
type ModifySagExpressConnectInterfaceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifySagExpressConnectInterfaceRequest creates a request to invoke ModifySagExpressConnectInterface API
func CreateModifySagExpressConnectInterfaceRequest() (request *ModifySagExpressConnectInterfaceRequest) {
	request = &ModifySagExpressConnectInterfaceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "ModifySagExpressConnectInterface", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifySagExpressConnectInterfaceResponse creates a response to parse from ModifySagExpressConnectInterface response
func CreateModifySagExpressConnectInterfaceResponse() (response *ModifySagExpressConnectInterfaceResponse) {
	response = &ModifySagExpressConnectInterfaceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
