package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/jawher/mow.cli"
	"github.com/jmcvetta/neoism"
)

func main() {
	app := cli.App("cyphertool", "simple utility for doing things with cypher")
	app.Command(
		"run",
		"run cypher statements against a neo4j database. Typically useful for bulk importing of data",
		func(cmd *cli.Cmd) {
			url := cmd.String(cli.StringOpt{
				Name:  "url",
				Desc:  "neo4j url",
				Value: "http://localhost:7474/db/data",
			})
			echo := cmd.Bool(cli.BoolOpt{
				Name:  "echo",
				Desc:  "echo queries as they are run",
				Value: false,
			})
			cmd.Action = func() {
				if err := run(*url, *echo); err != nil {
					log.Fatal(err)
				}
			}
		},
	)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(url string, echo bool) error {
	db, err := neoism.Connect(url)
	if err != nil {
		return err
	}

	br := bufio.NewReader(os.Stdin)
	for {
		stmt, err := br.ReadString(';')
		switch err {
		case nil:
			err = runQueryIfNotEmpty(db, stmt, echo)
			if err != nil {
				return err
			}
		case io.EOF:
			return runQueryIfNotEmpty(db, stmt, echo)
		}
	}
}

func runQueryIfNotEmpty(db *neoism.Database, query string, echo bool) error {
	stmt := strings.TrimSpace(query)
	if len(stmt) > 0 {
		if echo {
			defer fmt.Println(stmt)
		}

		cq := neoism.CypherQuery{
			Statement: string(stmt),
		}
		return db.Cypher(&cq)
	}
	return nil

}
