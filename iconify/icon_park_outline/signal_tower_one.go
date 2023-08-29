package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalTowerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="m12 44l8-40h8l8 40"/><path d="M15 10h18"/><path stroke-linejoin="round" d="M12 26h24m-21 1l20 12m-2-12L14 39m16-28L15 26m3-15l15 15"/></g>`),
		g.Group(children),
	)
}