package main

var mappings = map[string]string{
	"int":  "number",
	"uint": "number",
}

func MapToTs(field string) string {
	if translation, ok := mappings[field]; ok {
		return translation
	}
	return field
}
