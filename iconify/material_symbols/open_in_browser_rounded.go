package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInBrowserRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21h-3q-.425 0-.713-.288T15 20q0-.425.288-.713T16 19h3V7H5v12h3q.425 0 .713.288T9 20q0 .425-.288.713T8 21H5Zm6-1v-5.15l-.875.875q-.3.3-.713.288T8.7 15.7q-.275-.3-.287-.7t.287-.7l2.6-2.6q.15-.15.325-.212t.375-.063q.2 0 .375.063t.325.212l2.6 2.6q.3.3.287.7t-.287.7q-.3.3-.713.313t-.712-.288L13 14.85V20q0 .425-.288.713T12 21q-.425 0-.713-.288T11 20Z"/>`),
		g.Group(children),
	)
}