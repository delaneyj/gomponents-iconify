package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChristmasTreeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 5.5L12 3l4 4l-2 1l4 4l-1.5.5M17 17H5l4-4l-3-1l3-3m5 8v3a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-3M3 3l18 18"/>`),
		g.Group(children),
	)
}