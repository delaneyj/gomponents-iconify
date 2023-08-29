package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineTornado(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.11 8L23 3H1l2.89 5zM7.95 15L12 22l4.05-7zm11-5H5.05l1.74 3h10.42z"/>`),
		g.Group(children),
	)
}