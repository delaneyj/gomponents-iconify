package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterPosPauseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 12a2 2 0 0 0-2-2h-1V4H7v6H6a2 2 0 0 0-2 2v7h9c0-.7.13-1.37.35-2H6v-5h12v1.09c.33-.05.66-.09 1-.09c.34 0 .67.04 1 .09V12m-5-2H9V6h6v4m-8 5v-2h4v2H7m9 1h2v6h-2v-6m6 0v6h-2v-6h2Z"/>`),
		g.Group(children),
	)
}