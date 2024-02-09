package main
import (
	"bufio"
	"flag"
	"fmt"
	// "log"
	// "net/http"
	"os"
	"strings"
	"github.com/yutak1026/kawasql/db"
)

func showTitle() {
	title := `This is KawaDB created by kawmaura.`
	fmt.Println(title)
}

func client() {
	showTitle()
	stdin := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">>")
		stdin.Scan()
		q := stdin.Text()

		var err error
		if strings.HasPrefix(q, "exit") {
			fmt.Print("quit\n")
			db.Execute(q)
		} else {
			fmt.Print("start SQL command\n")
		}

		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	flag.Parse()

	// if *serverMode {
	// 	server()
	// 	return
	// }

	client()
}