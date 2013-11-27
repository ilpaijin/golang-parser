package main

import (
	"fmt"
	"ilpaijin"
	"log"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Missing arguments, 1st is the xmlType (ex. pregame), 2nd is the parseType (ex. full/light/jackpot)")
	}

	parsedSet, err := ilpaijin.Parser(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(parsedSet)

	db := ilpaijin.Db()
	defer db.Close()

	idutente := ilpaijin.GetUserByIdCoupon("19125605")

	fmt.Println(idutente)
}
