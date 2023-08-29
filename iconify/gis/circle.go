package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50 5C25.189 5 5 25.189 5 50s20.189 45 45 45s45-20.189 45-45S74.811 5 50 5z" color="currentColor"/>`),
		g.Group(children),
	)
}