package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireHydrantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21h14m-2 0v-4m2-2v-2a1 1 0 0 0-1-1h-1V8a5 5 0 0 0-8.533-3.538M7.08 7.1A5.03 5.03 0 0 0 7 8v4H6a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h1v5m5-9a2 2 0 1 0 2 2M6 8h2m4 0h6M3 3l18 18"/>`),
		g.Group(children),
	)
}