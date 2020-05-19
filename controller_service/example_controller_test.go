package controller_service

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

// creating an object to allow us to mock the number of passes and failures 
type mockResults struct {
	passes float64
	total float64
}

// implementing the function exactly how we want to with the mock results
// aka overriding this method
func (mr mockResults) RetrievePassesResults() (float64, float64){
	return mr.passes, mr.total
}


func TestDeterminePass(t *testing.T) {

	
	Convey("Given a result struct", t, func() {

		Convey("It should return true if the result is more than or equals to 0.5", func() {
			mr := mockResults{90,100}
			So(DeterminePass(mr), ShouldEqual, true )
		})

		Convey("It should return false if the result is less than 0.5",func() {
			mr := mockResults{1,100}
			So(DeterminePass(mr), ShouldEqual, false)
		})
	})


}
