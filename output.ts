declare namespace internal {
	export interface User {
		id : number;
		name : string;
		given : string;
	}

}


declare namespace external {
	export interface Adress {
		given : string;
	}

	export interface Person {
		name : string;
		given : string;
		adresses : Adress[];
	}

}


export interface Product {
	price : number;
	name : string;
}

