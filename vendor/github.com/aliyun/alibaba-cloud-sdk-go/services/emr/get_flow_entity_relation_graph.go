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

// GetFlowEntityRelationGraph invokes the emr.GetFlowEntityRelationGraph API synchronously
// api document: https://help.aliyun.com/api/emr/getflowentityrelationgraph.html
func (client *Client) GetFlowEntityRelationGraph(request *GetFlowEntityRelationGraphRequest) (response *GetFlowEntityRelationGraphResponse, err error) {
	response = CreateGetFlowEntityRelationGraphResponse()
	err = client.DoAction(request, response)
	return
}

// GetFlowEntityRelationGraphWithChan invokes the emr.GetFlowEntityRelationGraph API asynchronously
// api document: https://help.aliyun.com/api/emr/getflowentityrelationgraph.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFlowEntityRelationGraphWithChan(request *GetFlowEntityRelationGraphRequest) (<-chan *GetFlowEntityRelationGraphResponse, <-chan error) {
	responseChan := make(chan *GetFlowEntityRelationGraphResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFlowEntityRelationGraph(request)
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

// GetFlowEntityRelationGraphWithCallback invokes the emr.GetFlowEntityRelationGraph API asynchronously
// api document: https://help.aliyun.com/api/emr/getflowentityrelationgraph.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFlowEntityRelationGraphWithCallback(request *GetFlowEntityRelationGraphRequest, callback func(response *GetFlowEntityRelationGraphResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFlowEntityRelationGraphResponse
		var err error
		defer close(result)
		response, err = client.GetFlowEntityRelationGraph(request)
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

// GetFlowEntityRelationGraphRequest is the request struct for api GetFlowEntityRelationGraph
type GetFlowEntityRelationGraphRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageCount       requests.Integer `position:"Query" name:"PageCount"`
	OrderMode       string           `position:"Query" name:"OrderMode"`
	EntityId        string           `position:"Query" name:"EntityId"`
	PageNumber      requests.Integer `position:"Query" name:"PageNumber"`
	Limit           requests.Integer `position:"Query" name:"Limit"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	Relationship    string           `position:"Query" name:"Relationship"`
	CurrentSize     requests.Integer `position:"Query" name:"CurrentSize"`
	OrderField      string           `position:"Query" name:"OrderField"`
	Direction       string           `position:"Query" name:"Direction"`
	EntityGroupId   string           `position:"Query" name:"EntityGroupId"`
	EntityType      string           `position:"Query" name:"EntityType"`
}

// GetFlowEntityRelationGraphResponse is the response struct for api GetFlowEntityRelationGraph
type GetFlowEntityRelationGraphResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateGetFlowEntityRelationGraphRequest creates a request to invoke GetFlowEntityRelationGraph API
func CreateGetFlowEntityRelationGraphRequest() (request *GetFlowEntityRelationGraphRequest) {
	request = &GetFlowEntityRelationGraphRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "GetFlowEntityRelationGraph", "emr", "openAPI")
	return
}

// CreateGetFlowEntityRelationGraphResponse creates a response to parse from GetFlowEntityRelationGraph response
func CreateGetFlowEntityRelationGraphResponse() (response *GetFlowEntityRelationGraphResponse) {
	response = &GetFlowEntityRelationGraphResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
