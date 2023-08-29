package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularNodataOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19.425L17.6 20.8q-.275.275-.688.288T16.2 20.8q-.275-.275-.275-.7t.275-.7l1.4-1.4l-1.4-1.4q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.4 1.4l1.4-1.4q.275-.275.687-.287t.713.287q.275.275.275.7t-.275.7L20.425 18l1.375 1.4q.275.275.288.688t-.288.712q-.275.275-.7.275t-.7-.275L19 19.425ZM3.7 20.3L20.275 3.725q.475-.475 1.088-.212t.612.937v8.25q-.45-.275-.95-.438T19.975 12V6.85L6.825 20h6.425q.2.575.5 1.075t.675.925h-10q-.675 0-.937-.613T3.7 20.3Zm3.125-.3l13.15-13.15l-3.438 3.438l-3.024 3.024l-3.088 3.088l-3.6 3.6Z"/>`),
		g.Group(children),
	)
}