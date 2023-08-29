package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PeopleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPeopleNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 17c0 3.867-3.133 7-7 7s-7-3.133-7-7s3.133-7 7-7s7 3.133 7 7Zm9.5 9c3.039 0 5.5-2.461 5.5-5.5S36.539 15 33.5 15a5.499 5.499 0 0 0-5.5 5.5c0 3.039 2.461 5.5 5.5 5.5ZM17 26c2.734 0 7.183.851 10.101 2.545C28.293 29.758 29 31.081 29 32.4V38H4v-5.6c0-4.256 8.661-6.4 13-6.4Zm27 12H31v-5.6c0-1.416-.511-2.72-1.324-3.883c1.541-.345 3.058-.517 4.217-.517C37.62 28 44 29.787 44 33.333V38Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPeopleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}