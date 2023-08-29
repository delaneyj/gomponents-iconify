package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.667 12H2v8H0V0h12l.333 2H20l-3 6l3 6H8l-.333-2z"/>`),
		g.Group(children),
	)
}