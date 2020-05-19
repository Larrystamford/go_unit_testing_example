package model_service



import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/DATA-DOG/go-sqlmock"

)


// By mocking the db, we ensure that the SQL logic is correct
// We also ensure that the return results of pass, total and resultsList is correct
func TestGetAllResults(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Id", "Result"}).AddRow("1", "pass").AddRow("2", "pass").AddRow("3", "pass").AddRow("4", "fail").AddRow("5", "pass")

	mock.ExpectQuery("SELECT (.*) FROM FoodTests").WillReturnRows(rows)


	resultsList, passes, total := getAllResults(db)


	expectedResultsList := []ResultsDB{
		ResultsDB{"1", "pass"},
		ResultsDB{"2", "pass"},
		ResultsDB{"3", "pass"},
		ResultsDB{"4", "fail"},
		ResultsDB{"5", "pass"}}
	
		

	Convey("Given a this particular mock database", t, func() {
		Convey("Number of passes should be 4", func() {
			So(passes, ShouldEqual, 4)
		})
		Convey("Length of database should be 5", func() {
			So(total, ShouldEqual, 5)
		})
		Convey("Expected query results should match", func() {
			So(resultsList, ShouldResemble , expectedResultsList)
		})
	})

}