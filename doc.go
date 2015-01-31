/*
Package stringcase makes it so you can convert strings to different casing styles:

* lower case

* UPPER CASE

* Title Case

* camelCase

* PascalCase

* snake_case

* CONST_CASE

* property-case

* Header-Case

Samples
	s := stringcase.ToLowerCase("Hello world")
	// hello world
.
	s := stringcase.ToUpperCase("Hello world")
	// HELLO WORLD
.
	s := stringcase.ToTitleCase("Hello world")
	// Hello World
,
	s := stringcase.ToCamelCase("Hello world")
	// helloWorld
.
	s := stringcase.ToPascalCase("Hello world")
	// HelloWorld
.
	s := stringcase.ToSnakeCase("Hello world")
	// hello_world
.
	s := stringcase.ToConstCase("Hello world")
	// HELLO_WORLD
.
	s := stringcase.ToPropertyCase("Hello world")
	// hello-world
.
	s := stringcase.ToHeaderCase("Hello world")
	// Hello-World

*/
package stringcase
