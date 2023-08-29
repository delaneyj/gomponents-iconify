package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChartHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22H2V2h2v20Zm11-1H5v-3h10v3Zm3-5H5v-3h13v3Zm3-5H5V8h16v3Zm-8-5H5V3h8v3Z"/>`),
		g.Group(children),
	)
}