package instant

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"

	"golang.org/x/text/language"
)

// Coin is an instant answer
type Coin struct {
	Answer
}

func (c *Coin) setQuery(r *http.Request, qv string) answerer {
	c.Answer.setQuery(r, qv)
	return c
}

func (c *Coin) setUserAgent(r *http.Request) answerer {
	return c
}

func (c *Coin) setLanguage(lang language.Tag) answerer {
	c.language = lang
	return c
}

func (c *Coin) setType() answerer {
	c.Type = "coin toss"
	return c
}

func (c *Coin) setRegex() answerer {
	triggers := []string{
		"flip a coin", "heads or tails", "coin toss",
	}

	t := strings.Join(triggers, "|")
	c.regex = append(c.regex, regexp.MustCompile(fmt.Sprintf(`^(?P<trigger>%s)$`, t)))

	return c
}

func (c *Coin) solve(r *http.Request) answerer {
	choices := []string{"Heads", "Tails"}

	c.Solution = choices[rand.Intn(2)]

	return c
}

func (c *Coin) setCache() answerer {
	c.Cache = false
	return c
}

func (c *Coin) tests() []test {
	tests := []test{}

	for _, q := range []string{"flip a coin", "heads or tails", "Coin Toss"} {
		tst := test{
			query: q,
			expected: []Data{
				{
					Type:      "coin toss",
					Triggered: true,
					Solution:  "Heads",
					Cache:     false,
				},
				{
					Type:      "coin toss",
					Triggered: true,
					Solution:  "Tails",
					Cache:     false,
				},
			},
		}

		tests = append(tests, tst)
	}

	return tests
}
