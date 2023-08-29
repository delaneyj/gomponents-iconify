package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feQuoteRight0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feQuoteRight1" fill="currentColor"><path id="feQuoteRight2" d="M17 11a4 4 0 1 1 4-4c0 1.473-1.333 6.14-4 14h-2l2-10ZM7 11a4 4 0 1 1 4-4c0 1.473-1.333 6.14-4 14H5l2-10Z"/></g></g>`),
		g.Group(children),
	)
}