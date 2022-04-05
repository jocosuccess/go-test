package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Node struct {
	name        string
	description string
	forksCount  int
}

type Projects struct {
	Nodes []Node `json:"nodes"`
}

type Data struct {
	Projects Projects `json:"projects"`
}

type Result struct {
	Data Data `json:"data"`
}

func main() {
	os.Setenv("HOST", "https://gitlab.com/api/graphql")
	os.Setenv("TOKEN", "glpat-Sg6zNuZJyttvjcB72ZJc")

	var return_data Result

	jsonData := map[string]string{
		"query": `
            query last_projects($n: Int = 5) {
                projects(last:$n) {
                    nodes {
                        name
                        description
                        forksCount
                    }
                }
            }
        `,
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, err := http.NewRequest("POST", os.Getenv("HOST"), bytes.NewBuffer(jsonValue))
	request.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Bearer " + os.Getenv("TOKEN")},
	}
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	defer response.Body.Close()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	// fmt.Println(string(data))
	err1 := json.Unmarshal(data, &return_data)
	if err1 != nil {

		// if error is not nil
		// print error
		fmt.Println(err1)
	}
	// fmt.Println(return_data.Data)
	total := 0
	names := ""
	for i, p := range return_data.Data.Projects.Nodes {
		names += p.name
		fmt.Println("Node", (i + 1), p.name, p.description)
		total += p.forksCount
		names += ","
	}
	fmt.Println("Names", names)
	fmt.Println("Total", total)
}
