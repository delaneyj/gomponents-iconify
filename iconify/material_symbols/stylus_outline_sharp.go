package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StylusOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.95 20L2.7 21.3L4 15.05L8.95 20Zm0 0L4 15.05L16.875 2.175l4.95 4.95L8.95 20Zm-.3-2.525L19 7.125L16.875 5L6.525 15.35l2.125 2.125Z"/>`),
		g.Group(children),
	)
}