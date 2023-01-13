package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}
func splittoArray(line string) []byte {
	tab := strings.Fields(line)
	println("--- %d", len(tab))
	//toConvert := make([]byte, len(tab))

	return nil
}
func bytetostring(mybte []byte) string {
	return string(mybte)
}
func bytetostring2(mybtes [][]byte) []string {
	var s []string
	for _, mybte := range mybtes {
		s = append(s, bytetostring(mybte))
	}

	return s

}
func stringTobytes(line string) []byte {
	return []byte(line)
}
func stringTobytes2(tab []string) [][]byte {
	var s [][]byte
	for _, mybte := range tab {
		s = append(s, stringTobytes(mybte))
	}

	return s
}
func main() {
	myBytes := [][]byte{{65, 66, 67, 226, 130, 172},
		{65, 66, 67, 226, 130, 172, 12, 34, 21, 23}}
	mywriter := bytetostring2(myBytes)

	if err := writeLines(mywriter, "foo.in.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
	lines, err := readLines("foo.in.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	newBytes := stringTobytes2(lines)
	fmt.Printf("-----------> after cnversion : %v", newBytes)
	/*for i, line := range lines {
		fmt.Println(i, line)
		splittoArray(line)
	}

	*/

}
