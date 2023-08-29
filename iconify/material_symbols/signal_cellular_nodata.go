package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularNodata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.9 21.5l-1.4-1.4l2.1-2.1l-2.1-2.1l1.4-1.4l2.1 2.1l2.1-2.1l1.4 1.4l-2.075 2.1l2.075 2.1l-1.4 1.4l-2.1-2.075l-2.1 2.075ZM2 22L21.975 2.025V12.7q-.65-.4-1.425-.588T19 11.925q-2.525 0-4.3 1.775T12.925 18q0 1.15.375 2.138T14.425 22H2Z"/>`),
		g.Group(children),
	)
}