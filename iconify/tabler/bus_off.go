package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 17a2 2 0 1 0 4 0a2 2 0 1 0-4 0m12.18-.828a2 2 0 0 0 2.652 2.648"/><path d="M4 17H2V6a1 1 0 0 1 1-1h2m4 0h8c2.761 0 5 3.134 5 7v5h-1m-5 0H8"/><path d="m16 5l1.5 7H22M2 10h8m4 0h3M7 7v3m5-5v3M3 3l18 18"/></g>`),
		g.Group(children),
	)
}