package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PivotTable(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M42 4H6a2 2 0 0 0-2 2l.001 36a2 2 0 0 0 2 2h36a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M20.009 34.008H33.5v-14"/><path stroke-linecap="round" stroke-linejoin="round" d="M24.5 38.5L23 37l-3-3l3-3l1.5-1.5m4.5-5l1.5-1.5l3-3l3 3l1.5 1.5"/><path stroke-linecap="round" d="M14 4v40M4 14.038L44 14"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 4h20M8 44h20"/><path stroke-linecap="round" d="M44 8v10M4 8v10"/></g>`),
		g.Group(children),
	)
}