package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EllipsisH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 6h4v4H0V6zm6 0h4v4H6V6zm6 0h4v4h-4V6z"/>`),
		g.Group(children),
	)
}