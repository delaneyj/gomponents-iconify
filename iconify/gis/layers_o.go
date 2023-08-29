package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M28.135 10.357a3.5 3.5 0 0 0-2.668 1.235L.832 40.607a3.5 3.5 0 0 0 2.67 5.766l93-.064a3.5 3.5 0 0 0 2.666-5.766L74.59 11.592a3.5 3.5 0 0 0-2.668-1.235zm1.619 7H70.3l18.64 21.957l-77.873.053zM89.91 51.313l-9.178.007l8.211 9.67l-77.875.053l8.22-9.682l-9.188.008L.832 62.283a3.5 3.5 0 0 0 2.67 5.766l93-.065a3.5 3.5 0 0 0 2.666-5.765zm0 21.593l-9.178.008l8.211 9.67l-77.875.053l8.22-9.682l-9.188.008L.832 83.877a3.5 3.5 0 0 0 2.67 5.766l93-.065a3.5 3.5 0 0 0 2.666-5.766z" color="currentColor"/>`),
		g.Group(children),
	)
}