package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anticlockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.333 5a3 3 0 0 1 3-3H13v2H7.333a1 1 0 0 0-1 1v6.5H4.34L.586 7.446l1.467-1.36l2.28 2.462V5ZM8 6h15v15H8V6Zm2 2v11h11V8H10Z"/>`),
		g.Group(children),
	)
}