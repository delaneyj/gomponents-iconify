package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0m10.585-1.414a2 2 0 0 0 2.826 2.831"/><path d="M5 17H3V6a1 1 0 0 1 1-1h1m3.96 0H13v4m0 4v4m-4 0h6m6 0v-6h-6m-2-5h5l3 5M3 3l18 18"/></g>`),
		g.Group(children),
	)
}