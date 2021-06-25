declare namespace internal {
	export interface User {
		Id : number;
		Name : string;
		Given : string;
	}

}


declare namespace external {
	export interface Person {
		Name : string;
		Given : string;
	}

}


export interface Product {
	Price : number;
	Name : string;
}

