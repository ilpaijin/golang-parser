package ilpaijin

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Db() *sql.DB {

	var dbname = "staging"
	var user = "root"
	var pwd = "panissa"

	q, err := sql.Open("mysql", user+":"+pwd+"@/"+dbname)

	if err != nil {
		log.Fatal("Error db ", err)
	}

	return q
}

func GetUserByIdCoupon(db *sql.DB, idcoupon string) (r int) {
	var (
		idCoupon int
		idUtente int
		res      int
	)

	rows := db.QueryRow("SELECT idCoupon, idUtente FROM jackpot_coupons WHERE idCoupon=?", idcoupon).Scan(&idCoupon, &idUtente)

	switch {
	case rows == sql.ErrNoRows:
		res = 0
	case rows != nil:
		log.Fatal(rows)
	default:
		res = idUtente
	}
	return res
}
