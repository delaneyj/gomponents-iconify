package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheveronLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.05 9.293L6.343 10L12 15.657l1.414-1.414L9.172 10l4.242-4.243L12 4.343z"/>`),
		g.Group(children),
	)
}