package main

import(
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"regexp"
)

var ip string = "192.168.0.100"
var url_regex, errRx = regexp.Compile(`(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?`)

func main(){

	if errRx != nil {
		fmt.Fprintf(os.Stderr, "Error compiling regex")
		return
	}

	/*file_name := os.Args[1]*/

	file_name := "output/output2.txt"

	file, err := os.Open(file_name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "The file %s does not exist.\n", file_name)
		return
	}

	bfr := bufio.NewReader(file)

	for line, _, err := bfr.ReadLine(); err != io.EOF; line, _, err = bfr.ReadLine() {

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading line")
			continue
		}

		s := string(line)

		if strings.Contains(s, "les")  {
			fmt.Printf("%s\n", s)
		}
	}
}

/* deprecated */
func findUrl(line string){
	result := url_regex.MatchString(line)
	if result {
		fmt.Printf("%s\n", line)
	}
}
