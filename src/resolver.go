package httpkernel

import (
	"chime/components/httpcontext"
	_ "errors"
	_ "fmt"
	"reflect"
	_ "strings"
	// "chime/core/router"
	// "chime/core/httpContext"
	// "chime/core/utils"
)

type Resolver struct {
	ControllerMaps map[string]*reflect.Value
	ActionMaps     map[string]*reflect.Value
}

func NewResolver() *Resolver {
	return &Resolver{
		ControllerMaps: make(map[string]*reflect.Value),
		ActionMaps:     make(map[string]*reflect.Value),
	}
}

func (this *Resolver) GetController(req *httpcontext.Request) (controller *reflect.Value, method *reflect.Value, err error) {
	return nil, nil, nil
}

/*
func (this *ClassResolver) load(bundles map[string]*appBundle) {
    for key, value := range bundles {
        this.discoverStructs(key, value.Type)
    }
}

func (this *ClassResolver) addBundle(bundle *appBundle) {
    this.discoverStructs(bundle.Name, bundle.Type)
}

func (this *ClassResolver) discoverStructs(app string, cm interface{}) {
    cmt := reflect.TypeOf(cm)
    cmt_methodcount := cmt.NumMethod()
    for cmt_methodi := 0; cmt_methodi < cmt_methodcount; cmt_methodi++ {
        cmt_method := cmt.Method(cmt_methodi)
        params := make([]reflect.Value, 1)
        params[0] = reflect.ValueOf(cm)
        result := cmt_method.Func.Call(params)
        cv := result[0]

        name := cmt_method.Name
        if strings.HasSuffix(name, CONTROLLER_SUFFIX) {
            key := fmt.Sprintf("%v%v%v", app, SEPARATOR, name)
            this.ControllerMaps[key] = &cv
            this.discoverMethods(key, &cv)
        }
    }
}

func (this *ClassResolver) discoverMethods(key string, cv *reflect.Value) {
    cmt := cv.Type()
    cmt_methodcount := cmt.NumMethod()

    for cmt_methodi := 0; cmt_methodi < cmt_methodcount; cmt_methodi++ {
        cmt_method := cmt.Method(cmt_methodi)
        name := cmt_method.Name
        if strings.HasSuffix(name, ACTION_SUFFIX) {
            key := fmt.Sprintf("%v%v%v", key, SEPARATOR, name)
            // fmt.Printf("Found Method: %v\n", key)
            this.ActionMaps[key] = &cmt_method.Func
        }
    }
}

func (this *ClassResolver) GetCallable(route *router.Route) (controller *reflect.Value, method *reflect.Value, err error) {
    contKey := fmt.Sprintf("%v%v%v%v", route.Bundle, SEPARATOR, route.Controller, CONTROLLER_SUFFIX)
    //contKey := "frontend_" + route.Controller + "Controller"
    controller, contOk := this.ControllerMaps[contKey]
    if !contOk {
        err = errors.New(fmt.Sprintf("Controller %v not found\n", contKey))
        return nil, nil, err
    }

    methodKey := fmt.Sprintf("%v%v%v%v", contKey, SEPARATOR, route.Action, ACTION_SUFFIX)
    // methodKey := "frontend_" + route.Controller + "Controller_" + route.Action + "Action"
    method, methOk := this.ActionMaps[methodKey]
    if !methOk {
        err = errors.New(fmt.Sprintf("Action %v not found\n", methodKey))
        return controller, nil, err
    }

    return controller, method, nil
}

*/
func (this *Resolver) GetArguments(req *httpcontext.Request, controller *reflect.Value) []reflect.Value {
    params := make([]reflect.Value, 2)
    params[0] = *controller
    params[1] = reflect.ValueOf(req.GetAttributes())
    return params
}

func (this *Resolver) Call(method *reflect.Value, args []reflect.Value) (response interface{}, err error){
    output := method.Call(args)

    if (len(output) > 1) && (!output[1].IsNil()){
        return nil, output[1].Interface().(error)
    }

    return output[0].Interface(), nil
}
