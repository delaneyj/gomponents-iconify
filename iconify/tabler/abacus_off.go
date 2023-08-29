package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AbacusOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5v16m14 0v-2m0-4V3M5 7h2m4 0h8M5 15h10m-7-2v4m3-4v4m5-1v1M14 5v4m-3-4v2M8 8v1M3 21h18M3 3l18 18"/>`),
		g.Group(children),
	)
}