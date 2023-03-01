package tasks

import "fmt"

func (st *ServiceTaskExecutor) Digio(a int, b string) (interface{}, error) {
	// method logic here
	fmt.Println("Digio")
	var i interface{}
	i = "hello"
	return i, nil
}

func (st *ServiceTaskExecutor) Digio1(a int, b string) (interface{}, error) {
	// method logic here
	fmt.Println("Digio1")
	var i interface{}
	i = "hello1"
	return i, nil
}
