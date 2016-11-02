package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dur, err := time.ParseDuration(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(strconv.FormatFloat(dur.Seconds(), 'f', -1, 64))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
