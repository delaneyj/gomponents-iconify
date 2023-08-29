package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditBinocularBinocularBinocularsViewZoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="3" cy="11" r="2.5"/><path d="M5.5 3V1.5a1 1 0 0 0-1-1H3.32a1 1 0 0 0-1 .81L.55 10.5m3.95-5h5M5.5 8v3"/><circle cx="11" cy="11" r="2.5"/><path d="M8.5 3V1.5a1 1 0 0 1 1-1h1.18a1 1 0 0 1 1 .81l1.79 9.19M8.5 8v3"/></g>`),
		g.Group(children),
	)
}