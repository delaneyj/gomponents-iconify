package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Floppy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5.5L18.5 3H17v6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V3H5m7 1v5h3V4h-3m-5 8h10a1 1 0 0 1 1 1v6H6v-6a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}