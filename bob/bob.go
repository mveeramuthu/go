// Package bob says Hey!
package bob

import "strings"

const testVersion = 3

// Hey responds to remarks
func Hey(remark string) string {

	var response string = "Whatever."

	remark = strings.TrimSpace(remark)
	if isSilent(remark) {
		response = "Fine. Be that way!"
	} else if isYelling(remark) {
		response = "Whoa, chill out!"
	} else if isQuestion(remark) {
		response = "Sure."
	}

	return response
}

// Checks whether the input string is a question
func isQuestion(remark string) bool {
	//Check if string ends with a '?'
	return remark[len(remark)-1:] == "?"
}

// Checks whether the input string's emotion is yelling
func isYelling(remark string) bool {
	// Check if string is all CAPS
	return remark == strings.ToUpper(remark) && remark != strings.ToLower(remark)
}

func isSilent(remark string) bool {
	return len(remark) == 0
}

func main() {

}