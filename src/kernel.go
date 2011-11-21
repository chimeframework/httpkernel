package httpkernel

import (
	dispatcher "chime/components/eventdispatcher"
	"chime/components/httpcontext"
	"net/http"
)

type HttpKerneler interface{
}

type Kerneler interface{
	HttpKerneler
	RegisterBundles() []Bundler
	Boot()
	Shutdown()
	GetBundles() []Bundler
	GetBundle(string) []Bundler
	GetFirstBundle(string) Bundler
	GetName() string
	GetEnviornment() string
	IsDebug() bool
	GetRootDir() string
	GetStartTime() int
	// TODO: implement locateResource(), GetCacheDir(), GetLogDir()
}

/// HttpKernel
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

	if err != nil {
		panic("No controller found")
	}

	// raw controller
	contEvent := NewFilterControllerEvent(this, req, reqType, controller)
	this.dispatcher.Dispatch(KERNEL_EVENTS_CONTROLLER, contEvent)
	controller = contEvent.GetController()

	args := this.resolver.GetArguments(req, controller)
	response, err := this.resolver.Call(method, args)

	if err != nil {
		panic(err.Error())
	}

	responser, ok := response.(httpcontext.Responser)

	if !ok {
		viewEvent := NewResponseForControllerResultEvent(this, req, reqType, responser)
		this.dispatcher.Dispatch(KERNEL_EVENTS_VIEW, viewEvent)

		if event.HasResponse() {
			response = event.GetResponse()
		}

		responser, ok = response.(httpcontext.Responser)

		if !ok {
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
