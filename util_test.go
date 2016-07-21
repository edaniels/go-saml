package saml

import (
	"fmt"
	"regexp"

	. "gopkg.in/check.v1"
)

// -----------------------------------------------------------------------
// From: https://github.com/dropbox/godropbox/blob/master/gocheck2/checkers.go
// MultilineErrorMatches: Multiline ErrorMatches
// The standard gocheck ErrorMatches brackets the regular expression that
// the error must match in ^ and $, so that it can only match single-line
// errors messages. Most dropbox errors are created using godropbox.errors,
// which produce error objects that have multiline message, which means
// that our standard errors will never be matched by ErrorMatches.
//
// This is a variant of the normal ErrorMatches which avoids that problem,
// and works with dropbox errors.
// It takes two parameters:
// 1: An error object, and
// 2: a string containing a regular expression.
// The check succeeds if the error's message contains a match for the regular expression
// at any position.
type multilineErrorMatches struct{}

func (e multilineErrorMatches) Check(params []interface{}, names []string) (bool, string) {
	if len(params) != 2 {
		return false, "MultilineErrorMatches take 2 arguments: an error, and a regular expression"
	}
	errValue, errIsError := params[0].(error)
	if !errIsError {
		return false, "the first parameter value must be an error!"
	}
	regexpStr, reIsStr := params[1].(string)
	if !reIsStr {
		return false, "the second parameter value must be a string containing a regular expression"
	}
	matches, err := regexp.MatchString(regexpStr, errValue.Error())
	if err != nil {
		return false, fmt.Sprintf("Error in regular expression: %v", err.Error())
	}
	return matches, ""
}

func (h multilineErrorMatches) Info() *CheckerInfo {
	return &CheckerInfo{
		Name:   "MultilineErrorMatches",
		Params: []string{"error", "pattern"},
	}
}

var MultilineErrorMatches multilineErrorMatches = multilineErrorMatches{}
