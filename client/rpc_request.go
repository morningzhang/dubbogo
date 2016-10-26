/******************************************************
# DESC    : rpc client request
# AUTHOR  : Alex Stocks
# VERSION : 1.0
# LICENCE : Apache Licence 2.0
# EMAIL   : alexstocks@foxmail.com
# MOD     : 2016-06-30 10:45
# FILE    : rpc_request.go
******************************************************/

package client

type rpcRequest struct {
	service     string
	method      string
	contentType string
	request     interface{}
	opts        RequestOptions
}

func newRpcRequest(service, method string, request interface{}, contentType string, reqOpts ...RequestOption) Request {
	var opts RequestOptions

	for _, o := range reqOpts {
		o(&opts)
	}

	return &rpcRequest{
		service:     service,
		method:      method,
		request:     request,
		contentType: contentType,
		opts:        opts,
	}
}

func (r *rpcRequest) ContentType() string {
	return r.contentType
}

func (r *rpcRequest) Service() string {
	return r.service
}

func (r *rpcRequest) Method() string {
	return r.method
}

func (r *rpcRequest) Request() interface{} {
	return r.request
}

func (r *rpcRequest) Stream() bool {
	return r.opts.Stream
}
