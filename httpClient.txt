package main

import (
	"fmt"
	"net/http"
	"os"

	// "strings"
	"bufio"
)

func mainHttp() {
	// // get the data from prv example
	{
		resp, err := http.Get("http://localhost:8090/users")
		pancOnError(err)
		defer resp.Body.Close()

		// get stauts code
		fmt.Println("status code :", resp.StatusCode)

		reader := bufio.NewReader(resp.Body)
		s := bufio.NewScanner(reader)

		for s.Scan() {
			fmt.Println(s.Text())
		}

		err = s.Err()
		pancOnError(err)

	}
	// {
	// 	resp, err := http.Head("http://localhost:8090/users")
	// 	pancOnError(err)
	// 	defer resp.Body.Close()
	// 	// get status code
	// 	fmt.Println("status code :", resp.StatusCode)
	// }
	// {
	// 	// post request
	// 	json := `{"ID":6, "Name":"hana", "Score":20}`
	// 	b := strings.NewReader(json)

	// 	resp, err := http.Post("http://localhost:8090/users", "application/json", b)
	// 	pancOnError(err)
	// 	defer resp.Body.Close()

	// 	// get status code
	// 	fmt.Println("status code :", resp.StatusCode)
	// }
}

func pancOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
