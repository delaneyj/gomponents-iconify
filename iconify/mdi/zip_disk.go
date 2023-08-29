package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZipDisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3L3 5v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5l-4-2v2a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1V3H7m1 7h8a1 1 0 0 1 1 1v8H7v-8a1 1 0 0 1 1-1Z"/>`),
		g.Group(children),
	)
}