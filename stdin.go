//stdin is a simple utility to make it even easier to read in lines of text.
//
//Requires Go 1.1.

package stdin

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func Bytes() []byte {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("Error reading standard input:", err)
	}
	return bytes
}

func LineSlice() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading standard input:", err)
	}
	return lines
}

func LineChan() chan string {
	scanner := bufio.NewScanner(os.Stdin)
	ch := make(chan string)
	go func() {
		defer close(ch)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatalln("Error reading standard input:", err)
		}
	}()

	return ch
}
