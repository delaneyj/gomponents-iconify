package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripSplitThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 3a1 1 0 1 0-2 0v26a1 1 0 1 0 2 0V3ZM6.5 4h7v2h-7A2.5 2.5 0 0 0 4 8.5v15A2.5 2.5 0 0 0 6.5 26h7v2h-7A4.5 4.5 0 0 1 2 23.5v-15A4.5 4.5 0 0 1 6.5 4Zm19 22h-7v2h7a4.5 4.5 0 0 0 4.5-4.5v-15A4.5 4.5 0 0 0 25.5 4h-7v2h7A2.5 2.5 0 0 1 28 8.5v15a2.5 2.5 0 0 1-2.5 2.5ZM24 9a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0V9Zm1 5a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1Zm-1 7a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0v-2ZM7 8a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0V9a1 1 0 0 0-1-1Zm-1 7a1 1 0 1 1 2 0v2a1 1 0 1 1-2 0v-2Zm1 5a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}