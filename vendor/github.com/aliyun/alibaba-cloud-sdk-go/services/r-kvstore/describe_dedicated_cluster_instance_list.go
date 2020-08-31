package r_kvstore

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

// DescribeDedicatedClusterInstanceList invokes the r_kvstore.DescribeDedicatedClusterInstanceList API synchronously
// api document: https://help.aliyun.com/api/r-kvstore/describededicatedclusterinstancelist.html
func (client *Client) DescribeDedicatedClusterInstanceList(request *DescribeDedicatedClusterInstanceListRequest) (response *DescribeDedicatedClusterInstanceListResponse, err error) {
	response = CreateDescribeDedicatedClusterInstanceListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDedicatedClusterInstanceListWithChan invokes the r_kvstore.DescribeDedicatedClusterInstanceList API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describededicatedclusterinstancelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedClusterInstanceListWithChan(request *DescribeDedicatedClusterInstanceListRequest) (<-chan *DescribeDedicatedClusterInstanceListResponse, <-chan error) {
	responseChan := make(chan *DescribeDedicatedClusterInstanceListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDedicatedClusterInstanceList(request)
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

// DescribeDedicatedClusterInstanceListWithCallback invokes the r_kvstore.DescribeDedicatedClusterInstanceList API asynchronously
// api document: https://help.aliyun.com/api/r-kvstore/describededicatedclusterinstancelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedClusterInstanceListWithCallback(request *DescribeDedicatedClusterInstanceListRequest, callback func(response *DescribeDedicatedClusterInstanceListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDedicatedClusterInstanceListResponse
		var err error
		defer close(result)
		response, err = client.DescribeDedicatedClusterInstanceList(request)
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

// DescribeDedicatedClusterInstanceListRequest is the request struct for api DescribeDedicatedClusterInstanceList
type DescribeDedicatedClusterInstanceListRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EngineVersion        string           `position:"Query" name:"EngineVersion"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Engine               string           `position:"Query" name:"Engine"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	InstanceStatus       requests.Integer `position:"Query" name:"InstanceStatus"`
	DedicatedHostName    string           `position:"Query" name:"DedicatedHostName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	InstanceNetType      string           `position:"Query" name:"InstanceNetType"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

// DescribeDedicatedClusterInstanceListResponse is the response struct for api DescribeDedicatedClusterInstanceList
type DescribeDedicatedClusterInstanceListResponse struct {
	*responses.BaseResponse
	PageNumber int             `json:"PageNumber" xml:"PageNumber"`
	PageSize   int             `json:"PageSize" xml:"PageSize"`
	RequestId  string          `json:"RequestId" xml:"RequestId"`
	TotalCount int             `json:"TotalCount" xml:"TotalCount"`
	Instances  []InstancesItem `json:"Instances" xml:"Instances"`
}

// CreateDescribeDedicatedClusterInstanceListRequest creates a request to invoke DescribeDedicatedClusterInstanceList API
func CreateDescribeDedicatedClusterInstanceListRequest() (request *DescribeDedicatedClusterInstanceListRequest) {
	request = &DescribeDedicatedClusterInstanceListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "DescribeDedicatedClusterInstanceList", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDedicatedClusterInstanceListResponse creates a response to parse from DescribeDedicatedClusterInstanceList response
func CreateDescribeDedicatedClusterInstanceListResponse() (response *DescribeDedicatedClusterInstanceListResponse) {
	response = &DescribeDedicatedClusterInstanceListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
