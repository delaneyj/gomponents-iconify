package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EaseOutControlPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21S13 5 21 5M7 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0zm0 0h2m5 0h-2"/>`),
		g.Group(children),
	)
}