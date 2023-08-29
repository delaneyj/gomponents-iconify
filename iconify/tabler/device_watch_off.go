package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceWatchOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 6h5a3 3 0 0 1 3 3v5m-.89 3.132A2.99 2.99 0 0 1 15 18H9a3 3 0 0 1-3-3V9c0-.817.327-1.559.857-2.1"/><path d="M9 18v3h6v-3M9 5V3h6v3M3 3l18 18"/></g>`),
		g.Group(children),
	)
}