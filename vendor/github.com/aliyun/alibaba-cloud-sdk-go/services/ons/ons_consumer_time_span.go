package ons

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

// OnsConsumerTimeSpan invokes the ons.OnsConsumerTimeSpan API synchronously
// api document: https://help.aliyun.com/api/ons/onsconsumertimespan.html
func (client *Client) OnsConsumerTimeSpan(request *OnsConsumerTimeSpanRequest) (response *OnsConsumerTimeSpanResponse, err error) {
	response = CreateOnsConsumerTimeSpanResponse()
	err = client.DoAction(request, response)
	return
}

// OnsConsumerTimeSpanWithChan invokes the ons.OnsConsumerTimeSpan API asynchronously
// api document: https://help.aliyun.com/api/ons/onsconsumertimespan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsConsumerTimeSpanWithChan(request *OnsConsumerTimeSpanRequest) (<-chan *OnsConsumerTimeSpanResponse, <-chan error) {
	responseChan := make(chan *OnsConsumerTimeSpanResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsConsumerTimeSpan(request)
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

// OnsConsumerTimeSpanWithCallback invokes the ons.OnsConsumerTimeSpan API asynchronously
// api document: https://help.aliyun.com/api/ons/onsconsumertimespan.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsConsumerTimeSpanWithCallback(request *OnsConsumerTimeSpanRequest, callback func(response *OnsConsumerTimeSpanResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsConsumerTimeSpanResponse
		var err error
		defer close(result)
		response, err = client.OnsConsumerTimeSpan(request)
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

// OnsConsumerTimeSpanRequest is the request struct for api OnsConsumerTimeSpan
type OnsConsumerTimeSpanRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	InstanceId string `position:"Query" name:"InstanceId"`
	Topic      string `position:"Query" name:"Topic"`
}

// OnsConsumerTimeSpanResponse is the response struct for api OnsConsumerTimeSpan
type OnsConsumerTimeSpanResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateOnsConsumerTimeSpanRequest creates a request to invoke OnsConsumerTimeSpan API
func CreateOnsConsumerTimeSpanRequest() (request *OnsConsumerTimeSpanRequest) {
	request = &OnsConsumerTimeSpanRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsConsumerTimeSpan", "ons", "openAPI")
	request.Method = requests.POST
	return
}

// CreateOnsConsumerTimeSpanResponse creates a response to parse from OnsConsumerTimeSpan response
func CreateOnsConsumerTimeSpanResponse() (response *OnsConsumerTimeSpanResponse) {
	response = &OnsConsumerTimeSpanResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
