package main

import (
	"fmt"
	"os"

	"github.com/ikawaha/zenn-importer/cmd/hatena"
	"github.com/ikawaha/zenn-importer/cmd/qiita"
)

type subcommand struct {
	Name string
	Cmd  func([]string) error
}

var subcommands = []subcommand{
	{
		Name: "qiita",
		Cmd:  qiita.Cmd,
	},
	{
		Name: "hatena",
		Cmd:  hatena.Cmd,
	},
}

func main() {
	var cmd func([]string) error
	if len(os.Args) >= 2 {
		for i := range subcommands {
			if os.Args[1] == subcommands[i].Name {
				cmd = subcommands[i].Cmd
			}
		}
	}
	if cmd == nil {
		fmt.Fprintln(os.Stderr, "zenn-importer: a command to save blog posts in zenn format.")
		fmt.Fprintln(os.Stderr, "subcommands:")
		for _, v := range subcommands {
			fmt.Println("  -", v.Name)
		}
		os.Exit(1)
	}
	if err := cmd(os.Args[2:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
