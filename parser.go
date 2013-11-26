package main

import (
	"fmt"
	"ilpaijin"
	"os"
)

func main() {

	ilpaijin.Parser(os.Args[1], os.Args[2])

	db := ilpaijin.Db()
	defer db.Close()

	idutente := ilpaijin.GetUserByIdCoupon("19125605")

	fmt.Println(idutente)
}
