package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M28 3a1 1 0 0 1 1 1v10a1 1 0 1 1-2 0V6.414L6.414 27H14a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1V18a1 1 0 1 1 2 0v7.586L25.586 5H18a1 1 0 1 1 0-2h10Z"/>`),
		g.Group(children),
	)
}