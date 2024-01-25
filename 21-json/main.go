package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {

	fmt.Println("Hello, playground")

	encodeJson()
	decodeJson()
}

func encodeJson() {

	lco := []Course{
		{
			Name:     "Go",
			Price:    0,
			Platform: "Youtube",
			Tags:     []string{"Go", "Programming Language"},
		},
		{
			Name:     "Python",
			Price:    0,
			Platform: "Youtube",
			Tags:     []string{"Python", "Programming Language"},
		},
		{
			Name:     "Java",
			Price:    0,
			Platform: "Youtube",
			Tags:     nil,
		},
	}

	jsonData, err := json.MarshalIndent(lco, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))

}

func decodeJson() {

	jsonData := []byte(`
	{
		"coursename": "Go",
		"price": 0,
		"Platform": "Youtube",
		"tags": [
			"Go",
			"Programming Language"
		]
	}
	`)

	var lco Course

	checkValid := json.Valid(jsonData)

	if checkValid {

		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &lco)
		fmt.Printf("%#v \n", lco)
	} else {
		fmt.Println("JSON is not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)
	fmt.Printf("%#v \n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Println(k, v)
	}

}
