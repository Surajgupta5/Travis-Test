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

// DescribeSecurityGroupAttribute invokes the emr.DescribeSecurityGroupAttribute API synchronously
// api document: https://help.aliyun.com/api/emr/describesecuritygroupattribute.html
func (client *Client) DescribeSecurityGroupAttribute(request *DescribeSecurityGroupAttributeRequest) (response *DescribeSecurityGroupAttributeResponse, err error) {
	response = CreateDescribeSecurityGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSecurityGroupAttributeWithChan invokes the emr.DescribeSecurityGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/emr/describesecuritygroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSecurityGroupAttributeWithChan(request *DescribeSecurityGroupAttributeRequest) (<-chan *DescribeSecurityGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeSecurityGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSecurityGroupAttribute(request)
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

// DescribeSecurityGroupAttributeWithCallback invokes the emr.DescribeSecurityGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/emr/describesecuritygroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSecurityGroupAttributeWithCallback(request *DescribeSecurityGroupAttributeRequest, callback func(response *DescribeSecurityGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSecurityGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeSecurityGroupAttribute(request)
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

// DescribeSecurityGroupAttributeRequest is the request struct for api DescribeSecurityGroupAttribute
type DescribeSecurityGroupAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClusterId       string           `position:"Query" name:"ClusterId"`
}

// DescribeSecurityGroupAttributeResponse is the response struct for api DescribeSecurityGroupAttribute
type DescribeSecurityGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId                  string                     `json:"RequestId" xml:"RequestId"`
	AvailableGroupList         AvailableGroupList         `json:"AvailableGroupList" xml:"AvailableGroupList"`
	SecurityGroupAttributeList SecurityGroupAttributeList `json:"SecurityGroupAttributeList" xml:"SecurityGroupAttributeList"`
}

// CreateDescribeSecurityGroupAttributeRequest creates a request to invoke DescribeSecurityGroupAttribute API
func CreateDescribeSecurityGroupAttributeRequest() (request *DescribeSecurityGroupAttributeRequest) {
	request = &DescribeSecurityGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "DescribeSecurityGroupAttribute", "emr", "openAPI")
	return
}

// CreateDescribeSecurityGroupAttributeResponse creates a response to parse from DescribeSecurityGroupAttribute response
func CreateDescribeSecurityGroupAttributeResponse() (response *DescribeSecurityGroupAttributeResponse) {
	response = &DescribeSecurityGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
