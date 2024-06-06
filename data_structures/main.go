package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type Person struct {
	Id        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"userId"`
}

func basic() {
	buffer := make([]byte, 1024)
	file, err := os.Create("example2.txt")
	if err != nil {
		fmt.Println(err)
	}

	content := "hello fenil@1234#567"
	byte, err := io.WriteString(file, content)
	fmt.Println("total written bytes are ", byte)
	if err != nil {
		fmt.Println("error in writing", err)
	}
	file, err = os.Open("example2.txt")
	if err != nil {
		fmt.Println("error in opeaing file", err)
		return
	}
	defer file.Close()
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error in reading file", err)
			return
		}

		fmt.Println(string(buffer[:n]))
	}

	fmt.Println("size of buffer", len(buffer))
	fmt.Println("capacity of buffer", cap(buffer))
}
func dealWithUrl() {

	myUrl := "https://dummyjson.com/todos/1"
	newUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("error in parsing url", err)
	}
	fmt.Println("new parsed url is", newUrl)
}
func get() {
	res, err := http.Get("https://dummyjson.com/todos/1")
	if err != nil {
		fmt.Println("error in get product", err)
	}
	fmt.Printf("type of res is %T\n", res)
	data, error := io.ReadAll(res.Body)
	defer res.Body.Close()
	if error != nil {
		fmt.Println("error in reading data", error)
	}
	fmt.Printf("type of data is %T\n", data)
	fmt.Println("data is ", string(data))
}
func main() {
	//basic()

	//get method
	//get()

	//deal with url
	dealWithUrl()

	//dealwithJson

	res, err := http.Get("https://dummyjson.com/todos/1")

	if res.StatusCode != http.StatusOK {
		fmt.Println("error in getting data", err)
	}

	if err != nil {
		fmt.Println("error in get product", err)
	}
	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))

	//convert to json
	todos := Person{
		Id:        1,
		Todo:      " car drive",
		Completed: false,
		UserId:    132,
	}

	jsondata, err := json.Marshal(todos)
	if err != nil {
		fmt.Println("error in marshaling", err)
	}
	fmt.Println("json data is ", string(jsondata))

	//get simple original object from jsonData and store it to persomn1
	var person1 Person
	err = json.Unmarshal(jsondata, &person1)
	if err != nil {
		fmt.Println("error in unmarshaling", err)
	}
	fmt.Println("person1 is", person1)

	// efficient way of
	var person2 Person

	// Decode the JSON data into the person2 struct
	err = json.NewDecoder(res.Body).Decode(&person2)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("person2", person2)
	defer res.Body.Close()
}
