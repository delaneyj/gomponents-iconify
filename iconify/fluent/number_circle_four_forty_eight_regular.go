package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFourFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 24c0-9.665 7.835-17.5 17.5-17.5S41.5 14.335 41.5 24S33.665 41.5 24 41.5S6.5 33.665 6.5 24ZM24 4C12.954 4 4 12.954 4 24s8.954 20 20 20s20-8.954 20-20S35.046 4 24 4Zm3.5 10.58c0-1.636-2.124-2.274-3.026-.91l-9.209 13.925c-.681 1.03.058 2.405 1.293 2.405H25v2.75a1.25 1.25 0 1 0 2.5 0V30h2.25a1.25 1.25 0 1 0 0-2.5H27.5V14.58ZM25 17.406V27.5h-6.675L25 17.406Z"/>`),
		g.Group(children),
	)
}