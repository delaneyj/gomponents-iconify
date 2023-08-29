package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M6 16h8v16H6z"/><path stroke-linecap="round" d="M10 6v10m0 16v10"/><path d="M34 16h8v16h-8z"/><path stroke-linecap="round" d="M38 6v10m0 16v10"/><path d="M20 14h8v16h-8z"/><path stroke-linecap="round" d="M24 4v10m0 16v10"/></g>`),
		g.Group(children),
	)
}