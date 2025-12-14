//package main

/*
import (
	"fmt"
	"log"
)

type Json struct {
	rawData string
	data    string
}

func (j *Json) Decode() error {

	if len(j.rawData) < 3 {

		return fmt.Errorf("string is to short")
	}

	j.data = j.rawData
	return nil

}

func JsonHandler(d string) *Json {
	object := Json{
		rawData: d,
		data:    "",
	}

	return &object

}

func main() {

	j := JsonHandler("dd")

	err := j.Decode()

	if err != nil {
		log.Println("string is to short")
	}

	var things []string
	things = make([]string, 3)
	things[0] = "table"
	things[1] = "chair"

	for _, thing := range things {
		fmt.Println(thing)
	}
	for true {
		fmt.Println("hallo")

	}

}
*/
