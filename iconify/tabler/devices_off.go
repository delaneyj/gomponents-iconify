package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 9a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v8m-1 3h-6a1 1 0 0 1-1-1v-6m5-5V5a1 1 0 0 0-1-1H8M4 4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h9m3-9h2M3 3l18 18"/>`),
		g.Group(children),
	)
}