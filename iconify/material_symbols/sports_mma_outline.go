package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMmaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21q-.425 0-.713-.288T7 20v-3h10v3q0 .425-.288.713T16 21H8ZM19 8v3q0 .075-.05.2L18.2 15q-.075.45-.413.725T17 16H7q-.45 0-.788-.275T5.8 15l-.75-3.8Q5 11.075 5 11V5q0-.825.588-1.413T7 3h8q.825 0 1.413.588T17 5v3q0-.425.288-.713T18 7q.425 0 .713.288T19 8ZM7.65 14h8.7l.65-3.4V10h-2V5H7v5.6l.65 3.4ZM8 10h6V7H8v3Zm4-.5Z"/>`),
		g.Group(children),
	)
}