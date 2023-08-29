package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FinancingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M12 9.927V7a3 3 0 0 1 3-3h26a3 3 0 0 1 3 3v26a3 3 0 0 1-3 3h-2.983"/><rect width="34" height="34" x="4" y="10" stroke-linejoin="round" rx="3"/><path stroke-linecap="round" stroke-linejoin="round" d="m15 17l6 6l6-6m-13 8h14m-14 6h14m-7-6v11"/></g>`),
		g.Group(children),
	)
}