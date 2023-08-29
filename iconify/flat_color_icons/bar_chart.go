package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#00BCD4" d="M19 22h10v20H19zM6 12h10v30H6zm26-6h10v36H32z"/>`),
		g.Group(children),
	)
}