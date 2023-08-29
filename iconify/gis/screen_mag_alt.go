package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenMagAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M2.5 12.5A2.5 2.5 0 0 0 0 15v70a2.5 2.5 0 0 0 2.5 2.5h95A2.5 2.5 0 0 0 100 85V15a2.5 2.5 0 0 0-2.5-2.5h-95zm47.645 15.072a22.36 22.36 0 0 1 15.863 6.59a22.386 22.386 0 0 1 0 31.727a22.384 22.384 0 0 1-31.725 0a22.386 22.386 0 0 1 0-31.727a22.359 22.359 0 0 1 15.862-6.59z" color="currentColor"/>`),
		g.Group(children),
	)
}