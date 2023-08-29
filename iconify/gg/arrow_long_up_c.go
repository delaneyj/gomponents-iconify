package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongUpC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.758 5.23l4.274-4.21l4.21 4.275l-1.424 1.403l-1.804-1.83l-.071 12.287a3.001 3.001 0 0 1-1.029 5.825a3 3 0 0 1-.971-5.835l.071-12.315l-1.853 1.826L7.758 5.23Zm4.175 13.75a1 1 0 1 0-.01 2a1 1 0 0 0 .01-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}