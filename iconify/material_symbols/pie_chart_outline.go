package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11h6.95q-.375-2.75-2.288-4.663T13 4.05V11Zm-2 8.95V4.05q-3.025.375-5.013 2.638T4 12q0 3.05 1.988 5.313T11 19.95Zm2 0q2.75-.35 4.675-2.275T19.95 13H13v6.95ZM12 12Zm0 10q-2.075 0-3.9-.787t-3.175-2.138q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.888.788t3.174 2.15q1.363 1.362 2.15 3.175T22 12q0 2.05-.788 3.875t-2.137 3.188q-1.35 1.362-3.175 2.15T12 22Z"/>`),
		g.Group(children),
	)
}