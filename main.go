package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	GenerateConfigYamlCmd := flag.NewFlagSet("generateYML", flag.ExitOnError) //Generate config file
	outputDir := GenerateConfigYamlCmd.String("dir", "./output", "Directory in which the .ts files will be created")
	scanPath := GenerateConfigYamlCmd.String("spath", "./", "Path to the project to scan")

	GenerateFromYamlCmd := flag.NewFlagSet("generateTS", flag.ExitOnError) //Generate ts files
	projectPath := GenerateFromYamlCmd.String("spath", "./", "Path to the project to scan")

	if len(os.Args) < 1 {
		fmt.Println("expected 'generateYML' or 'generateTS' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generateYML":
		GenerateConfigYamlCmd.Parse(os.Args[2:])
		fmt.Printf("outputDir: %v\n", *outputDir)
		fmt.Printf("scanPath: %v\n", *scanPath)
		GenerateConfigYaml(*scanPath, *outputDir)
	case "generateTS":
		GenerateFromYamlCmd.Parse(os.Args[2:])
		fmt.Printf("projectPath: %v\n", *projectPath)
		GenerateFromYaml(*projectPath)
	default:
		fmt.Println("expected 'generateYML' or 'generateTS' subcommands")
		os.Exit(1)
	}

}
