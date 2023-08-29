package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Currency(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12a7 7 0 1 0 14 0a7 7 0 1 0-14 0M4 4l3 3m13-3l-3 3M4 20l3-3m13 3l-3-3"/>`),
		g.Group(children),
	)
}