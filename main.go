package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const cow = `         \  ^__^
          \ (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
		`

const cat = `                    \
                     \ 
   *                  |\___/|
             .        )     (             .              '
                     =\     /=
'                      )===(       *               .
              *       /     \
                      |     |                 *
      '              /       \
                     \       /      '             
              _/\_/\_/\__  _/_/\_/\_/\_/\_/\_/\_/\_/\_/\_
         '    |  |  |  |( (  |  |  |  |  |  |  |  |  |  |
              |  |  |  | ) ) |  |  |  |  |  |  |  |  |  |
.             |  |  |  |(_(  |  |  |  |  |  |  |  |  |  |
              |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
         .    |  |  |  |  |  |  |  |  |  |  |  |  |  |  |`

const cat2 = `                    \
                     \  
     '          *     |\___/|
                     =) ^Y^ (=            .              '
 .                    \  ^  /
                       )=*=(       *             '
              '       /     \
                      |     |                        *   
       *             /| | | |\          '
                     \| | |_|/\                     .  
 .            _/\_/\_//_// ___/\_/\_/\_/\_/\_/\_/\_/\_/\_
              |  |  |  | \_) |  |  |  |  |  |  |  |  |  |
         *    |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
     .        |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
              |  |  |  |  |  |  |  |  |  |  |  |  |  |  |
              |  |  |  |  |  |  |  |  |  |  |  |  |  |  |`

// var sideBorders = []rune{'_', '-'}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("xsCowsay only works with stdin pipes.")
		fmt.Println("Usage: echo \"Hello\" | xsaCowsay")
		os.Exit(1)
	}

	var lines []string

	rd := bufio.NewReader(os.Stdin)
	for {
		line, _, err := rd.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	lines = tabsToSpace(lines)
	maxWidth := calcMaxWidth(lines)
	lines = normalizeLines(lines, maxWidth)
	lines = addBorders(lines, maxWidth)

	for _, v := range lines {
		fmt.Printf("%s\n", v)
	}
	fmt.Println(cat2)
}

func tabsToSpace(lines []string) []string {
	var newLines []string
	for _, v := range lines {
		newLines = append(newLines, strings.ReplaceAll(v, "\t", "    "))
	}
	return newLines
}

func calcMaxWidth(lines []string) int {
	maxWidth := 0
	for _, v := range lines {
		if len(v) > maxWidth {
			maxWidth = len(v)
		}
	}
	return maxWidth
}

func normalizeLines(lines []string, maxWidth int) []string {
	var newLines []string

	for _, v := range lines {
		padding := maxWidth - len(v)
		newLines = append(newLines, v+strings.Repeat(" ", padding))
	}
	return newLines
}

func addBorders(lines []string, maxWidth int) []string {
	newLines := []string{" " + strings.Repeat("_", maxWidth+2)}
	count := len(lines)

	if count == 1 {
		newLines = append(newLines, "< "+lines[0]+" >")
	} else if count == 2 {
		newLines = append(newLines, "/ "+lines[0]+" \\")
		newLines = append(newLines, "\\ "+lines[1]+" /")
	} else {
		newLines = append(newLines, "/ "+lines[0]+" \\")
		for i := 1; i < count-1; i++ {
			newLines = append(newLines, "| "+lines[i]+" |")
		}
		newLines = append(newLines, "\\ "+lines[len(lines)-1]+" /")
	}

	newLines = append(newLines, " "+strings.Repeat("-", maxWidth+2))

	return newLines
}
