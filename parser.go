package main

import (
	"fmt"
	"ilpaijin"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Missing arguments, 1st is the xmlType (ex. pregame/coupons), 2nd is the parseType (ex. full/light/jackpot)")
	}

	parsedSet := ilpaijin.Parse(os.Args[1], os.Args[2])

	fmt.Println(parsedSet)
	// parsedSet.Data.Csvize()

	db := ilpaijin.Db()
	defer db.Close()

	idutente := ilpaijin.GetUserByIdCoupon(db, "19125605")

	fmt.Println(idutente)
}
