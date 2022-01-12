package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/phlashdev/sherlock/codeanalysis/compilation"
	"github.com/phlashdev/sherlock/codeanalysis/syntax"
)

type ConsoleColor string

const (
	ColorReset ConsoleColor = "\033[0m"
	ColorGray  ConsoleColor = "\033[90m"
	ColorRed   ConsoleColor = "\033[31m"
)

func main() {
	showTree := false

	for {
		fmt.Print("> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		line := scanner.Text()
		if line == "" {
			return
		}

		if line == "#showTree" {
			showTree = !showTree
			if showTree {
				fmt.Println("Showing parse trees")
			} else {
				fmt.Println("Not showing parse trees")
			}
			continue
		} else if line == "#cls" {
			// todo
			// Clear console
			continue
		}

		syntaxTree := syntax.Parse(line)
		compilation := compilation.NewCompilation(syntaxTree)
		result := compilation.Evaluate()

		diagnostics := result.Diagnostics()

		if showTree {
			fmt.Print(ColorGray)
			prettyPrint(syntaxTree.Root(), "", true)
			fmt.Print(ColorReset)
		}

		if len(diagnostics) == 0 {
			fmt.Println(result.Value())
		} else {
			for _, diagnostic := range diagnostics {
				fmt.Println()

				fmt.Print(ColorRed)
				fmt.Println(diagnostic)

				span := diagnostic.Span()
				prefix := line[0:span.Start()]
				err := line[span.Start():span.Length()]
				suffix := line[span.End():]

				fmt.Print("    ")
				fmt.Print(prefix)

				fmt.Print(ColorRed)
				fmt.Print(err)
				fmt.Print(ColorReset)

				fmt.Print(suffix)

				fmt.Println()
			}
		}
	}
}

func prettyPrint(node syntax.SyntaxNode, indent string, isLast bool) {
	var marker string
	if isLast {
		marker = "└── "
	} else {
		marker = "├── "
	}

	fmt.Printf("%s%s%v", indent, marker, node.Kind())

	if t, ok := node.(*syntax.SyntaxToken); ok && t.Value() != nil {
		fmt.Printf(" %v", t.Value())
	}

	fmt.Println()

	if isLast {
		indent += "    "
	} else {
		indent += "│   "
	}

	children := node.GetChildren()
	childrenCount := len(children)
	for index, child := range children {
		prettyPrint(child, indent, index == childrenCount-1)
	}
}
