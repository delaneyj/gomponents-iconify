package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analysis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16a1 1 0 0 1-.707-1.707l7-7a1 1 0 0 1 1.414 1.414l-7 7A.997.997 0 0 1 14 16zM3 17a1 1 0 0 1-.707-1.707l6-6a1 1 0 0 1 1.414 1.414l-6 6A.997.997 0 0 1 3 17z"/><path fill="currentColor" d="M14 16a.997.997 0 0 1-.707-.293l-5-5a1 1 0 0 1 1.414-1.414l5 5A1 1 0 0 1 14 16Z"/>`),
		g.Group(children),
	)
}