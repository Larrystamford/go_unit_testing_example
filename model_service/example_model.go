package model_service

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

func goDotEnvVariable(key string) string {
  // load .env file
  err := godotenv.Load(".env")
  if err != nil {
    log.Fatalf("Error loading .env file")
  }
  return os.Getenv(key)
}

var db *sql.DB

var (
  host     = goDotEnvVariable("DB_HOST")
  port     = goDotEnvVariable("DB_PORT")
  user     = goDotEnvVariable("DB_USER")
  password = goDotEnvVariable("DB_PASSWORD")
  dbname   = goDotEnvVariable("DB_NAME")
)


type ResultsDB struct {
	Id string  `json:"Id"`
	Result string `json:"Result"`
}



// initialise database
func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	// ping checks if the database exist ( connecting to postgres alone does not mean table is there )
	if err = db.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("You connected to your database.")
}

// SELECT all results from table and return results as list
func getAllResults() ([]ResultsDB){
	rows, err := db.Query("SELECT * FROM FoodTests;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	resultsList := make([]ResultsDB, 0)
	for rows.Next() {
		eachResult := ResultsDB{}
		err := rows.Scan(&eachResult.Id, &eachResult.Result) // order matters
		if err != nil {
			panic(err)
		}
		resultsList = append(resultsList, eachResult)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	return resultsList
}


// Do not define interfaces before they are used: without a realistic example of usage, it is too difficult to see whether an interface is even necessary, let alone what methods it ought to contain.
// Hence no interface declared here, ResultStruct is an object that executes the RetrievePassesResults() method. 

type ResultStruct struct {
	passes float64
	total float64
}


func (rs ResultStruct) RetrievePassesResults() (float64, float64){
	// return the total number of passes and database table length

	fmt.Println("Database is queried.")
	resultsList := getAllResults()
	

	for _, eachRes := range resultsList {
		if eachRes.Result == "pass"{
			rs.passes++
		}
		rs.total++
	}
	fmt.Printf("Test Results: %.0f / %.0f tests passed.\n", rs.passes, 	rs.total)
	return rs.passes, rs.total
}




