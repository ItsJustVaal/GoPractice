package cmd

import (
	"fmt"
	"os"
)

func main() {
	// home, _ := homedir.Dir()
	// dbPath := filepath.Join(home, "tasks.db")
	// // must(cmd.Init(dbPath))
	// // must(cmd.Execute())

}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
