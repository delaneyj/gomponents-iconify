package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotsHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><rect width="4" height="4" x="3" y="10" rx="2"/><rect width="4" height="4" x="10" y="10" rx="2"/><rect width="4" height="4" x="17" y="10" rx="2"/></g>`),
		g.Group(children),
	)
}