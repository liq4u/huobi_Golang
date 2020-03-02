package client

import (
	"../../internal"
	"../../internal/requestbuilder"
	"../../pkg/getrequest"
	"../../pkg/postrequest"
	"../../pkg/response/order"
	"encoding/json"
	"fmt"
)

type OrderClient struct {
	privateUrlBuilder *requestbuilder.PrivateUrlBuilder
}

func (p *OrderClient) Init(accessKey string, secretKey string, host string) *OrderClient {
	p.privateUrlBuilder = new(requestbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return p
}

func (p *OrderClient) PlaceOrder(request *postrequest.PlaceOrderRequest) (*order.PlaceOrderResponse, error) {
	postBody, jsonErr := postrequest.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v1/order/orders/place", nil)
	postResp, postErr := internal.HttpPost(url, postBody)
	if postErr != nil {
		return nil, postErr
	}

	result := order.PlaceOrderResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) PlaceOrders(request []postrequest.PlaceOrderRequest) (*order.PlaceOrdersResponse, error) {

	postBody, jsonErr := postrequest.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v1/order/batch-orders", nil)
	postResp, postErr := internal.HttpPost(url, string(postBody))
	if postErr != nil {
		return nil, postErr
	}

	result := order.PlaceOrdersResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) CancelOrderById(orderId string) (*order.CancelOrderByIdResponse, error) {
	path := fmt.Sprintf("/v1/order/orders/%s/submitcancel", orderId)
	url := p.privateUrlBuilder.Build("POST", path, nil)
	postResp, postErr := internal.HttpPost(url, "")
	if postErr != nil {
		return nil, postErr
	}

	result := order.CancelOrderByIdResponse{}
	jsonErr := json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) CancelOrderByClientOrderId(clientOrderId string) (*order.CancelOrderByClientResponse, error) {
	url := p.privateUrlBuilder.Build("POST", "/v1/order/orders/submitCancelClientOrder", nil)
	body := fmt.Sprintf("{\"client-order-id\":\"%s\"}", clientOrderId)
	postResp, postErr := internal.HttpPost(url, body)
	if postErr != nil {
		return nil, postErr
	}

	result := order.CancelOrderByClientResponse{}
	jsonErr := json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetOpenOrders(request *getrequest.GetRequest) (*order.GetOpenOrdersResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/order/openOrders", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetOpenOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) CancelOrdersByCriteria(request *postrequest.CancelOrdersByCriteriaRequest) (*order.CancelOrdersByCriteriaResponse, error) {
	postBody, jsonErr := postrequest.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v1/order/orders/batchCancelOpenOrders", nil)
	postResp, postErr := internal.HttpPost(url, string(postBody))
	if postErr != nil {
		return nil, postErr
	}

	result := order.CancelOrdersByCriteriaResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) CancelOrdersByIds(request *postrequest.CancelOrdersByIdsRequest) (*order.CancelOrdersByIdsResponse, error) {
	postBody, jsonErr := postrequest.ToJson(request)

	url := p.privateUrlBuilder.Build("POST", "/v1/order/orders/batchcancel", nil)
	postResp, postErr := internal.HttpPost(url, string(postBody))
	if postErr != nil {
		return nil, postErr
	}

	result := order.CancelOrdersByIdsResponse{}
	jsonErr = json.Unmarshal([]byte(postResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetOrderById(orderId string) (*order.GetOrderResponse, error) {
	path := fmt.Sprintf("/v1/order/orders/%s", orderId)
	url := p.privateUrlBuilder.Build("GET", path, nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetOrderByCriteria(request *getrequest.GetRequest) (*order.GetOrderResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/order/orders/getClientOrder", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetMatchResultsById(orderId string) (*order.GetMatchResultsResponse, error) {
	path := fmt.Sprintf("/v1/order/orders/%s/matchresults", orderId)
	url := p.privateUrlBuilder.Build("GET", path, nil)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetMatchResultsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetHistoryOrders(request *getrequest.GetRequest) (*order.GetHistoryOrdersResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/order/orders", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetHistoryOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetLast48hOrders(request *getrequest.GetRequest) (*order.GetHistoryOrdersResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/order/history", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetHistoryOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetMatchResultsByCriteria(request *getrequest.GetRequest) (*order.GetMatchResultsResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v1/order/matchresults", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetMatchResultsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) GetFeeRate(request *getrequest.GetRequest) (*order.GetFeeResponse, error) {
	url := p.privateUrlBuilder.Build("GET", "/v2/reference/transact-fee-rate", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := order.GetFeeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return &result, nil
}

func (p *OrderClient) Test() {

	request := new(getrequest.GetRequest).Init()
	request.AddParam("currency", "pax")
	request.AddParam("amount", "1000")

	request.AddParam("type", "buy")

	url := p.privateUrlBuilder.Build("GET", "/v1/stable-coin/quote", request)
	getResp, getErr := internal.HttpGet(url)
	fmt.Println(getResp)
	fmt.Println(getErr)

}