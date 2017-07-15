package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/atotto/clipboard"
)

/*---------------------------+
| Comment generated via prt! |
+---------------------------*/
func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		str := strings.Join(args, " ")
		str = fmt.Sprintf("| %s |", str)

		// -3 to offset for the added spaces
		strlen := utf8.RuneCountInString(str) - 3
		line := strings.Repeat("-", strlen)

		top := fmt.Sprintf("/*%s+", line)
		bot := fmt.Sprintf("+%s*/", line)

		fmt.Println(top)
		fmt.Println(str)
		fmt.Print(bot)

		output := fmt.Sprintf("%s\n%s\n%s", top, str, bot)
		if !clipboard.Unsupported {
			clipboard.WriteAll(output)
		}
		return
	}
	fmt.Println("Please provide a comment to prettify")
}
