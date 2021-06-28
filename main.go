package main

func main() {
	basepath := "./test"
	outputDir := "./output"
	GenerateConfigYaml(basepath, outputDir)
	GenerateFromYaml(basepath)

}
