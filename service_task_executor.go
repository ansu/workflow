// package tasks

// import (
// 	"fmt"
// 	"reflect"
// )

// type ServiceTaskExecutor struct{}

// func (st *ServiceTaskExecutor) Execute(methodName string, params interface{}, successScreenID int, errorScreenID int) (int, error) {
// 	// Get the method to execute
// 	fmt.Println("coming")
// 	method := reflect.ValueOf(st).MethodByName(methodName)
// 	fmt.Println("coming1")

// 	// Make sure the method exists
// 	if !method.IsValid() {
// 		return errorScreenID, fmt.Errorf("Service task method %s not found", methodName)
// 	}

// 	// Convert the params to a slice of reflect.Values
// 	var paramValues []reflect.Value
// 	if params != nil {
// 		value := reflect.ValueOf(params)
// 		if value.Kind() != reflect.Slice {
// 			return errorScreenID, fmt.Errorf("Service task params must be a slice of values")
// 		}
// 		for i := 0; i < value.Len(); i++ {
// 			paramValues = append(paramValues, value.Index(i))
// 		}
// 	}

// 	// Call the method with the given params
// 	result := method.Call(paramValues)

// 	// Check for errors
// 	if len(result) < 1 {
// 		return errorScreenID, fmt.Errorf("Service task method %s did not return a value", methodName)
// 	}
// 	err := result[0].Interface()
// 	if err != nil {
// 		return errorScreenID, fmt.Errorf("Service task method %s returned an error: %v", methodName, err)
// 	}

// 	// Check for success
// 	if len(result) < 2 {
// 		return errorScreenID, fmt.Errorf("Service task method %s did not return a success value", methodName)
// 	}
// 	success := result[1].Interface().(bool)
// 	if success {
// 		return successScreenID, nil
// 	}
// 	return errorScreenID, nil
// }

// // Example service task method
// func (st *ServiceTaskExecutor) ExampleServiceTask(param1 string, param2 int) (error, bool) {
// 	// Execute some logic here
// 	// ...
// 	fmt.Println("Coming")
// 	return nil, true
// }

package tasks

import (
	"fmt"
	"reflect"
)

type ServiceTaskExecutor struct{}

func (st *ServiceTaskExecutor) Method2(inputData map[string]interface{}) (interface{}, error) {
	// method logic here
	fmt.Println("Method2")
	var i interface{}
	i = "method2"
	return i, nil

}

func (st *ServiceTaskExecutor) Execute(methodName string, inputData []interface{}) (interface{}, error) {

	method := reflect.ValueOf(st).MethodByName(methodName)
	methodType := method.Type()

	numExpectedParams := methodType.NumIn()
	if !method.IsValid() {
		return nil, fmt.Errorf("invalid method name: %s", methodName)
	}
	numParams := len(inputData)
	in := make([]reflect.Value, numParams)
	for i := 0; i < numParams; i++ {
		in[i] = reflect.ValueOf(inputData[i])
	}
	// var paramValues []reflect.Value
	// if params != nil {
	// 	value := reflect.ValueOf(params)
	// 	if value.Kind() != reflect.Slice {
	// 		return errorScreenID, fmt.Errorf("Service task params must be a slice of values")
	// 	}
	// 	for i := 0; i < value.Len(); i++ {
	// 		paramValues = append(paramValues, value.Index(i))
	// 	}
	// }

	// err := st.ValidateInputParams(inputData)
	// if err != nil {
	// 	return nil, err
	// }

	out := method.Call(in)
	// if err != nil {
	// 	return nil, err
	// }
	// Check if the method returned an error
	if errVal := out[numExpectedParams-1]; !errVal.IsNil() {
		return nil, errVal.Interface().(error)
	}

	// Return the result of the method call
	return out[0].Interface(), nil
}
