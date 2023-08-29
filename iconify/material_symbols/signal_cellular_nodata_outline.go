package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularNodataOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.9 21.5l-1.4-1.4l2.1-2.1l-2.1-2.1l1.4-1.4l2.1 2.1l2.1-2.1l1.4 1.4l-2.075 2.1l2.075 2.1l-1.4 1.4l-2.1-2.075l-2.1 2.075ZM2 22L21.975 2.025V12.7q-.45-.275-.95-.438T19.975 12V6.85L6.825 20h6.425q.2.575.5 1.075t.675.925H2Zm4.825-2l13.15-13.15l-3.438 3.438l-3.024 3.024l-3.088 3.088l-3.6 3.6Z"/>`),
		g.Group(children),
	)
}