package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretSquareLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm10.781 2.281l-6 6l-.687.719l.687.719l6 6l1.438-1.438L13.937 16l5.282-5.281z"/>`),
		g.Group(children),
	)
}