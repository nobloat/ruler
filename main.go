package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type config map[string][]string

func readConfig() (*config, error) {
	file, err := os.Open(".ruler.cfg")
	if err != nil {
		return nil, errors.New("Config file .ruler.cfg not found")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := config{}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && []byte(line)[0] != '#' {
			parts := strings.Split(line, ":")
			if len(parts) != 2 {
				return nil, errors.New("Invalid syantax key:val1,val2 in line " + strconv.Itoa(i))
			}
			key := parts[0]
			values := strings.Split(parts[1], ",")
			for _, v := range values {
				result[key] = append(result[key], strings.TrimSpace(v))
			}

		}
		i++
	}
	result["cfg"] = append(result["cfg"], ".ruler.cfg", ".gitignore")
	return &result, nil
}

func skip(path string) bool {
	if path == ".git" || path == ".svn" || path == ".idea" {
		return true
	}
	return false
}

func match(pattern, name string) bool {
	name = strings.Replace(name, "./", "/", 1)
	if strings.HasPrefix(pattern, "*") {
		return strings.HasSuffix(name, pattern[1:])
	} else if strings.HasSuffix(pattern, "*") {
		return strings.HasPrefix(name, pattern[:len(pattern)-1])
	}
	return pattern == name
}

var Order = [...]string{"test", "example", "doc", "func", "cfg", "res"}

func categorizeFile(path string, c config) string {
	for _, key := range Order {
		if patterns, ok := c[key]; ok {
			for _, pattern := range patterns {
				if match(pattern, path) {
					return key
				}
			}
		}
	}
	return "other"
}

func countLines(path string) (int, error) {
	file, _ := os.Open(path)
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}
	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		//TODO: Proper binary file detection
		if count > 1024 {
			return 0, nil
		}
		switch {
		case err == io.EOF:
			return count + 1, nil
		case err != nil:
			return count, err
		}
	}
}

func main() {
	args := os.Args
	verbose := false
	if len(args) > 1 {
		if args[1] == "--verbose" {
			verbose = true
		}
	}

	c, err := readConfig()
	if err != nil {
		panic(err)
	}

	lines := map[string]int{}
	bytes := map[string]int64{}

	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if info.IsDir() && skip(path) {
				if verbose {
					log.Print("Ignoring " + path)
				}
				return filepath.SkipDir
			} else if !info.IsDir() {
				category := categorizeFile(path, *c)
				l, _ := countLines(path)
				lines[category] += l
				bytes[category] += info.Size()

				if verbose {
					fmt.Println("[" + category + "] " + path + "\t" + strconv.Itoa(l) + " lines\t" + strconv.Itoa(int(info.Size())) + " bytes")
				}
			}
			return nil
		})

	if err != nil {
		panic(err)
	}

	printReport(lines, bytes)
}

func printReport(lines map[string]int, bytes map[string]int64) {
	fmt.Println("RULER-LINES-DOC: " + strconv.Itoa(lines["doc"]))
	fmt.Println("RULER-BYTES-DOC: " + strconv.Itoa(int(bytes["doc"])))
	fmt.Println("RULER-LINES-FUNC: " + strconv.Itoa(lines["func"]))
	fmt.Println("RULER-BYTES-FUNC: " + strconv.Itoa(int(bytes["func"])))
	fmt.Println("RULER-LINES-CFG: " + strconv.Itoa(lines["cfg"]))
	fmt.Println("RULER-BYTES-CFG: " + strconv.Itoa(int(bytes["cfg"])))
	fmt.Println("RULER-LINES-TEST: " + strconv.Itoa(lines["test"]))
	fmt.Println("RULER-BYTES-TEST: " + strconv.Itoa(int(bytes["test"])))
	fmt.Println("RULER-LINES-EXAMPLE: " + strconv.Itoa(lines["example"]))
	fmt.Println("RULER-BYTES-EXAMPLE: " + strconv.Itoa(int(bytes["example"])))
	fmt.Println("RULER-LINES-OTHER: " + strconv.Itoa(lines["other"]))
	fmt.Println("RULER-BYTES-OTHER: " + strconv.Itoa(int(bytes["other"])))
}
