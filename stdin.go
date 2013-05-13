//Package stdin is a simple utility to make it even easier to read in lines of text.
//
//Warning: It calls log.Fatal if stdin returns an error for whatever reason.
//
//Requires Go 1.1.
package stdin

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

//Bytes returns all stdin as a slice of bytes.
func Bytes() []byte {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalln("Error reading standard input:", err)
	}
	return bytes
}

//Chan returns a channel to which each line of stdin is sent.
func Chan() <-chan string {
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

//Slice returns each line of stdin in a slice of strings.
func Slice() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalln("Error reading standard input:", err)
	}
	return lines
}
