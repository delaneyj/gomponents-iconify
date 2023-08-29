package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 17l4 4l4-4"/><path d="M7 21V10a3 3 0 0 1 3-3h11"/><path d="m17 11l4-4l-4-4"/></g>`),
		g.Group(children),
	)
}