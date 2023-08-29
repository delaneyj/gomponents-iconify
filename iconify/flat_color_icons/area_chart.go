package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#3F51B5" d="M42 37H6V25l10-15l14 7L42 6z"/><path fill="#00BCD4" d="M42 42H6V32l10-8l14 2l12-9z"/>`),
		g.Group(children),
	)
}