package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("logs.txt")
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	scanner.Scan()

	for scanner.Scan(){
		line := scanner.Text()
		row := strings.Fields(line)

		fmt.Printf("Name : %s",row[0])
	}

	
	// if err = scanner.Err(); handleError(err)
	
	// columns := strings.Fields(line)

	// // data := (string(columns))
	// // test := strings.Split(data," ")
	// fmt.Println(columns[0])

}
