package main

import (
	"flag"
	"fmt"
	"os"
	"xkcd/cli/pkg/api"
	"xkcd/cli/pkg/postgres"
	"xkcd/cli/pkg/search"
	"xkcd/cli/pkg/table"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func init() {
	Conn, _ = postgres.GetConnection()
	table.CreateTable(Conn)
}

func main() {

	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)

	searchString := searchCmd.String("keyword", "", "Search by keyword")

	getStart := getCmd.Int("start", 0, "Start Index")

	getEnd := getCmd.Int("end", 0, "End Index")

	if len(os.Args) < 2 {
		fmt.Println("Expected Subcommand")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "get":
		getCmd.Parse(os.Args[2:])

	case "search":
		searchCmd.Parse(os.Args[2:])

	default:
		fmt.Println("Expected Subcomman")
		os.Exit(0)
	}

	if getCmd.Parsed() {
		if (*getStart == 0 && *getEnd == 0) || (*getStart >= *getEnd) {
			fmt.Println("Invalid Start And End Index")
			os.Exit(0)
		}

		for i := *getStart; i <= *getEnd; i++ {
			data, _ := api.GetComic(int64(i))
			table.Insert(data, Conn)

		}

	}

	if searchCmd.Parsed() {
		if *searchString == "" {
			fmt.Println("Invalid Search Param Use -keyword <VALUE>")
			os.Exit(0)
		}

		search.SearchTitle(Conn, *searchString)

	}

}
