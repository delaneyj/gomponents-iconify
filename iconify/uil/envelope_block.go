package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeBlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 13.26a5 5 0 1 0-3.5-1.47a5 5 0 0 0 3.5 1.47Zm2.12-2.88a3 3 0 0 1-3.4.58l4-4a3 3 0 0 1-.6 3.42Zm-4.24-4.24a3 3 0 0 1 2.12-.88a3 3 0 0 1 1.28.3l-4 4a3 3 0 0 1 .6-3.42Zm5.12 8.12a1 1 0 0 0-1 1v4a1 1 0 0 1-1 1h-14a1 1 0 0 1-1-1V9.67l5.88 5.88a3 3 0 0 0 2.11.88a3 3 0 0 0 2.16-.91a1 1 0 0 0 0-1.39a1 1 0 0 0-1.43 0a1 1 0 0 1-1.4 0L4.91 8.26H9.5a1 1 0 0 0 0-2h-5a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-4a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}