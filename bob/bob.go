// Package bob implemets the bob exercise
// Bob is a lackadaisical teenager. In conversation, his responses are very limited.
// Bob answers 'Sure.' if you ask him a question.
// He answers 'Whoa, chill out!' if you yell at him.
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
// He says 'Fine. Be that way!' if you address him without actually saying anything.
// He answers 'Whatever.' to anything else.
package bob

import "strings"
import "unicode"

// Hey returns Bob's response on remark
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	switch {
	case remark == "":
		return "Fine. Be that way!"
	case isYelling(remark) && isQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isYelling(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

func isYelling(remark string) bool {
	return hasLetters(remark) && strings.ToUpper(remark) == remark
}

func isQuestion(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func hasLetters(text string) bool {
	for _, char := range text {
		if unicode.IsLetter(char) {
			return true
		}
	}

	return false
}
