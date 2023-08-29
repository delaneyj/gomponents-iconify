package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionAdjustmentTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m18 10l6-6m0 0l6 6m-6-6v10m-6 24l6 6m0 0l6-6m-6 6V34m14-16l6 6m0 0l-6 6m6-6H34m-24-6l-6 6m0 0l6 6m-6-6h10"/><circle cx="24" cy="24" r="4"/></g>`),
		g.Group(children),
	)
}