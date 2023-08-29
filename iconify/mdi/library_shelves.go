package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibraryShelves(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 9V1.5h-3V9h-3V1.5h-3V9h-3V1.5H4.65V9H3v1.5h18V9h-1.5m0 4.5h-3V21h-3v-7.5h-3V21h-3v-7.5H4.65V21H3v1.5h18V21h-1.5v-7.5Z"/>`),
		g.Group(children),
	)
}