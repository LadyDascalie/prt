package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/atotto/clipboard"
	"bytes"
	"github.com/fatih/color"
)

/*-------------------------+
| This little haiku,       |
| Has been generated with, |
| This commenting tool.    |
+-------------------------*/
func main() {
	var lines []string
	var longest int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		lines = append(lines, txt)

		// count the line's length
		count := utf8.RuneCountInString(txt)
		// set the longest line length
		if count > longest {
			longest = count
		}
	}

	if len(lines) > 0 {
		// create a buffer
		var buf bytes.Buffer

		// create a decorator
		decorator := strings.Repeat("-", longest+1)

		// add the top decorators
		buf.WriteString(fmt.Sprintf("/*%s+", decorator))

		for k, line := range lines {
			line = stripComments(line)

			// first pad the line to the longest length (left-justified)
			// -[1]* means pad from the left (-) with the value held in longest ([1]*)
			line = fmt.Sprintf("%-[1]*s", longest, line)

			// then surround with decorations
			line = fmt.Sprintf("| %s |", line)

			// this is the first iteration
			// prepend a newline
			if k == 0 {
				line = "\n" + line
			}

			// append a newline
			line += "\n"

			// buffer it
			buf.WriteString(line)
		}
		// add the bottom decorators
		buf.WriteString(fmt.Sprintf("+%s*/", decorator))

		// write
		out := buf.String()

		fmt.Println(out)
		if !clipboard.Unsupported {
			clipboard.WriteAll(out)
			color.HiGreen("contents copied to the clipboard")
		}
		return
	}
	fmt.Println("Please provide a comment to prettify")
}

// stripComments does it's best to remove any
// statements beginning or closing a comment from the string.
// this is to prevent a user accidentally generating a broken comment.
func stripComments(line string) string {
	var comments = []string{"/*", "*/"}

	for _, statement := range comments {
		for {
			idx := strings.Index(line, statement)
			if idx == -1 {
				break
			}
			line = line[:idx] + line[idx+2:]
		}
	}
	return line
}
