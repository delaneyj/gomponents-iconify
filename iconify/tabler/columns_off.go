package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h2m-2 4h5.5M4 14h5.5M4 18h5.5m5-12H20m-5.5 4H20m-2 4h2m-5.5 4H18M3 3l18 18"/>`),
		g.Group(children),
	)
}