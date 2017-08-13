package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func outputCert(outPut error) error {
	f, err := os.OpenFile(*resultsDir+"/results.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()

	_, err = f.WriteString(outPut.Error() + "\n")
	check(err)
	return nil
}

func outputProblemCert(outPut string, problem string) {
	f, err := os.OpenFile(*resultsDir+"/results.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	check(err)
	defer f.Close()
	// First remove the commas from the problem variable, otherwise new columns get made for each comma
	problem = strings.Replace(problem, ",", "", -1)
	_, err = f.WriteString(outPut + ",,,," + problem + "\n")
	check(err)
	return
}

func createOutPutFile() {
	// write output file
	f, err := os.Create(*resultsDir + "/results.csv")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	// String to put into file
	_, err = fmt.Fprintf(w, "%s", columnNames+"\n")
	check(err)
	w.Flush()
	f.Close()
}

func refreshResults() {
	err := os.Remove(*resultsDir + "/results.csv")
	createOutPutFile()
	processHosts()
	check(err)
}

func addHost(s string) {
	f, err := os.OpenFile(*hostsFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("\n" + s); err != nil {
		panic(err)
	}
}

func removeHost(s string) {
	input, err := ioutil.ReadFile(*hostsFile)
	check(err)

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, s) && s != "" {
			lines[i] = ""
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(*hostsFile, []byte(output), 0644)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
