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

// OnsTraceQueryByMsgId invokes the ons.OnsTraceQueryByMsgId API synchronously
// api document: https://help.aliyun.com/api/ons/onstracequerybymsgid.html
func (client *Client) OnsTraceQueryByMsgId(request *OnsTraceQueryByMsgIdRequest) (response *OnsTraceQueryByMsgIdResponse, err error) {
	response = CreateOnsTraceQueryByMsgIdResponse()
	err = client.DoAction(request, response)
	return
}

// OnsTraceQueryByMsgIdWithChan invokes the ons.OnsTraceQueryByMsgId API asynchronously
// api document: https://help.aliyun.com/api/ons/onstracequerybymsgid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsTraceQueryByMsgIdWithChan(request *OnsTraceQueryByMsgIdRequest) (<-chan *OnsTraceQueryByMsgIdResponse, <-chan error) {
	responseChan := make(chan *OnsTraceQueryByMsgIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.OnsTraceQueryByMsgId(request)
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

// OnsTraceQueryByMsgIdWithCallback invokes the ons.OnsTraceQueryByMsgId API asynchronously
// api document: https://help.aliyun.com/api/ons/onstracequerybymsgid.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) OnsTraceQueryByMsgIdWithCallback(request *OnsTraceQueryByMsgIdRequest, callback func(response *OnsTraceQueryByMsgIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *OnsTraceQueryByMsgIdResponse
		var err error
		defer close(result)
		response, err = client.OnsTraceQueryByMsgId(request)
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

// OnsTraceQueryByMsgIdRequest is the request struct for api OnsTraceQueryByMsgId
type OnsTraceQueryByMsgIdRequest struct {
	*requests.RpcRequest
	PreventCache requests.Integer `position:"Query" name:"PreventCache"`
	InstanceId   string           `position:"Query" name:"InstanceId"`
	Topic        string           `position:"Query" name:"Topic"`
	MsgId        string           `position:"Query" name:"MsgId"`
	EndTime      requests.Integer `position:"Query" name:"EndTime"`
	BeginTime    requests.Integer `position:"Query" name:"BeginTime"`
}

// OnsTraceQueryByMsgIdResponse is the response struct for api OnsTraceQueryByMsgId
type OnsTraceQueryByMsgIdResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	HelpUrl   string `json:"HelpUrl" xml:"HelpUrl"`
	QueryId   string `json:"QueryId" xml:"QueryId"`
}

// CreateOnsTraceQueryByMsgIdRequest creates a request to invoke OnsTraceQueryByMsgId API
func CreateOnsTraceQueryByMsgIdRequest() (request *OnsTraceQueryByMsgIdRequest) {
	request = &OnsTraceQueryByMsgIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ons", "2019-02-14", "OnsTraceQueryByMsgId", "ons", "openAPI")
	return
}

// CreateOnsTraceQueryByMsgIdResponse creates a response to parse from OnsTraceQueryByMsgId response
func CreateOnsTraceQueryByMsgIdResponse() (response *OnsTraceQueryByMsgIdResponse) {
	response = &OnsTraceQueryByMsgIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
