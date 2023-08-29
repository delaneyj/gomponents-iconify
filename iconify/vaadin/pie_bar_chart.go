package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieBarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 11h3v5H5v-5zm-4 3h3v2H1v-2zm12-2h3v4h-3v-4zM9 9h3v7H9V9zM5 0a5 5 0 1 0 .001 10.001A5 5 0 0 0 5 0zm0 9a4 4 0 0 1 0-8v4h4a4 4 0 0 1-4 4z"/>`),
		g.Group(children),
	)
}