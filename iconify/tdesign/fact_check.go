package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3h22Zm-2 2H3v14h18V5Zm-1.086 5.5L15 15.414L12.086 12.5l1.414-1.414l1.5 1.5l3.5-3.5l1.414 1.414ZM11 17H5v-2h6v2Zm0-8H5V7h6v2Zm0 4H5v-2h6v2Z"/>`),
		g.Group(children),
	)
}