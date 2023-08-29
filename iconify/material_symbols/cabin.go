package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cabin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19h12v-2H6v2Zm0-4h12v-2H6v2Zm0-4h12v-.9L16.55 9h-9.1L6 10.1v.9Zm4.05-4h3.9L12 5.525L10.05 7ZM4 21v-9.375L2.2 13L1 11.4l3-2.3V6h2v1.575L12 3l11 8.4l-1.2 1.575l-1.8-1.35V21H4ZM4 5q0-1.25.875-2.125T7 2q.425 0 .713-.288T8 1h2q0 1.25-.875 2.125T7 4q-.425 0-.713.288T6 5H4Z"/>`),
		g.Group(children),
	)
}