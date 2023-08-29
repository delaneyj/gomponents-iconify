package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 19a1 1 0 0 0 1 1h2v2h-2.5c-.55 0-1.5-.45-1.5-1c0 .55-.95 1-1.5 1H8v-2h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H8V2h2.5c.55 0 1.5.45 1.5 1c0-.55.95-1 1.5-1H16v2h-2a1 1 0 0 0-1 1v14Z"/>`),
		g.Group(children),
	)
}