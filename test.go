package main

import "fmt"
import "time"
//import "encoding/json"

const (
	NUM_SENDERS = 5
)

type Response struct {
	stream []byte
	idx int
}

type Location struct {
	loc interface{}
}

func (l Location) pushToLoc() (res int){
	return 1
}

func generateResponseMessages(amount int) ([]Response){
	
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var resps []Response

	//create Response structs
	for i := 0; i < amount; i++ {
		r := Response{byt, i}
		resps = append(resps, r)
	}

	return resps
}

/*
func mapToArray(listptr *Response, idx int, newresp Response) {
	listptr[idx] = newresp
}
*/

func parallelResponseCopy (numNewResponses int) {
	//"parallel" copy to toOut array
	for i := 0; i < numNewResponses; i++ {
		// move goroutine to inlined func? 
		go func (listptr []Response, idx int, newresp Response) {
			listptr[idx] = newresp
		}(toOut, i, responses[i])
	}
}



func main() {	

	NUM_RESPONSES := 10000000
	toOut := make([]Response, NUM_RESPONSES)

	//create test Responses 
	responses := generateResponseMessages(NUM_RESPONSES)

	start := time.Now()
	//"parallel" copy to toOut array
	for i := 0; i < NUM_RESPONSES; i++ {
		go func (listptr []Response, idx int, newresp Response) {
			listptr[idx] = newresp
		}(toOut, i, responses[i])
	}

	//test request/func call to external source
	testChan := make(chan int) 
	for i := 0; i < NUM_RESPONSES; i++ {
		loc := Location{&toOut}
		go func (){
			testChan <- loc.pushToLoc()
		}()
	}	

	// recieve channel updates 
	gone := make([]int, NUM_RESPONSES)
	for i := 0; i < NUM_RESPONSES; i++ {
		go func (idx int){
			gone[idx] = <-testChan
		}(i)
	}

	elapsed := time.Since(start)
	fmt.Println(NUM_RESPONSES, "NUM_RESPONSES took: ", elapsed)
}