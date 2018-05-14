package main

import "fmt"

func main() {
	q := "package main\n\nimport \"fmt\"\n\nfunc main() {\n\tq := %q\n\tfmt.Printf(q, q)\n}\n"
	fmt.Printf(q, q)
}
