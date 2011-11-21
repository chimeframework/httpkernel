package httpkernel

import "reflect"


type Bundler interface{
    Boot()
    Shutdown()
    GetParent() string
    GetName() string
    GetPath() string
    GetCallable(string, string) *Callable
    // TODO: implement build()
}

type Callable struct {
    Controller *reflect.Value
    Method *reflect.Value
}

func NewCallable(controller *reflect.Value, method *reflect.Value) *Callable {
    return &Callable{ Controller: controller, Method:method }
}
