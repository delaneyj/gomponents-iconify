package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3.055A9.001 9.001 0 1 0 20.945 13H11V3.055Zm2 0V11h7.945A9.004 9.004 0 0 0 13 3.055ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Z"/>`),
		g.Group(children),
	)
}