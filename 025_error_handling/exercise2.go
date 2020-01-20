package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	 if err !=nil{
		 log.Println(err)
		return
	 }
	fmt.Println(string(bs))
}	


func toJSON(v interface{}) ([]byte, error){
	bs, err := json.Marshal(v)
	  if err!=nil{
	    return []byte{},  fmt.Errorf("my error is",err)
    }
	  return bs, nil
	
}