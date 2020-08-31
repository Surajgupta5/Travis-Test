package bssopenapi

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

// AllocateCostUnitResource invokes the bssopenapi.AllocateCostUnitResource API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/allocatecostunitresource.html
func (client *Client) AllocateCostUnitResource(request *AllocateCostUnitResourceRequest) (response *AllocateCostUnitResourceResponse, err error) {
	response = CreateAllocateCostUnitResourceResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateCostUnitResourceWithChan invokes the bssopenapi.AllocateCostUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/allocatecostunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateCostUnitResourceWithChan(request *AllocateCostUnitResourceRequest) (<-chan *AllocateCostUnitResourceResponse, <-chan error) {
	responseChan := make(chan *AllocateCostUnitResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateCostUnitResource(request)
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

// AllocateCostUnitResourceWithCallback invokes the bssopenapi.AllocateCostUnitResource API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/allocatecostunitresource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateCostUnitResourceWithCallback(request *AllocateCostUnitResourceRequest, callback func(response *AllocateCostUnitResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateCostUnitResourceResponse
		var err error
		defer close(result)
		response, err = client.AllocateCostUnitResource(request)
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

// AllocateCostUnitResourceRequest is the request struct for api AllocateCostUnitResource
type AllocateCostUnitResourceRequest struct {
	*requests.RpcRequest
	ResourceInstanceList *[]AllocateCostUnitResourceResourceInstanceList `position:"Query" name:"ResourceInstanceList"  type:"Repeated"`
	FromUnitId           requests.Integer                                `position:"Query" name:"FromUnitId"`
	ToUnitId             requests.Integer                                `position:"Query" name:"ToUnitId"`
	FromUnitUserId       requests.Integer                                `position:"Query" name:"FromUnitUserId"`
	ToUnitUserId         requests.Integer                                `position:"Query" name:"ToUnitUserId"`
}

// AllocateCostUnitResourceResourceInstanceList is a repeated param struct in AllocateCostUnitResourceRequest
type AllocateCostUnitResourceResourceInstanceList struct {
	ResourceId     string `name:"ResourceId"`
	CommodityCode  string `name:"CommodityCode"`
	ResourceUserId string `name:"ResourceUserId"`
}

// AllocateCostUnitResourceResponse is the response struct for api AllocateCostUnitResource
type AllocateCostUnitResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateAllocateCostUnitResourceRequest creates a request to invoke AllocateCostUnitResource API
func CreateAllocateCostUnitResourceRequest() (request *AllocateCostUnitResourceRequest) {
	request = &AllocateCostUnitResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "AllocateCostUnitResource", "", "")
	request.Method = requests.POST
	return
}

// CreateAllocateCostUnitResourceResponse creates a response to parse from AllocateCostUnitResource response
func CreateAllocateCostUnitResourceResponse() (response *AllocateCostUnitResourceResponse) {
	response = &AllocateCostUnitResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
