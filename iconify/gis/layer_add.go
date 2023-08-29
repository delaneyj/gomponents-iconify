package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayerAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M80.5 0c-7.145 0-13.252 4.246-15.977 10.357H28.135a3.5 3.5 0 0 0-2.668 1.235L.832 40.607a3.5 3.5 0 0 0 2.67 5.766l93-.064a3.5 3.5 0 0 0 2.666-5.766l-7.87-9.272A17.428 17.428 0 0 0 98 17.5C98 7.805 90.195 0 80.5 0zM77 5h7v9h9v7h-9v9h-7v-9h-9v-7h9V5zm12.91 46.313l-9.178.007l8.211 9.67l-77.875.053l8.22-9.682l-9.188.008L.832 62.283a3.5 3.5 0 0 0 2.67 5.766l93-.065a3.5 3.5 0 0 0 2.666-5.765L89.91 51.312zm0 21.593l-9.178.008l8.211 9.67l-77.875.053l8.22-9.682l-9.188.008L.832 83.877a3.5 3.5 0 0 0 2.67 5.766l93-.065a3.5 3.5 0 0 0 2.666-5.766L89.91 72.906z" color="currentColor"/>`),
		g.Group(children),
	)
}