package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type config map[string][]string

func readConfig() (*config, error) {
	file, err := os.Open(".ruler.cfg") // For read access.
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

func categorizeFile(path string, c config) string {
	order := []string{"test", "example", "doc", "source", "config"}
	for _, key := range order {
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
	c, err := readConfig()
	if err != nil {
		panic(err)
	}

	lines := map[string]int64{}
	bytes := map[string]int64{}

	err = filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				panic(err)
			}
			if info.IsDir() && skip(path) {
				return filepath.SkipDir
			} else if !info.IsDir() {
				category := categorizeFile(path, *c)
				l, _ := countLines(path)
				lines[category] += int64(l)
				bytes[category] += info.Size()
			}
			return nil
		})

	if err != nil {
		panic(err)
	}

	fmt.Println(lines)
	fmt.Println(bytes)
}
