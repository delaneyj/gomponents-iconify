package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feQuoteLeft0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feQuoteLeft1" fill="currentColor"><path id="feQuoteLeft2" d="M7 21a4 4 0 0 1-4-4c0-1.473 1.333-6.14 4-14h2L7 13a4 4 0 1 1 0 8Zm10 0a4 4 0 0 1-4-4c0-1.473 1.333-6.14 4-14h2l-2 10a4 4 0 1 1 0 8Z"/></g></g>`),
		g.Group(children),
	)
}