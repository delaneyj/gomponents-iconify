package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LowLevelNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsLowLevelNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24.493 36.569c.375.116.78.091 1.139-.07l2.589-1.373c.732-.985.845-1.472.58-2.502l-1.72-2.372a1.615 1.615 0 0 0-.991-.565l-13.983-.8c-1.978-.311-2.454 1.657-.541 2.25l12.927 5.432ZM6.109 24a17.926 17.926 0 0 1 4.956-10.52l.228.227l2.121 2.121l.707.707l1.415-1.414l-.708-.707l-2.121-2.121l-.172-.172A17.925 17.925 0 0 1 23 8.027V13h2V8.027a17.924 17.924 0 0 1 10.465 4.094l-2.172 2.172l-.707.707L34 16.414l.707-.707l2.228-2.228A17.926 17.926 0 0 1 41.891 24h-4.207v2H42a17.905 17.905 0 0 1-4.106 11.427l-.48-.48l-2.233-2.233l-.707-.707l-1.414 1.414l.707.707L36 38.361l.537.537l-.013.012l1.391 1.437A19.928 19.928 0 0 0 44 25.99C44 14.95 35.045 6 24 6S4 14.949 4 25.99a19.923 19.923 0 0 0 5.79 14.065l1.42-1.407l-.019-.02l.066-.065l2.233-2.233l.707-.707l-1.414-1.414l-.708.707l-2.21 2.211A17.902 17.902 0 0 1 6 26h4.053v-2H6.109Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsLowLevelNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}