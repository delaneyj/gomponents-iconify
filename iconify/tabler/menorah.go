package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Menorah(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4v16M8 4v2a4 4 0 1 0 8 0V4"/><path d="M4 4v2a8 8 0 1 0 16 0V4M10 20h4"/></g>`),
		g.Group(children),
	)
}