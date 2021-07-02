package main

import (
	"flag"
	"fmt"
	"os"
)

const help string = "\n\nGototo generates TS files out of annotated structs in a project.\n" +
	"Find more information at https://github.com/Rulioos/GoToTo. \n" +
	"Commands : \n\t" +
	"-GenerateYML => Generate the YAML configuration file that can be modified to setup filesname and manage contexts in files." +
	"\n\t\t Usage: Gototo GenerateYML -spath=[project_path] -dir=[path_to_output_dir]" +
	"\n\t" +
	"-UpdateYML => Update an existing YAML file without modifying the previous changes to filenames and context management." +
	"\n\t\t Usage: Gototo UpdateYML -spath=[project_path] -dir=[path_to_new_dir] --no-more-files=true" +
	"\n\t" +
	"-GenerateTS => Generate TS files according to the corresponding YAML file in the project." +
	"\n\t\t Usage: Gototo GenerateTS -spath=[project_path] --ignore-pending=False" +
	"Use Gototo command [--help] for more information on a given command.\n\n"

func main() {
	ParseEnumsGoFile("project_example/domain/enums/enums.go")
	//Generate YAML FLAGS
	GenerateConfigYamlCmd := flag.NewFlagSet("generateYML", flag.ExitOnError) //Generate config file
	outputDir := GenerateConfigYamlCmd.String("dir", "", "Directory in which the .ts files will be created")
	scanPath := GenerateConfigYamlCmd.String("spath", "./", "Path to the project to scan")

	//UPDATE YAML FLAGS
	UpdateConfigYamlCmd := flag.NewFlagSet("updateYML", flag.ExitOnError) //Update config file
	UpdateYamlDir := UpdateConfigYamlCmd.String("dir", "", "Directory in which the .ts files will be created")
	UpdateYamlPath := UpdateConfigYamlCmd.String("spath", "./", "Path to the project to scan")
	NoMoreFiles := UpdateConfigYamlCmd.Bool("no-more-files", true, "Boolean to generate more files or not according to context. True by default. ")

	//Generate TS FLAGS
	GenerateFromYamlCmd := flag.NewFlagSet("generateTS", flag.ExitOnError) //Generate ts files
	projectPath := GenerateFromYamlCmd.String("spath", "./", "Path to the project to scan")
	ignorePending := GenerateFromYamlCmd.Bool("ignore-pending", false, "Ignore if there is pending contexts ie not in a file. Value false by default")

	if len(os.Args) < 2 {
		fmt.Println(help)
		os.Exit(1)
	}
	switch os.Args[1] {
	case "generateYML":
		GenerateConfigYamlCmd.Parse(os.Args[2:])
		fmt.Printf("outputDir: %v\n", *outputDir)
		fmt.Printf("scanPath: %v\n", *scanPath)
		GenerateConfigYaml(*scanPath, *outputDir)
	case "updateYML":
		UpdateConfigYamlCmd.Parse(os.Args[2:])
		fmt.Printf("outputDir: %v\n", *UpdateYamlDir)
		fmt.Printf("scanPath: %v\n", *UpdateYamlPath)
		fmt.Printf("No more files: %v\n", *NoMoreFiles)
		UpdateConfigYaml(*UpdateYamlPath, *UpdateYamlDir, *NoMoreFiles)
	case "generateTS":
		GenerateFromYamlCmd.Parse(os.Args[2:])
		fmt.Printf("projectPath: %v\n", *projectPath)
		fmt.Printf("ignorePending: %v\n", *ignorePending)
		GenerateFromYaml(*projectPath, *ignorePending)

	case "help":
		fmt.Print(help)

	default:
		fmt.Println(help)
		os.Exit(1)
	}

}
