package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Layer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M28.135 21.154a3.5 3.5 0 0 0-2.668 1.235L.832 51.404a3.5 3.5 0 0 0 2.67 5.766l93-.065a3.5 3.5 0 0 0 2.666-5.765L74.59 22.389a3.5 3.5 0 0 0-2.668-1.235H28.135zM89.91 62.11l-9.178.008l8.211 9.67l-77.875.053l8.22-9.682l-9.188.008L.832 73.08a3.5 3.5 0 0 0 2.67 5.766l93-.065a3.5 3.5 0 0 0 2.666-5.765L89.91 62.109z" color="currentColor"/>`),
		g.Group(children),
	)
}