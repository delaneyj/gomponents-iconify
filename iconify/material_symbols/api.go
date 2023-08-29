package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Api(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 14l-2-2l2-2l2 2l-2 2ZM9.875 8.125l-2.5-2.5L12 1l4.625 4.625l-2.5 2.5L12 6L9.875 8.125Zm-4.25 8.5L1 12l4.625-4.625l2.5 2.5L6 12l2.125 2.125l-2.5 2.5Zm12.75 0l-2.5-2.5L18 12l-2.125-2.125l2.5-2.5L23 12l-4.625 4.625ZM12 23l-4.625-4.625l2.5-2.5L12 18l2.125-2.125l2.5 2.5L12 23Z"/>`),
		g.Group(children),
	)
}