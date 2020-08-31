package vpc

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

// ModifyIPv6TranslatorAclListEntry invokes the vpc.ModifyIPv6TranslatorAclListEntry API synchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6translatoracllistentry.html
func (client *Client) ModifyIPv6TranslatorAclListEntry(request *ModifyIPv6TranslatorAclListEntryRequest) (response *ModifyIPv6TranslatorAclListEntryResponse, err error) {
	response = CreateModifyIPv6TranslatorAclListEntryResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyIPv6TranslatorAclListEntryWithChan invokes the vpc.ModifyIPv6TranslatorAclListEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6translatoracllistentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIPv6TranslatorAclListEntryWithChan(request *ModifyIPv6TranslatorAclListEntryRequest) (<-chan *ModifyIPv6TranslatorAclListEntryResponse, <-chan error) {
	responseChan := make(chan *ModifyIPv6TranslatorAclListEntryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyIPv6TranslatorAclListEntry(request)
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

// ModifyIPv6TranslatorAclListEntryWithCallback invokes the vpc.ModifyIPv6TranslatorAclListEntry API asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyipv6translatoracllistentry.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyIPv6TranslatorAclListEntryWithCallback(request *ModifyIPv6TranslatorAclListEntryRequest, callback func(response *ModifyIPv6TranslatorAclListEntryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyIPv6TranslatorAclListEntryResponse
		var err error
		defer close(result)
		response, err = client.ModifyIPv6TranslatorAclListEntry(request)
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

// ModifyIPv6TranslatorAclListEntryRequest is the request struct for api ModifyIPv6TranslatorAclListEntry
type ModifyIPv6TranslatorAclListEntryRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AclId                string           `position:"Query" name:"AclId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	AclEntryComment      string           `position:"Query" name:"AclEntryComment"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	AclEntryId           string           `position:"Query" name:"AclEntryId"`
}

// ModifyIPv6TranslatorAclListEntryResponse is the response struct for api ModifyIPv6TranslatorAclListEntry
type ModifyIPv6TranslatorAclListEntryResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyIPv6TranslatorAclListEntryRequest creates a request to invoke ModifyIPv6TranslatorAclListEntry API
func CreateModifyIPv6TranslatorAclListEntryRequest() (request *ModifyIPv6TranslatorAclListEntryRequest) {
	request = &ModifyIPv6TranslatorAclListEntryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyIPv6TranslatorAclListEntry", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyIPv6TranslatorAclListEntryResponse creates a response to parse from ModifyIPv6TranslatorAclListEntry response
func CreateModifyIPv6TranslatorAclListEntryResponse() (response *ModifyIPv6TranslatorAclListEntryResponse) {
	response = &ModifyIPv6TranslatorAclListEntryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
