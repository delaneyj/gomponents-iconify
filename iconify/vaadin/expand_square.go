package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M11 2H2v9l1-1V3h7zM5 14h9V5l-1 1v7H6z"/><path fill="currentColor" d="M16 0h-5l1.8 1.8l-4.5 4.5l1.4 1.4l4.5-4.5L16 5zM7.7 9.7L6.3 8.3l-4.5 4.5L0 11v5h5l-1.8-1.8z"/>`),
		g.Group(children),
	)
}