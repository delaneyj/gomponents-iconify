package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 0h1v12h-1V0Zm-3 3v9H9V3h1ZM6 6v6H5V6h1Zm-5 6V9h1v3H1Zm14 3H0v-1h15v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}