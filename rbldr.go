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
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/hatingthefruit/rbldr/util"
)

func old_main() {
	templateName := flag.String("template", "templates/default.json", "Name of template to use")
	resumeDir := flag.String("resumedir", "", "Location of resume files")
	ifName := flag.String("if", "", "Input file name; default is STDIN")
	ofName := flag.String("of", "", "Output file name; default is STDOUT")
	flag.Parse()

	var inputFile, outputFile *os.File
	var inputErr error
	if *ifName == "" {
		//fmt.Println("pulling from stdin, printing to stdout")
		inputFile = os.Stdin
	} else {
		inputFile, inputErr = os.Open(*ifName)
		util.CheckErr(inputErr)
	}
	if *ofName == "" {
		outputFile = os.Stdout
	} else {
		outputFile, inputErr = os.Create(*ofName)
		util.CheckErr(inputErr)
	}

	defer inputFile.Close()
	defer outputFile.Close()

	finalResume := util.NewResume(inputFile, *resumeDir)

	//fmt.Println(finalResume)

	rWriter := bufio.NewWriter(outputFile)

	templateFile, err := os.Open(*templateName)
	if err != nil {
		log.Fatal(err)
	}

	var template util.Template
	json.NewDecoder(templateFile).Decode(&template)

	buildResume := util.BuildResume(finalResume, template)
	_, err = rWriter.WriteString(buildResume)
	util.CheckErr(err)
	rWriter.Flush()
}

func main() {
	templateName := flag.String("template", "templates/resume.templ.tex", "Name of template to use")
	resumeDir := flag.String("resumedir", "", "Location of resume files")
	ifName := flag.String("if", "", "Input file name; default is STDIN")
	ofName := flag.String("of", "", "Output file name; default is STDOUT")
	flag.Parse()

	var inputFile, outputFile *os.File
	var inputErr error
	if *ifName == "" {
		//fmt.Println("pulling from stdin, printing to stdout")
		inputFile = os.Stdin
	} else {
		inputFile, inputErr = os.Open(*ifName)
		util.CheckErr(inputErr)
	}
	if *ofName == "" {
		outputFile = os.Stdout
	} else {
		outputFile, inputErr = os.Create(*ofName)
		util.CheckErr(inputErr)
	}

	defer inputFile.Close()
	defer outputFile.Close()

	finalResume := util.NewResume(inputFile, *resumeDir)

	//fmt.Println(finalResume)

	rWriter := bufio.NewWriter(outputFile)

	templateString, err := os.ReadFile(*templateName)
	if err != nil {
		log.Fatal(err)
	}

	err = util.FillTemplate(finalResume, string(templateString), rWriter)
	util.CheckErr(err)
	rWriter.Flush()
}
