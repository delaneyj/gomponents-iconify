package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11 .958v9.039C11 10.551 10.551 11 9.997 11H.958A9.1 9.1 0 1 0 11 .958zm-2 0A9.098 9.098 0 0 0 .958 9H9V.958z"/>`),
		g.Group(children),
	)
}