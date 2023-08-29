package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileXmlBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5c-1.11 0-2 .89-2 2v14c0 1.11.89 2 2 2h14c1.11 0 2-.89 2-2V5c0-1.11-.89-2-2-2M8 15H6.5L6 13l-.5 2H4l.75-3L4 9h1.5l.5 2l.5-2H8l-.75 3L8 15m7.5 0H14v-4.5h-1V14h-1.5v-3.5h-1V15H9v-4c0-1.1.9-2 2-2h2.5a2 2 0 0 1 2 2v4m4.5 0h-3V9h1.5v4.5H20V15Z"/>`),
		g.Group(children),
	)
}