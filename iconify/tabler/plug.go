package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.785 6L18 14.215l-2.054 2.054a5.81 5.81 0 1 1-8.215-8.215L9.785 6zM4 20l3.5-3.5M15 4l-3.5 3.5M20 9l-3.5 3.5"/>`),
		g.Group(children),
	)
}