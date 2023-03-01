package main

import (
	"example.com/WORKFLOW/orchestrator"
	"example.com/WORKFLOW/tasks"
)

func main() {

	serviceTask := tasks.ServiceTaskExecutor{}
	orchestrator := orchestrator.NewOrchestrator(serviceTask)
	orchestrator.Dummy()

	// // serviceTask := &tasks.ServiceTaskExecutor{}
	// // inputData := map[string]interface{}{
	// // 	"a": "1",
	// // 	"b": 1,

	// // 	// any other fields required by the specific service task being executed
	// // }

	// inputData := []interface{}{1, "test"}
	// // Example 1: Call the "sum" method with input parameters 2 and 3
	// result, err := serviceTask.Execute("Digio", inputData)
	// if err != nil {
	// 	log.Fatalf("Error calling service task: %s", err.Error())
	// }
	// fmt.Println(result)

	// result1, err1 := serviceTask.Execute("Digio1", inputData)
	// if err1 != nil {
	// 	log.Fatalf("Error calling service task: %s", err1.Error())
	// }
	// fmt.Println(result1)

	// // Example 2: Call the "concat" method with input parameters "Hello" and "World"
	// result, err = serviceTask.Execute("Method2", inputData)
	// if err != nil {
	// 	log.Fatalf("Error calling service task: %s", err.Error())
	// }
	// fmt.Printf("Result of concat operation: %s\n", result.(string))
	// sum := math_util.Add(2, 3)
	// instant := &services.ServiceTask{}

	// instant1 := &services.ExampleServiceTask{}

	// inputData := map[string]interface{}{
	// 	"a": "1",
	// 	"b": 1,

	// 	// any other fields required by the specific service task being executed
	// }
	// inputData := []interface{}{"1", 1}

	// serviceTaskName := "ExampleServiceTask"

	// result, err := instant.Execute(serviceTaskName, inputData, 1, 2)
	// if err != nil {
	// 	// handle error
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(result)

	// fmt.Println("The sum is:", sum)
}
