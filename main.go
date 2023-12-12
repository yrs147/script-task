package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	currentCoreValue   int
	currentMemoryValue int
	targetPodName      = "cms-api-server-55484f97b-xtb2m" // Set the Name of the Pod to monitor here
	thresholdMemory    = 203                              // set the threshhols memory value here
	thresholdCores     = 70                               // srt the threshold core value here
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	for {
		file, err := os.Open("logs.txt")
		handleError(err)

		scanner := bufio.NewScanner(file)
		scanner.Scan()
		for scanner.Scan() {
			line := scanner.Text()
			row := strings.Fields(line)

			if row[0] == targetPodName {
				currentCoreValue, err = strconv.Atoi(strings.TrimSuffix(row[1], "m"))
				handleError(err)

				currentMemoryValue, err = strconv.Atoi(strings.TrimSuffix(row[2], "Mi"))
				handleError(err)

			}
		}
		file.Close()

		if currentCoreValue > thresholdCores {
			fmt.Println("Alert: CPU Cores exceeded threshold!")
		}

		if currentMemoryValue > thresholdMemory {
			fmt.Println("Alert: Memory exceeded threshold!")
		}

		time.Sleep(5 * time.Second) // I assumed the script would scan the pods ever 5 sec
	}

}
