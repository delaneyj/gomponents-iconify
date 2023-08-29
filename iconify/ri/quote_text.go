package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuoteText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4H3v2h18V4Zm0 7H8v2h13v-2Zm0 7H8v2h13v-2ZM5 11H3v9h2v-9Z"/>`),
		g.Group(children),
	)
}