package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.5 0H6v9h9v-.5A8.5 8.5 0 0 0 6.5 0Z"/><path fill="currentColor" d="M12.826 10H5V2.174A6.5 6.5 0 1 0 12.826 10Z"/>`),
		g.Group(children),
	)
}