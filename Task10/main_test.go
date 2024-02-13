package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	files, _ := func() ([]fs.FileInfo, error) {
		f, err := os.Open("tests")
		if err != nil {
			return nil, err
		}
		list, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, err
		}
		sort.Slice(list, func(i, j int) bool {
			numI, _ := strconv.Atoi(strings.TrimSuffix(list[i].Name(), ".a"))
			numJ, _ := strconv.Atoi(strings.TrimSuffix(list[j].Name(), ".a"))
			return numI < numJ
		})
		return list, nil
	}()
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".a") {
			testNum := strings.TrimSuffix(file.Name(), ".a")
			t.Run(fmt.Sprintf("test%s", testNum), func(t *testing.T) {
				inputFile, _ := os.Open(fmt.Sprintf("tests/%s", testNum))
				rightAnswers, _ := os.ReadFile(fmt.Sprintf("tests/%s.a", testNum))
				var output bytes.Buffer
				run(inputFile, &output)
				if strings.TrimSpace(output.String()) != strings.TrimSpace(string(rightAnswers)) {
					t.Errorf("Test %s: FAIL\n", testNum)
					outputLines := strings.Split(strings.TrimSpace(output.String()), "\n")
					rightLines := strings.Split(strings.TrimSpace(string(rightAnswers)), "\n")
					for lineNum := 0; lineNum < len(outputLines) && lineNum < len(rightLines); lineNum++ {
						if outputLines[lineNum] != rightLines[lineNum] {
							t.Errorf("Line %d: got %s, but expected %s\n", lineNum+1, outputLines[lineNum], rightLines[lineNum])
							return
						}
					}
				} else {
					t.Logf("Test %s: OK\n", testNum)
				}
			})
		}
	}
}
