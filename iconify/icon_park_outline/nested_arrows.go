package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestedArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M30 17V4H4v26h13"/><path d="M43 43V17H17v26h26ZM33 30H17m16 0l-5-5l5 5Zm0 0l-5 5l5-5ZM17 17v26"/></g>`),
		g.Group(children),
	)
}