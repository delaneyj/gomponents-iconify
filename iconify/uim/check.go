package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.84 17.08a.997.997 0 0 1-.707-.293l-3.84-3.84a1 1 0 1 1 1.414-1.414l3.133 3.133l7.453-7.453a1 1 0 0 1 1.414 1.414l-8.16 8.16a.997.997 0 0 1-.707.293Z"/>`),
		g.Group(children),
	)
}