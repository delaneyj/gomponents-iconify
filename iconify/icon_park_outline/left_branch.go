package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftBranch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M26 8c1.5-.012 5.5 0 6.571 5.062C33.655 18.178 38.143 22.848 40 24M26 40c1.5 0 5.5 0 6.571-5.062C33.655 29.821 38.143 25.152 40 24"/><circle r="4" fill="currentColor" transform="matrix(-1 0 0 1 40 24)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M40 24H26m-8 0H6M18 8H6m12 32H6"/></g>`),
		g.Group(children),
	)
}