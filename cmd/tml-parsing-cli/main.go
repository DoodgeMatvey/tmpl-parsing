package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	parsetmpl "github.com/DoodgeMatvey/rbac-parsing-tmpl"
	mysql "github.com/DoodgeMatvey/rbac-parsing-tmpl/mysql"

	"gopkg.in/yaml.v2"
)

func main() {
	// Define flags
	inputFilePath := flag.String("input", "./Rest-api.yaml", "Input YAML file path")
	outputFilePath := flag.String("output", "./result.yaml", "Output YAML file path")
	flag.Parse()

	//Read from file
	inputSpec, err := ioutil.ReadFile(*inputFilePath)
	if err != nil {
		panic(err.Error())
	}

	// Define variables
	var openApiSpec parsetmpl.OpenApiSpec
	yaml.Unmarshal(inputSpec, &openApiSpec)

	result := openApiSpec.Parsetmpl()

	// Create open file
	outputFile, err := os.Create(*outputFilePath)
	if err != nil {
		panic(err.Error())
	}

	// Write data into outputFile
	yaml.NewEncoder(outputFile).Encode(result)

	fmt.Println(result)
	// Close file
	err = outputFile.Close()
	if err != nil {
		panic(err.Error())
	}

	// Open mysql db
	db, err := mysql.Open("root:q3l8vtRp890@tcp(127.0.0.1:3306)/parsetmpl_db")
	if err != nil {
		panic(err)
	}

	// Insert data into db
	mysql.InsertData(db, result)

	// Close mysql db
	defer db.Close()
}
