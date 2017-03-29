package main

import (
	"bufio"
	"fmt"
	"os"
)

func outPutFile(outPut error) error {
	f, err := os.OpenFile("results/results.csv", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(outPut.Error() + "\n")
	if err != nil {
		return err
	}
	return nil

}

func createOutPutFile() {
	// write output file
	f, err := os.Create("results/results.csv")
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
	err := os.Remove("results/results.csv")
	createOutPutFile()
	processHosts()
	if err != nil {
		fmt.Println(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
