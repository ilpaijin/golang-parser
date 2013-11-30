package main

import (
	"encoding/csv"
	"fmt"
	"ilpaijin"
	"log"
	"os"
	"reflect"
	// "strings"
)

func main() {

	if len(os.Args) == 1 {
		log.Fatal("Missing arguments, 1st is the xmlType (ex. pregame), 2nd is the parseType (ex. full/light/jackpot)")
	}

	log.Print(os.Args)

	parsedSet, err := ilpaijin.Parser(os.Args[0], os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	c := parsedSet[0].Interface()

	// log.Fatal(c)

	f, err := os.Create("./arrayOfSport.csv")
	if err != nil {
		log.Fatal("Error creating csv file")
	}
	defer f.Close()

	w := csv.NewWriter(f)

	switch reflect.TypeOf(c).Kind() {

	case reflect.Struct:

		s := reflect.ValueOf(c)

		// fName := s.FieldByName("XMLName").Interface()

		// typeOfT := s.Type()

		var record []string

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)

			// m := f.Interface()
			// realMappings := m.(map[string]string)
			// log.Print(realMappings)
			// log.Fatal(typeOfT.Field(i).Name, f.Type(), f.Interface())
			// record = append(record, strings.ToLower(typeOfT.Field(i).Name))
			v := f.Interface()
			val := reflect.ValueOf(v)
			// log.Print(v)
			// log.Print(val)
			// fmt.Printf("%s", val.Interface())
			// log.Fatal(val.String())
			record = append(record, val.Interface().(string))
			w.Write(record)
		}

	}

	w.Flush()

	db := ilpaijin.Db()
	defer db.Close()

	idutente := ilpaijin.GetUserByIdCoupon(db, "19125605")

	fmt.Println(idutente)
}
