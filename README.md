# GoToTo ( GTT )

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


# V0.1.0

### Input
```go
//@tsInterface[context="user"]
type User struct {
	Id     uint    `json:"id"`
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Phones []phone `json:"phones,omitempty"`
}

//@tsInterface[context="user"]
type Phone struct {
	Id          uint   `json:"id"`
	Brand       string `json:"brand"`
	PhoneNumber string `json:"phone_number"`
}

//@tsInterface[context="product"]
type Product struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	InStock bool   `json:"in_stock"`
}
```
### Output

> **File Tree**

:file_folder: My Project
┣ :clipboard:gototoConf.yaml <br/>
┣ :file_folder: src <br/>
┣ :file_folder: output_dir <br/>
&nbsp;&nbsp;&nbsp;&nbsp;┣:page_facing_up: UserModel.ts <br/>
&nbsp;&nbsp;&nbsp;&nbsp;┣ :page_facing_up: ProductModel.ts <br/>
┣ :file_folder: src <br/>

> **GototoConf.yaml**
```yaml
outputdirpath: ./modelTS
contextlist:
  - product
  - user
filenamecontextsmap:
  ProductModel:
    - product
  UserModel:
    - user
batches:
  product:
    - name: Item
      fields:
        - name: name
          ftype: string
          omitempty: false
        - name: id
          ftype: number
          omitempty: false
        - name: in_stock
          ftype: boolean
          omitempty: false
  user:
    - name: User
      fields:
        - name: id
          ftype: number
          omitempty: false
        - name: name
          ftype: string
          omitempty: false
        - name: age
          ftype: number
          omitempty: false
        - name: phones
          ftype: []Phone
          omitempty: true
    - name: Phone
      fields:
        - name: id
          ftype: number
          omitempty: false
        - name: brand
          ftype: string
          omitempty: false
        - name: phone_number
          ftype: string
          omitempty: false     
```


> **Model.ts**

```ts
export interface User {
	id: number;
	name : string;
	age : number;
	phones ?: []Phone
}

export interface Phone {
	id: number;
	brand : string;
	phone_number: string;

}
```

> **ProductModel.ts**

```ts
export interface Product {
	id: number;
	name : string;
	in_stock: boolean;
}
```
> **Note:** YAML file and TS files are rewritten each time corresponding command is launched.
Still implementing tests.

