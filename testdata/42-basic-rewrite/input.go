package main

import "github.com/matthewmueller/joy/js"

func main() {
	test("hi")
}

func test(msg string) {
	js.Rewrite("console.log($1)", msg)
}
