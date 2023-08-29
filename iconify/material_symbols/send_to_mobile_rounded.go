package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendToMobileRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.15 13H13q-.425 0-.713-.288T12 12q0-.425.288-.713T13 11h4.15l-.875-.9Q16 9.825 16 9.412t.3-.712q.275-.275.7-.275t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.688.288T16.3 15.3q-.275-.275-.288-.688t.263-.712l.875-.9ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v12h10v-1h2v4q0 .825-.588 1.413T17 23H7Z"/>`),
		g.Group(children),
	)
}