package httpkernel

import (
    "chime/components/httpcontext"
    "net/http"
    dispatcher "chime/components/eventdispatcher"
)

const (
	KERNEL_EVENTS_REQUEST    = "kernel.request"
	KERNEL_EVENTS_VIEW       = "kernel.view"
	KERNEL_EVENTS_RESPONSE   = "kernel.response"
	KERNEL_EVENTS_CONTROLLER = "kernel.controller"
	MASTER_REQUEST           = 1
	SUB_REQUEST              = 2
)

type HttpKernel struct {
	dispatcher *dispatcher.EventDispatcher
	resolver   *Resolver
}

func (this *HttpKernel) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    reqContext := httpcontext.NewRequest(request)
    this.HandleMasterRequest(reqContext).Send(writer)
}

func (this *HttpKernel) HandleMasterRequest(req *httpcontext.Request) httpcontext.Responser {
	return this.handleRequest(req, MASTER_REQUEST)
}

func (this *HttpKernel) HandleSubRequest(req *httpcontext.Request) httpcontext.Responser {
	return this.handleRequest(req, SUB_REQUEST)
}

func (this *HttpKernel) handleRequest(req *httpcontext.Request, reqType int) httpcontext.Responser {
	event := NewResponseEvent(this, req, reqType)
	this.dispatcher.Dispatch(KERNEL_EVENTS_REQUEST, event)

	if event.HasResponse() {
		return this.filterResponse(req, reqType, event.GetResponse())
	}

	controller, method, err := this.resolver.GetController(req)

	if err != nil{
		panic("No controller found")
	}

    // raw controller
    contEvent := NewFilterControllerEvent(this, req, reqType, controller)
    this.dispatcher.Dispatch(KERNEL_EVENTS_CONTROLLER, contEvent)
    controller = contEvent.GetController()

    args := this.resolver.GetArguments(req, controller)
    response, err := this.resolver.Call(method, args)

    if err != nil{
        panic(err.Error())
    }

    responser, ok := response.(httpcontext.Responser)

    if !ok{
        viewEvent := NewResponseForControllerResultEvent(this, req, reqType, responser)
        this.dispatcher.Dispatch(KERNEL_EVENTS_VIEW, viewEvent)

        if event.HasResponse(){
            response = event.GetResponse()
        }

        responser, ok = response.(httpcontext.Responser)

        if !ok{
            panic("Controller didn't return a response")
        }
    }
	return responser
}

func (this *HttpKernel) filterResponse(req *httpcontext.Request, reqType int, res *httpcontext.Response) *httpcontext.Response {
	event := NewFilterResponseEvent(this, req, reqType, res)
	this.dispatcher.Dispatch(KERNEL_EVENTS_RESPONSE, event)
	return event.GetResponse()
}
