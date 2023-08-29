package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmstripSplitThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 3a1 1 0 1 0-2 0v26a1 1 0 1 0 2 0V3ZM6.5 4h7v24h-7A4.5 4.5 0 0 1 2 23.5v-15A4.5 4.5 0 0 1 6.5 4ZM6 8a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0V9a1 1 0 0 0-1-1Zm-1 7v2a1 1 0 1 0 2 0v-2a1 1 0 1 0-2 0Zm1 5a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1Zm19.5 8h-7V4h7A4.5 4.5 0 0 1 30 8.5v15a4.5 4.5 0 0 1-4.5 4.5ZM25 9v2a1 1 0 1 0 2 0V9a1 1 0 1 0-2 0Zm1 5a1 1 0 0 0-1 1v2a1 1 0 1 0 2 0v-2a1 1 0 0 0-1-1Zm-1 7v2a1 1 0 1 0 2 0v-2a1 1 0 1 0-2 0Z"/>`),
		g.Group(children),
	)
}