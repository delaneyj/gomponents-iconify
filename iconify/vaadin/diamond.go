package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 6h4l3 8.6L0 6zm16 0h-4l-3 8.6L16 6zm-8 9L5 6h6l-3 9zM4 5H0l2-3l2 3zm12 0h-4l2-3l2 3zm-6 0H6l2-3l2 3zM3.34 2H7L5 5L3.34 2zM9 2h4l-2 3l-2-3z"/>`),
		g.Group(children),
	)
}