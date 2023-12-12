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

var(
	currentCore 	int
	currentMemory	int
)

func main() {
	file, err := os.Open("logs.txt")
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	targetName := "cms-api-server-55484f97b-xtb2m"
	// thresholdMemory := 203
	// threshHoldCores := 70
	scanner.Scan()


	for{
		for scanner.Scan(){
			line := scanner.Text()
			row := strings.Fields(line)
	
			if(row[0]==targetName){
				currentCore,err = strconv.Atoi(strings.TrimSuffix(row[1],"m"))
				handleError(err)
		
				currentMemory, err = strconv.Atoi(strings.TrimSuffix(row[2],"Mi"))
				handleError(err)
	
			}else{
				break
			}		
			
		}
		fmt.Printf("Name : %s\n",targetName)
				fmt.Printf("CPUCores: %dm\n", currentCore)
				fmt.Printf("Memory: %dMi\n", currentMemory)
				fmt.Println("----------------")
	}
	
}
