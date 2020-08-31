package polardb

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

// DescribeDBClusterParameters invokes the polardb.DescribeDBClusterParameters API synchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterparameters.html
func (client *Client) DescribeDBClusterParameters(request *DescribeDBClusterParametersRequest) (response *DescribeDBClusterParametersResponse, err error) {
	response = CreateDescribeDBClusterParametersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterParametersWithChan invokes the polardb.DescribeDBClusterParameters API asynchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterparameters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterParametersWithChan(request *DescribeDBClusterParametersRequest) (<-chan *DescribeDBClusterParametersResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterParametersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterParameters(request)
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

// DescribeDBClusterParametersWithCallback invokes the polardb.DescribeDBClusterParameters API asynchronously
// api document: https://help.aliyun.com/api/polardb/describedbclusterparameters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBClusterParametersWithCallback(request *DescribeDBClusterParametersRequest, callback func(response *DescribeDBClusterParametersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterParametersResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterParameters(request)
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

// DescribeDBClusterParametersRequest is the request struct for api DescribeDBClusterParameters
type DescribeDBClusterParametersRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterParametersResponse is the response struct for api DescribeDBClusterParameters
type DescribeDBClusterParametersResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	Engine            string            `json:"Engine" xml:"Engine"`
	DBType            string            `json:"DBType" xml:"DBType"`
	DBVersion         string            `json:"DBVersion" xml:"DBVersion"`
	RunningParameters RunningParameters `json:"RunningParameters" xml:"RunningParameters"`
}

// CreateDescribeDBClusterParametersRequest creates a request to invoke DescribeDBClusterParameters API
func CreateDescribeDBClusterParametersRequest() (request *DescribeDBClusterParametersRequest) {
	request = &DescribeDBClusterParametersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterParameters", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterParametersResponse creates a response to parse from DescribeDBClusterParameters response
func CreateDescribeDBClusterParametersResponse() (response *DescribeDBClusterParametersResponse) {
	response = &DescribeDBClusterParametersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
