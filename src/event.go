package httpkernel

import (
	dispatcher "chime/components/eventdispatcher"
	"chime/components/httpcontext"
	"reflect"
)

/** START KERNEL EVENT **/

type KernelEvent struct {
	*dispatcher.Event
	kernel      *HttpKernel
	request     *httpcontext.Request
	requestType int
}

func NewKernelEvent(ker *HttpKernel, req *httpcontext.Request, reqType int) *KernelEvent {
	return &KernelEvent{kernel: ker, request: req, requestType: reqType}
}

func (this *KernelEvent) GetKernel() *HttpKernel {
	return this.kernel
}

func (this *KernelEvent) GetRequest() *httpcontext.Request {
	return this.request
}

func (this *KernelEvent) GetRequestType() int {
	return this.requestType
}

/** END KERNEL EVENT **/

/** START RESPONSE EVENT **/

type ResponseEvent struct {
	*KernelEvent
	response *httpcontext.Response
}

func NewResponseEvent(ker *HttpKernel, req *httpcontext.Request, reqType int) *ResponseEvent {
	this := &ResponseEvent{}
	this.kernel = ker
	this.request = req
	this.requestType = reqType
	return this
}

func (this *ResponseEvent) SetResponse(response *httpcontext.Response) {
	this.response = response
	this.StopPropagation()
}

func (this *ResponseEvent) GetResponse() *httpcontext.Response {
	return this.response
}

func (this *ResponseEvent) HasResponse() bool {
	return this.response != nil
}

/** END RESPONSE EVENT **/

/** START FILTER RESPONSE EVENT **/

type FilterResponseEvent struct {
	*KernelEvent
	response *httpcontext.Response
}

func NewFilterResponseEvent(ker *HttpKernel, req *httpcontext.Request, reqType int, res *httpcontext.Response) *FilterResponseEvent {
	this := &FilterResponseEvent{}
	this.kernel = ker
	this.request = req
	this.requestType = reqType
	this.SetResponse(res)
	return this
}

func (this *FilterResponseEvent) SetResponse(response *httpcontext.Response) {
	this.response = response
}

func (this *FilterResponseEvent) GetResponse() *httpcontext.Response {
	return this.response
}

/** END FILTER RESPONSE EVENT **/

/** START FILTER CONTROLLER EVENT **/

// TODO: replace controller with a real controller
type FilterControllerEvent struct {
	*KernelEvent
	controller *reflect.Value
}

func NewFilterControllerEvent(ker *HttpKernel, req *httpcontext.Request, reqType int, controller *reflect.Value) *FilterControllerEvent {
	this := &FilterControllerEvent{}
	this.kernel = ker
	this.request = req
	this.requestType = reqType
	this.SetController(controller)
	return this
}

func (this *FilterControllerEvent) SetController(controller *reflect.Value) {
	// TODO: check for a valid controller
	this.controller = controller
}

func (this *FilterControllerEvent) GetController() *reflect.Value {
	return this.controller
}
/** END FILTER CONTROLLER EVENT **/

/** START RESPONSE FOR CONTROLLER RESULT EVENT **/

type ResponseForControllerResultEvent struct {
	*ResponseEvent
	controllerResult interface{}
}

func NewResponseForControllerResultEvent(ker *HttpKernel, req *httpcontext.Request, reqType int, controllerResult interface{}) *ResponseForControllerResultEvent {
	this := &ResponseForControllerResultEvent{}
	this.kernel = ker
	this.request = req
	this.requestType = reqType
	this.controllerResult = controllerResult
	return this
}

/** END RESPONSE FOR CONTROLLER RESULT EVENT **/
