// Exercism side exercise about Bob.

package bob

import (
	"strings"
)

// Hey represents Bob's answers, based on input 'remark' argument.
// Bob answers 'Sure.' if you ask him a question, such as "How are you?".
// He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
// He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
// He says 'Fine. Be that way!' if you address him without actually saying anything.
// He answers 'Whatever.' to anything else.
func Hey(remark string) string {

	switch {
	case isQuestion(remark) && isShouting(remark):
		return "Calm down, I know what I'm doing!"
	case isShouting(remark):
		return "Whoa, chill out!"
	case isQuestion(remark):
		return "Sure."
	case isEmpty(remark):
		return "Fine. Be that way!"
	}

	return "Whatever."

}

func isEmpty(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}

func isShouting(text string) bool {
	return text == strings.ToUpper(text) && strings.ToLower(text) != text
}

func isQuestion(text string) bool {
	return strings.HasSuffix(strings.TrimSpace(text), "?")
}
