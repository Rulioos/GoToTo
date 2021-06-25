package main

func main() {
	batches := ParseGoFile("test.go")
	GenerateAll_TS_namespace("output", batches)
	GenerateAll_TS_multiple_files("output_dir", batches)
}
