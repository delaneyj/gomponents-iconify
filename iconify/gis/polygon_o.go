package gis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolygonO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M33.162 19.463a3.5 3.5 0 0 0-3.111 1.974L9.34 64.239a3.5 3.5 0 0 0 2.632 4.985l75.015 11.275a3.5 3.5 0 0 0 3.846-4.55L76.566 32.456a3.5 3.5 0 0 0-2.431-2.293l-40.04-10.586a3.5 3.5 0 0 0-.933-.115zm1.934 7.621l35.412 9.361l11.904 36.287l-64.7-9.724l17.384-35.924z" color="currentColor"/>`),
		g.Group(children),
	)
}