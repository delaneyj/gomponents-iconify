package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 20a10 10 0 1 1 0-20a10 10 0 0 1 0 20zm-5.6-4.29a9.95 9.95 0 0 1 11.2 0a8 8 0 1 0-11.2 0zm6.12-7.64l3.02-3.02l1.41 1.41l-3.02 3.02a2 2 0 1 1-1.41-1.41z"/>`),
		g.Group(children),
	)
}