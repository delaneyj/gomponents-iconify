package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDoubleLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.25 16a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.414 1.414L14.664 12l2.293 2.293A1 1 0 0 1 16.25 16zm-5.5 0a.997.997 0 0 1-.707-.293l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.414 1.414L9.164 12l2.293 2.293A1 1 0 0 1 10.75 16z"/>`),
		g.Group(children),
	)
}