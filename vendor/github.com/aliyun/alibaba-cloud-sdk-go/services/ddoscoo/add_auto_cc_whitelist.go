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

// AddAutoCcWhitelist invokes the ddoscoo.AddAutoCcWhitelist API synchronously
// api document: https://help.aliyun.com/api/ddoscoo/addautoccwhitelist.html
func (client *Client) AddAutoCcWhitelist(request *AddAutoCcWhitelistRequest) (response *AddAutoCcWhitelistResponse, err error) {
	response = CreateAddAutoCcWhitelistResponse()
	err = client.DoAction(request, response)
	return
}

// AddAutoCcWhitelistWithChan invokes the ddoscoo.AddAutoCcWhitelist API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/addautoccwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddAutoCcWhitelistWithChan(request *AddAutoCcWhitelistRequest) (<-chan *AddAutoCcWhitelistResponse, <-chan error) {
	responseChan := make(chan *AddAutoCcWhitelistResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddAutoCcWhitelist(request)
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

// AddAutoCcWhitelistWithCallback invokes the ddoscoo.AddAutoCcWhitelist API asynchronously
// api document: https://help.aliyun.com/api/ddoscoo/addautoccwhitelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddAutoCcWhitelistWithCallback(request *AddAutoCcWhitelistRequest, callback func(response *AddAutoCcWhitelistResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddAutoCcWhitelistResponse
		var err error
		defer close(result)
		response, err = client.AddAutoCcWhitelist(request)
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

// AddAutoCcWhitelistRequest is the request struct for api AddAutoCcWhitelist
type AddAutoCcWhitelistRequest struct {
	*requests.RpcRequest
	ExpireTime requests.Integer `position:"Query" name:"ExpireTime"`
	Whitelist  string           `position:"Query" name:"Whitelist"`
	InstanceId string           `position:"Query" name:"InstanceId"`
	SourceIp   string           `position:"Query" name:"SourceIp"`
}

// AddAutoCcWhitelistResponse is the response struct for api AddAutoCcWhitelist
type AddAutoCcWhitelistResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddAutoCcWhitelistRequest creates a request to invoke AddAutoCcWhitelist API
func CreateAddAutoCcWhitelistRequest() (request *AddAutoCcWhitelistRequest) {
	request = &AddAutoCcWhitelistRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2020-01-01", "AddAutoCcWhitelist", "ddoscoo", "openAPI")
	return
}

// CreateAddAutoCcWhitelistResponse creates a response to parse from AddAutoCcWhitelist response
func CreateAddAutoCcWhitelistResponse() (response *AddAutoCcWhitelistResponse) {
	response = &AddAutoCcWhitelistResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
