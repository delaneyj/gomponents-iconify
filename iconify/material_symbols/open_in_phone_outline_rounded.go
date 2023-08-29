package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInPhoneOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23q-.825 0-1.413-.588T5 21v-6h2v3h10V6H7v3H5V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Zm0-3v1h10v-1H7Zm3.15-7H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h7.15l-.875-.9Q9 9.825 9 9.412t.3-.712q.275-.275.7-.275t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T9.3 15.3q-.275-.275-.288-.688t.263-.712l.875-.9ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}