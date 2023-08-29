package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartArrowsVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 21V7m-9 8l3-3l3 3m0-5l3-3l3 3M3 21h18m-9 0v-9M3 6l3-3l3 3M6 21V3"/>`),
		g.Group(children),
	)
}