package main

//Mappings is used to map go types to ts types
var mappings = map[string]string{
	"int":     "number",
	"int8":    "number",
	"int32":   "number",
	"int64":   "number",
	"uint":    "number",
	"uint8":   "number",
	"uint32":  "number",
	"uint64":  "number",
	"uintptr": "number",
	"bool":    "boolean",
}

//MapToTs maps go type to ts type
func MapToTs(field string) string {
	if translation, ok := mappings[field]; ok {
		return translation
	}
	return field
}
