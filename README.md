# GoToTo ( GTT )

GTT generates TS Files out of commented structs.
First it scans the project looking for commented structs and generate a YAML file containing all structs, contexts,filenames and output Directory path for .ts files.
Then it creates ts files reading this YAML File.

When building in project.

```bash
$ go test
$ go build
$ ./Gototo GenerateYML -spath="my/project/path" -dir="my/output/dir/path"
$ ./Gototo GenerateYML -spath="anotherprojectpath" -dir="./insidemyprojectpath"
$ ./Gototo UpdateYML -spath="project_path" -dir="./anewdir" --no-more-files=true
$ ./Gototo GenerateTS -spath="my/project/path" --ignore-pending=false
$ ./Gototo help
```


# V1.0

### Input 
Look at files in project_example domain. 
Structs are annotated with @tsInterface[context="mycontext"] so they can be generated.
Enums needs to be in an enums.go file and formated as in the example so it can be generated. 
No annotations needed around the enum in go.

### Output

> **File Tree**

:file_folder: My Project <br/>
┣ :clipboard:gototoConf.yaml <br/>
┣ :file_folder: src <br/>
┣ :file_folder: output_dir <br/>
&nbsp;&nbsp;&nbsp;&nbsp;┣:page_facing_up: UserModel.ts <br/>
&nbsp;&nbsp;&nbsp;&nbsp;┣ :page_facing_up: ProductModel.ts <br/>
&nbsp;&nbsp;&nbsp;&nbsp;┣ :page_facing_up: enums.ts <br/>
┣ :file_folder: src <br/>

> **GototoConf.yaml**
```yaml
outputdirpath: ./ModelsTS
contextlist:
  - product
  - newcontext
  - user
pending: []
filenamecontextsmap:
  NewcontextModel:
    - newcontext
  ProductModel:
    - product
  UserModel:
    - user
batchesinterface:
  newcontext:
    - name: Address
      fields:
        - name: id
          ftype: number
          omitempty: false
        - name: street
          ftype: string
          omitempty: false
        - name: h_number
          ftype: string
          omitempty: false
        - name: city
          ftype: string
          omitempty: false
        - name: country_code
          ftype: string
          omitempty: false
  product:
    - name: Item
      fields:
        - name: name
          ftype: string
          omitempty: false
        - name: price
          ftype: number
          omitempty: false
        - name: in_stock
          ftype: boolean
          omitempty: false
  user:
    - name: User
      fields:
        - name: person
          ftype: Person
          omitempty: false
        - name: pwd
          ftype: string
          omitempty: false
        - name: login
          ftype: string
          omitempty: false
    - name: Person
      fields:
        - name: name
          ftype: string
          omitempty: false
        - name: given
          ftype: string
          omitempty: false
        - name: gender
          ftype: string
          omitempty: false
        - name: phone
          ftype: string
          omitempty: true
        - name: email
          ftype: string
          omitempty: false
batchesenums:
  - - name: paymentMethods
      consts:
        CASH: CASH
        CB: CB
        GIFT_CARD: GIFT_CARD
    - name: roles
      consts:
        CHEF_COOK: CHEF_COOK
        GENERAL_MANAGER: GENERAL_MANAGER
        OWNER: OWNER
        SERVER: SERVER
  - - name: productType
      consts:
        CLOTH: CLOTH
        FOOD: FOOD
        TOY: TOY

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

