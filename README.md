# GoToTo ( GTT )
v0.1.0
GTT generates TS Files out of commented structs.
First it scans the project looking for commented structs and generate a YAML file containing all structs, contexts,filenames and output Directory path for .ts files.
Then it creates ts files reading this YAML File.

When building in project.

```bash
$ go test
$ go build
$ ./Gototo generateYML -spath="my/project/path" -"my/output/dir/path"
$ ./Gototo generateYML -spath="anotherprojectpath" -"./insidemyprojectpath"
$ ./Gototo generateTS -spath="my/project/path"
```
YAML file and TS files are rewritten each time corresponding command is launched.
Still implementing tests.
