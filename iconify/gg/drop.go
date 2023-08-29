package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M6.343 19.52a8 8 0 0 1 0-11.313L12 2.55l5.657 5.657A8 8 0 0 1 6.343 19.521Z"/>`),
		g.Group(children),
	)
}