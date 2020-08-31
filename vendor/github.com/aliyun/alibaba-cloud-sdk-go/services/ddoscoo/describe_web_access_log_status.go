package ddoscoo

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

// DescribeWebAccessLogStatus invokes the ddoscoo.DescribeWebAccessLogStatus API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/describewebaccesslogstatus.html
func (client *Client) DescribeWebAccessLogStatus(request *DescribeWebAccessLogStatusRequest) (response *DescribeWebAccessLogStatusResponse, err error) {
	response = CreateDescribeWebAccessLogStatusResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeWebAccessLogStatusWithChan invokes the ddoscoo.DescribeWebAccessLogStatus API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describewebaccesslogstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebAccessLogStatusWithChan(request *DescribeWebAccessLogStatusRequest) (<-chan *DescribeWebAccessLogStatusResponse, <-chan error) {
	responseChan := make(chan *DescribeWebAccessLogStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebAccessLogStatus(request)
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

// DescribeWebAccessLogStatusWithCallback invokes the ddoscoo.DescribeWebAccessLogStatus API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/describewebaccesslogstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebAccessLogStatusWithCallback(request *DescribeWebAccessLogStatusRequest, callback func(response *DescribeWebAccessLogStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebAccessLogStatusResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebAccessLogStatus(request)
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

// DescribeWebAccessLogStatusRequest is the request struct for api DescribeWebAccessLogStatus
type DescribeWebAccessLogStatusRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Domain          string `position:"Query" name:"Domain"`
	Lang            string `position:"Query" name:"Lang"`
}

// DescribeWebAccessLogStatusResponse is the response struct for api DescribeWebAccessLogStatus
type DescribeWebAccessLogStatusResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	SlsStatus   bool   `json:"SlsStatus" xml:"SlsStatus"`
	SlsLogstore string `json:"SlsLogstore" xml:"SlsLogstore"`
	SlsProject  string `json:"SlsProject" xml:"SlsProject"`
}

// CreateDescribeWebAccessLogStatusRequest creates a request to invoke DescribeWebAccessLogStatus API
func CreateDescribeWebAccessLogStatusRequest() (request *DescribeWebAccessLogStatusRequest) {
	request = &DescribeWebAccessLogStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "DescribeWebAccessLogStatus", "ddoscoo", "openAPI")
	return
}

// CreateDescribeWebAccessLogStatusResponse creates a response to parse from DescribeWebAccessLogStatus response
func CreateDescribeWebAccessLogStatusResponse() (response *DescribeWebAccessLogStatusResponse) {
	response = &DescribeWebAccessLogStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
