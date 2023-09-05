package table

import (
	"context"
	"fmt"
	"strings"
	"xkcd/cli/pkg/api"

	"github.com/jackc/pgx/v5"
)

func Insert(c *api.Comic, conn *pgx.Conn) bool {

	transcript := strings.ReplaceAll(c.Transcript, "\n", "")

	_, err := conn.Exec(context.Background(), "INSERT INTO comics(id,month,link,year,news,safe_title,transcript,alt,img,title,day) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", c.Num, c.Month, c.Link, c.Year, c.News, c.SafeTitle, transcript, c.Alt, c.Img, c.Title, c.Day)

	if err != nil {
		return false
	}
	fmt.Printf("Written Comic With Id %d\n", c.Num)
	return true
}
