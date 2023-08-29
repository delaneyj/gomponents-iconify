package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterPosPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 13.09V12a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v7h9c0-3.31 2.69-6 6-6c.34 0 .67.04 1 .09M10 14H6v-2h4v2m7-5H7V4h10v5m5 10l-5 3v-6l5 3Z"/>`),
		g.Group(children),
	)
}