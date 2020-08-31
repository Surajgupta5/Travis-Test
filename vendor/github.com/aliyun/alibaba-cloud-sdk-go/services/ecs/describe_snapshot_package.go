package ecs

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

// DescribeSnapshotPackage invokes the ecs.DescribeSnapshotPackage API synchronously
// api document: https://help.aliyun.com/api/ecs/describesnapshotpackage.html
func (client *Client) DescribeSnapshotPackage(request *DescribeSnapshotPackageRequest) (response *DescribeSnapshotPackageResponse, err error) {
	response = CreateDescribeSnapshotPackageResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSnapshotPackageWithChan invokes the ecs.DescribeSnapshotPackage API asynchronously
// api document: https://help.aliyun.com/api/ecs/describesnapshotpackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSnapshotPackageWithChan(request *DescribeSnapshotPackageRequest) (<-chan *DescribeSnapshotPackageResponse, <-chan error) {
	responseChan := make(chan *DescribeSnapshotPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSnapshotPackage(request)
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

// DescribeSnapshotPackageWithCallback invokes the ecs.DescribeSnapshotPackage API asynchronously
// api document: https://help.aliyun.com/api/ecs/describesnapshotpackage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSnapshotPackageWithCallback(request *DescribeSnapshotPackageRequest, callback func(response *DescribeSnapshotPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSnapshotPackageResponse
		var err error
		defer close(result)
		response, err = client.DescribeSnapshotPackage(request)
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

// DescribeSnapshotPackageRequest is the request struct for api DescribeSnapshotPackage
type DescribeSnapshotPackageRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSnapshotPackageResponse is the response struct for api DescribeSnapshotPackage
type DescribeSnapshotPackageResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	TotalCount       int              `json:"TotalCount" xml:"TotalCount"`
	PageNumber       int              `json:"PageNumber" xml:"PageNumber"`
	PageSize         int              `json:"PageSize" xml:"PageSize"`
	SnapshotPackages SnapshotPackages `json:"SnapshotPackages" xml:"SnapshotPackages"`
}

// CreateDescribeSnapshotPackageRequest creates a request to invoke DescribeSnapshotPackage API
func CreateDescribeSnapshotPackageRequest() (request *DescribeSnapshotPackageRequest) {
	request = &DescribeSnapshotPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeSnapshotPackage", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSnapshotPackageResponse creates a response to parse from DescribeSnapshotPackage response
func CreateDescribeSnapshotPackageResponse() (response *DescribeSnapshotPackageResponse) {
	response = &DescribeSnapshotPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
