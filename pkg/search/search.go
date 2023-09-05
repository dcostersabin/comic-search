package search

import (
	"context"
	"encoding/json"
	"fmt"
	"xkcd/cli/pkg/api"

	"github.com/jackc/pgx/v5"
)

func SearchTitle(conn *pgx.Conn, query string) {

	rows, err := conn.Query(context.Background(), "SELECT month,link,year,news,safe_title as safeTitle,transcript,alt,img,title,day,id as num from comics where transcript LIKE "+"'%"+query+"%'"+" OR "+"month LIKE "+"'%"+query+"%'"+" OR "+"news LIKE "+"'%"+query+"%'"+" OR "+"link LIKE "+"'%"+query+"%'")

	if err != nil {
		fmt.Printf("%v", err)
	}
	rowDetails, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[api.Comic])

	if err != nil {
		fmt.Printf("%v", err)
	}
	for _, p := range rowDetails {
		str, _ := json.MarshalIndent(p, "", "\t")
		fmt.Println(string(str))
	}

}
