package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1.586 3H21v2H6.414l7 7l-7 7H21v2H1.586l9-9l-9-9Z"/>`),
		g.Group(children),
	)
}