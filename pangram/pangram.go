package pangram

import "strings"

type PangramChecker struct {
	Alphabet []rune
}

func (p *PangramChecker) Check(input string) bool {
	input = strings.ToLower(input)
	missingRunes := p.createRuneSet()
	counter := len(missingRunes)

	for _, r := range input {
		if value, exists := missingRunes[r];
			exists && !value {
			missingRunes[r] = true
			counter--
		}

		if counter == 0 {
			return true
		}
	}

	return false
}

func (p *PangramChecker) createRuneSet() map[rune]bool {
	set := map[rune]bool{}
	for _, v := range p.Alphabet {
		set[v] = false
	}

	return set
}
