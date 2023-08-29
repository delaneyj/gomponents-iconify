package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EaseInOutControlPoints(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20a2 2 0 1 0 4 0a2 2 0 0 0-4 0zm0 0h-2M7 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm0 0h2m5 0h-2m0 16h-2m-7 0c8 0 10-16 18-16"/>`),
		g.Group(children),
	)
}