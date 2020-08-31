package rds

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

// RecoveryDBInstance invokes the rds.RecoveryDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/recoverydbinstance.html
func (client *Client) RecoveryDBInstance(request *RecoveryDBInstanceRequest) (response *RecoveryDBInstanceResponse, err error) {
	response = CreateRecoveryDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RecoveryDBInstanceWithChan invokes the rds.RecoveryDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/recoverydbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecoveryDBInstanceWithChan(request *RecoveryDBInstanceRequest) (<-chan *RecoveryDBInstanceResponse, <-chan error) {
	responseChan := make(chan *RecoveryDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecoveryDBInstance(request)
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

// RecoveryDBInstanceWithCallback invokes the rds.RecoveryDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/recoverydbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecoveryDBInstanceWithCallback(request *RecoveryDBInstanceRequest, callback func(response *RecoveryDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecoveryDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.RecoveryDBInstance(request)
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

// RecoveryDBInstanceRequest is the request struct for api RecoveryDBInstance
type RecoveryDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId       requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceStorage     requests.Integer `position:"Query" name:"DBInstanceStorage"`
	ClientToken           string           `position:"Query" name:"ClientToken"`
	ResourceGroupId       string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceDescription string           `position:"Query" name:"DBInstanceDescription"`
	DBInstanceId          string           `position:"Query" name:"DBInstanceId"`
	DBInstanceStorageType string           `position:"Query" name:"DBInstanceStorageType"`
	RestoreTime           string           `position:"Query" name:"RestoreTime"`
	Period                string           `position:"Query" name:"Period"`
	ResourceOwnerAccount  string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId              string           `position:"Query" name:"BackupId"`
	OwnerAccount          string           `position:"Query" name:"OwnerAccount"`
	OwnerId               requests.Integer `position:"Query" name:"OwnerId"`
	UsedTime              string           `position:"Query" name:"UsedTime"`
	DBInstanceClass       string           `position:"Query" name:"DBInstanceClass"`
	DbNames               string           `position:"Query" name:"DbNames"`
	VSwitchId             string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress      string           `position:"Query" name:"PrivateIpAddress"`
	TargetDBInstanceId    string           `position:"Query" name:"TargetDBInstanceId"`
	VPCId                 string           `position:"Query" name:"VPCId"`
	PayType               string           `position:"Query" name:"PayType"`
	InstanceNetworkType   string           `position:"Query" name:"InstanceNetworkType"`
}

// RecoveryDBInstanceResponse is the response struct for api RecoveryDBInstance
type RecoveryDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	OrderId      string `json:"OrderId" xml:"OrderId"`
}

// CreateRecoveryDBInstanceRequest creates a request to invoke RecoveryDBInstance API
func CreateRecoveryDBInstanceRequest() (request *RecoveryDBInstanceRequest) {
	request = &RecoveryDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RecoveryDBInstance", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecoveryDBInstanceResponse creates a response to parse from RecoveryDBInstance response
func CreateRecoveryDBInstanceResponse() (response *RecoveryDBInstanceResponse) {
	response = &RecoveryDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
