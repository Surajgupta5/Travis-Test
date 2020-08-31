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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListFlowClusterHost invokes the emr.ListFlowClusterHost API synchronously
// api document: https://help.aliyun.com/api/emr/listflowclusterhost.html
func (client *Client) ListFlowClusterHost(request *ListFlowClusterHostRequest) (response *ListFlowClusterHostResponse, err error) {
	response = CreateListFlowClusterHostResponse()
	err = client.DoAction(request, response)
	return
}

// ListFlowClusterHostWithChan invokes the emr.ListFlowClusterHost API asynchronously
// api document: https://help.aliyun.com/api/emr/listflowclusterhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowClusterHostWithChan(request *ListFlowClusterHostRequest) (<-chan *ListFlowClusterHostResponse, <-chan error) {
	responseChan := make(chan *ListFlowClusterHostResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListFlowClusterHost(request)
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

// ListFlowClusterHostWithCallback invokes the emr.ListFlowClusterHost API asynchronously
// api document: https://help.aliyun.com/api/emr/listflowclusterhost.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListFlowClusterHostWithCallback(request *ListFlowClusterHostRequest, callback func(response *ListFlowClusterHostResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListFlowClusterHostResponse
		var err error
		defer close(result)
		response, err = client.ListFlowClusterHost(request)
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

// ListFlowClusterHostRequest is the request struct for api ListFlowClusterHost
type ListFlowClusterHostRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
	ProjectId string `position:"Query" name:"ProjectId"`
}

// ListFlowClusterHostResponse is the response struct for api ListFlowClusterHost
type ListFlowClusterHostResponse struct {
	*responses.BaseResponse
	RequestId string                        `json:"RequestId" xml:"RequestId"`
	HostList  HostListInListFlowClusterHost `json:"HostList" xml:"HostList"`
}

// CreateListFlowClusterHostRequest creates a request to invoke ListFlowClusterHost API
func CreateListFlowClusterHostRequest() (request *ListFlowClusterHostRequest) {
	request = &ListFlowClusterHostRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ListFlowClusterHost", "emr", "openAPI")
	return
}

// CreateListFlowClusterHostResponse creates a response to parse from ListFlowClusterHost response
func CreateListFlowClusterHostResponse() (response *ListFlowClusterHostResponse) {
	response = &ListFlowClusterHostResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
