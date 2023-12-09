package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit := flag.Int("limit",30,"time limit for your quiz")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*limit)*time.Second)

	correct := 0
	problemLoop:
		for i, p := range problems {
			fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
			answerCh := make(chan string)
			go func(){
				var answer string
				fmt.Scanf("%s\n", &answer)
				answerCh <- answer
				
			}()
			select{
			case <-timer.C:
				fmt.Println()
				break problemLoop
			case answer:= <- answerCh:
				if answer == p.a {
					correct++
				}
			}
		}

		fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}