package main

import (
	"encoding/csv"
	"fmt"
	"ilpaijin"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Missing arguments, 1st is the xmlType (ex. pregame/coupons), 2nd is the parseType (ex. full/light/jackpot)")
	}

	parsedSet := ilpaijin.Parser(os.Args[1], os.Args[2])

	f, err := os.Create("data/arrayOfSport.csv")
	if err != nil {
		log.Fatal("Error creating csv file")
	}
	defer f.Close()

	w := csv.NewWriter(f)

	var record []string

	for _, q := range parsedSet.Quote {

		record = append(record, q.TipoQuota)

		w.Write(record)
	}

	w.Flush()

	db := ilpaijin.Db()
	defer db.Close()

	idutente := ilpaijin.GetUserByIdCoupon(db, "19125605")

	fmt.Println(idutente)
}
