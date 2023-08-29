package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Windy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 5a1 1 0 1 0 0 2h9v2h-9a3 3 0 1 1 0-6h4v2h-4ZM5.5 8a1.5 1.5 0 1 0 0 3H22v2H5.5a3.5 3.5 0 1 1 0-7H8v2H5.5ZM5 18a3 3 0 0 1 3-3h10v2H8a1 1 0 1 0 0 2h4.5v2H8a3 3 0 0 1-3-3Z"/>`),
		g.Group(children),
	)
}