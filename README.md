# go_unit_testing_example
Unit testing with dependency injection via interface and methods creation

# Model
The model simulates retrieving data from the databse
The main focus is the RetrievePassesResults() function. By converting it to a method, the method can be overridden for testing purposes in the controller folder.

# Controller
The controller simulates manipulating and creating logic for the results that were retrieved from the database.
To ensure that unit testing can be done on the controller, an interface is passed as a parameter to the functions that are created in this folder. Hence, this allows us to customise the child object during testing. 
