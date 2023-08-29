package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CakeFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M9 27h30l-4.688 17H13.689L9 27Zm30 0H9c0-5.5 5.5-8 5.5-8c0-2 2-8 9.5-8s9.5 6 9.5 8c0 0 5.5 2.5 5.5 8Z"/><path stroke-linecap="round" d="M6 27h36"/><path d="M28 12a4 4 0 0 0-8 0"/><path stroke-linecap="round" d="m24 8l4-4"/><path d="M20 27v17m8-17v17"/></g>`),
		g.Group(children),
	)
}