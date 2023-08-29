package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 4.594L18.281 6.28l-9 9l-.687.719l.687.719l9 9L20 27.406zm-2 4.843v13.126L11.437 16z"/>`),
		g.Group(children),
	)
}