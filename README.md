# go_unit_testing_example
Unit testing with dependency injection via interface and methods creation

### Model
The model simulates retrieving data from the databse
The main focus is the RetrievePassesResults() function. By converting it to a method, the method can be overridden for testing purposes in the controller folder.

### Controller
The controller simulates manipulating and creating logic for the results that were retrieved from the database.
To ensure that unit testing can be done on the controller, an interface is passed as a parameter to the functions that are created in this folder. Hence, this allows us to customise the child object during testing. 

### Importing GoConvey 
$ go get github.com/smartystreets/goconvey <br />
$ $GOPATH/bin/goconvey <br />
(more info: https://smartystreets.com/blog/2015/05/go-testing-part-3-convey-behavior/)

### Using testing coverage
$ go test -coverprofile=coverage.out  <br />
$ go tool cover -html=coverage.out <br />
 
