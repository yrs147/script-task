package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

		cores,err := strconv.Atoi(strings.TrimSuffix(row[1],"m"))
		handleError(err)

		memory, err := strconv.Atoi(strings.TrimSuffix(row[2],"Mi"))
		handleError(err)


		fmt.Printf("Name : %s\n",row[0])
		fmt.Printf("CPUCores: %dm\n", cores)
		fmt.Printf("Memory: %dMi\n", memory)
		fmt.Println("----------------")
	}

	
	
}
