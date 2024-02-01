package main

import (
	"fmt"
	"os"

	//"flag"
	"strings"
)

func mainCmd() {
	// read cmd args
	// {
	// 	args := os.Args
	// 	path := os.Args[0]
	// 	fmt.Printf("path : %v \n", path)
	// 	args = os.Args[1:]

	// 	for i, v := range args {
	// 		fmt.Printf("index:%v, Value:%v \n",i, v )
	// 	}
	// }
	// read cmd flags

	// {
	// 	fname := flag.String("fname", "","first name as string" )
	// 	score := flag.Int("score", 0, "score as int")
	// 	start := flag.Bool("start", false, "start as boolian")

	// 	var lname string
	// 	flag.StringVar(&lname, "lname", "", "last name as string")

	// 	flag.Parse()

	// 	fmt.Printf("first name :%v, Score:%v, last name :%v, start:%v", *fname, *score,lname, *start)

	// 	fmt.Println("run of the args", flag.Args())
	// }

	{
		fmt.Println("GAME_SCORE", os.Getenv("GAME_SCORE"))
		os.Setenv("GAME_SCORE", "100")
		fmt.Println("GAME_SCORE 2", os.Getenv("GAME_SCORE"))

		fmt.Println("all env values")

		for _, e := range os.Environ() {
			envRow := strings.Split(e, "=")
			fmt.Printf("key :%v, value :%v \n", envRow[0], envRow[1])
		}

	}
}
