package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.032 1.017l-4.274 4.21L9.16 6.653l1.859-1.83l-.056 18.155l2 .006l.056-18.113l1.798 1.825l1.425-1.403l-4.21-4.275Z"/>`),
		g.Group(children),
	)
}