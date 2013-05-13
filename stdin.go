package stdin

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func All() string {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("Error reading standard input:", err)
	}
	return string(bytes)
}

func AllLines() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading standard input:", err)
	}
	return lines
}

func Lines() chan string {
	scanner := bufio.NewScanner(os.Stdin)
	ch := make(chan string)
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatalln("Error reading standard input:", err)
		}
	}()

	return ch
}
