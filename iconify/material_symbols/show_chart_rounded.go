package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.725 17.75Q2.4 17.425 2.4 17t.325-.75l6.05-6.05q.3-.3.7-.3t.7.3l3.3 3.3l6.4-7.225q.275-.325.713-.325t.737.3q.275.275.288.662t-.263.688l-7.175 8.1q-.275.325-.713.338t-.737-.288l-3.25-3.25l-5.25 5.25q-.325.325-.75.325t-.75-.325Z"/>`),
		g.Group(children),
	)
}