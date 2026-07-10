package main

import (
	"os"

	"gitlab.com/shaninalex/forgecore/app/cmd"
)

func main() {
	os.Exit(cmd.Execute())
}
