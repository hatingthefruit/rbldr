/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/hatingthefruit/rbldr/util"
)

func main() {
	allArgs := os.Args[1:]
	template := flag.String("template", "default", "Name of template to use")
	flag.Parse()

	var inputFile, outputFile *os.File
	var inputErr error
	if len(allArgs) == 0 {
		//fmt.Println("pulling from stdin, printing to stdout")
		inputFile = os.Stdin
		outputFile = os.Stdout
	} else if len(allArgs) == 1 {
		//fmt.Printf("Pulling from %s, printing to stdout\n", allArgs[0])
		inputFile, inputErr = os.Open(allArgs[0])
		util.CheckErr(inputErr)
		outputFile = os.Stdout
	} else {
		//fmt.Printf("Infile %s and outfile %s\n", allArgs[0], allArgs[1])

		inputFile, inputErr = os.Open(allArgs[0])
		util.CheckErr(inputErr)
		outputFile, inputErr = os.Open(allArgs[1])
		util.CheckErr(inputErr)

	}

	defer inputFile.Close()
	defer outputFile.Close()

	finalResume := util.NewResume(inputFile)

	//fmt.Println(finalResume)

	rWriter := bufio.NewWriter(outputFile)
	buildResume := util.BuildResume(finalResume, *template)
	_, err := rWriter.WriteString(buildResume)
	util.CheckErr(err)
	rWriter.Flush()
}
