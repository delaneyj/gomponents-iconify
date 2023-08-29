package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 14.118l-2.317 2.317a4 4 0 1 1-2.121-2.121l2.317-2.317L4.21 6.329a2 2 0 0 1 0-2.828l.708-.707L12 9.875l7.081-7.081l.708.707a2 2 0 0 1 0 2.828l-5.668 5.668l2.317 2.316a4 4 0 1 1-2.121 2.121L12 14.119Zm-6 5.879a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm12 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/>`),
		g.Group(children),
	)
}