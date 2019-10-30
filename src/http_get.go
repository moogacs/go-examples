package main

import (

	"fmt"
	
	//package implements encoding and decoding of JSON. Mapping between JSON and Go values are done by Marshal and Unmarshal functions provided by json package
	"encoding/json"

	//package ioutil implements some I/O utility functions.
	"io/ioutil"

	//package http provides HTTP client and server implementations.
	"net/http"
)

func main() {

	// mapping json keys to the struct fields. Make sure first character of each field value is uppercase, otherwise mapping will not work
	type User struct {
		Id        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	}

	// mapping json keys to the struct fields. Make sure first character of each field value is uppercase, otherwise mapping will not work. We are taking this struct as json is nested object where "data" is the key for inner object.
	type Data struct {
		User User `json:"data"`
	}

	var data Data

	/*response from dummy API: {
		"data":{
			"id":7,
			"email":"michael.lawson@reqres.in",
			"first_name":"Michael",
			"last_name":"Lawson",
			"avatar":"https://s3.amazonaws.com/uifaces/faces/twitter/follettkyle/128.jpg"
		}
	}
	*/
	
	response, err := http.Get("http://reqres.in/api/users/7")
	if err != nil {
		fmt.Printf("Request Failed %s\n", err) // print error if err is not empty
	} else {
		jsonData, _ := ioutil.ReadAll(response.Body) // json.Decode can also be used to do this conversion. Declared blank identifier("_") to avoid unused variables from the returns values.
		fmt.Printf("%s\n", jsonData)                 // print response as json string
		json.Unmarshal(jsonData, &data)              // there are many ways to do this
		fmt.Printf("%+v", data.User)                 // print struct values with field name
	}
}
