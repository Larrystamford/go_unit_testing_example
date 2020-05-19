package main



import (
	"fmt"
	"learningTesting/model_service"
	"learningTesting/controller_service"
)


func main() {
	resultsObject := model_service.ResultStruct{} 
	// instead of calling RetrieveFailureResults() directly, we instantiate this specific Results-Object from controller service that 
	// calls the RetrieveFailureResults() function that we want to implement 

	respond := controller_service.DeterminePass(resultsObject)

	// The final contoller that will determine if URS deployment is allowed
	if respond {
		fmt.Println("URS passed please proceed to deployment.")
	} else {
		fmt.Println("URS failed not allowed to deployed.")
	}
}
