package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterPos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 10H6a2 2 0 0 0-2 2v7h16v-7a2 2 0 0 0-2-2m0 4h-4v-2h4m-1-3H7V4h10Z"/>`),
		g.Group(children),
	)
}