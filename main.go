package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	interval := flag.Int("timer", 30, "Timer")
	flag.Parse()

	fmt.Println(*interval)
	file, err := os.OpenFile("practice.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
	}

	newRecords := make(map[string]int)
	for _, record := range records {
		x, _ := strconv.Atoi(record[1])
		newRecords[record[0]] = x
	}

	timer := time.NewTimer(time.Duration(*interval) * time.Second)
	correct, wrong := 0, 0
	
	for key, value := range newRecords {
		select {
		case <-timer.C:
			fmt.Println("TOO LONG")
			return
		default:
			fmt.Printf("What is the answer of: %s ", key)
			var answer int
			fmt.Scanln(&answer)
			if answer == value {
				correct++
			} else {
				wrong++
			}
		}
	}
	fmt.Printf("Correct: %d, Wrong: %d", correct, wrong)
}
