export interface User {
	person : Person;
	pwd : string;
	login : string;
}

export interface Person {
	name : string;
	given : string;
	gender : string;
	phone ?: string;
	email : string;
}

