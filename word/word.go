package word

import (
	"log"
)

type Word struct {
	English string
	German string
	Gender string
	Pronounce string
	SentenceGerman string
	SentenceEnglish string
}

// word loader: make an interface and use that to load words.
// for example, I'm going to use a json type of thing.

// TODO: I can have the guesses return something like % of correct characters
// for an incorrect guess

// TODO: Should I put the configuration in here that will decide about
// case insensitivity and special characters?

// TODO: maybe I should just have a Matcher passed to the Guess functions that does 
// the above.

// GuessEnglish is used to guess the translated word in English
func (w *Word) GuessEnglish(g string, m Matcher) bool {
	if m == nil {
		log.Println("Matcher is nil. Cannot compare words.")
		return false
	}

	return m.Match(g, w.English)
}

// GuessGerman is used to guess the translated word in German
func (w *Word) GuessGerman(g string, m Matcher) bool {
	if m == nil {
		log.Println("Matcher is nil. Cannot compare words.")
		return false
	}

	return m.Match(g, w.German)
}

// GuessGender is used to guess the gender of a noun
func (w *Word) GuessGender(g string) bool {
	if g == w.Gender {
		return true
	}
	return false
}
