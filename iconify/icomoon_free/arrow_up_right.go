package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m4.5 0l4 4L0 12.5L3.5 16L12 7.5l4 4V0H4.5z"/>`),
		g.Group(children),
	)
}