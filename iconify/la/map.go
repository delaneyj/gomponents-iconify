package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Map(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28 4.469l-1.406.625l-6.625 2.843l-7.625-2.875l-.375-.125l-.375.157l-7 3l-.594.25V27.53l1.406-.625l6.625-2.843l7.625 2.875l.375.125l.375-.157l7-3l.594-.25zM13 7.437l6 2.25v14.876l-6-2.25zM11 7.5v14.844L6 24.5V9.656zm15 0v14.844L21 24.5V9.656z"/>`),
		g.Group(children),
	)
}