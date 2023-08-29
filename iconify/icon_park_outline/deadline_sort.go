package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeadlineSort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M6 5v25.004h36V5M30 37l-6 6l-6-6m6-7v13"/><path d="M18.344 20.658L29.658 9.344m-11.314-.001l11.313 11.314"/></g>`),
		g.Group(children),
	)
}