package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.05 17l4.15-2.65q.45-.3.45-.85t-.45-.85L11.05 10q-.5-.325-1.025-.05t-.525.875v5.35q0 .6.525.875T11.05 17ZM4 21q-.825 0-1.413-.588T2 19V7q0-.425.288-.713T3 6h5V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v2h5q.425 0 .713.288T22 7v12q0 .825-.588 1.413T20 21H4Zm6-15h4V4h-4v2Z"/>`),
		g.Group(children),
	)
}