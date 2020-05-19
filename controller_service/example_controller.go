package controller_service

import (
)


// Results interface required here, hence it is declared here.
type Results interface {
	RetrievePassesResults() (float64, float64)
}

// This is the unit function that we want to test. Hence, RetrievePassesResults was isolated as a method (implemented with interface) instead of existing as a function.
func DeterminePass(results Results) (bool) {
	// if the percentage of passes is less than 50% return false
	// else return true

	passNum, totalNum := results.RetrievePassesResults()

	if passNum/totalNum >= 0.5 {
		return true
	}
	
	return false
}