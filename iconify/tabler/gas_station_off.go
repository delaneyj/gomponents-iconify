package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasStationOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 11a2 2 0 0 1 2 2m3 3V9l-3-3M4 20V6c0-.548.22-1.044.577-1.405M8 4h4a2 2 0 0 1 2 2v4m0 4v6M3 20h12"/><path d="M18 7v1a1 1 0 0 0 1 1h1M4 11h7M3 3l18 18"/></g>`),
		g.Group(children),
	)
}