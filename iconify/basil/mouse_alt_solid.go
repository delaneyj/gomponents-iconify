package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.014 2.654a6.5 6.5 0 0 1 5.486 6.42v5.852a6.5 6.5 0 0 1-13 0V9.074a6.5 6.5 0 0 1 7.514-6.42ZM11.25 10a.75.75 0 0 0 1.5 0V7a.75.75 0 0 0-1.5 0v3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}