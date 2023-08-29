package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CottageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-9.375L2.2 13L1 11.4l3-2.3V6h2v1.575L12 3l11 8.4l-1.2 1.575l-1.8-1.35V21H4Zm2-2h5v-4h2v4h5v-8.9l-6-4.575L6 10.1V19ZM4 5q0-1.25.875-2.125T7 2q.425 0 .713-.288T8 1h2q0 1.25-.875 2.125T7 4q-.425 0-.713.288T6 5H4Zm2 14h12H6Z"/>`),
		g.Group(children),
	)
}