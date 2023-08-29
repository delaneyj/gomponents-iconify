package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatAddOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9h6q.425 0 .713-.288T15 8q0-.425-.288-.713T14 7H8q-.425 0-.713.288T7 8q0 .425.288.713T8 9Zm0 4h3q.425 0 .713-.288T12 12q0-.425-.288-.713T11 11H8q-.425 0-.713.288T7 12q0 .425.288.713T8 13Zm9 4h-2q-.425 0-.713-.288T14 16q0-.425.288-.713T15 15h2v-2q0-.425.288-.713T18 12q.425 0 .713.288T19 13v2h2q.425 0 .713.288T22 16q0 .425-.288.713T21 17h-2v2q0 .425-.288.713T18 20q-.425 0-.713-.288T17 19v-2ZM6 17l-2.15 2.15q-.25.25-.55.125T3 18.8V5q0-.825.588-1.413T5 3h12q.825 0 1.413.588T19 5v5.075q-.25-.05-.5-.063T18 10q-2.525 0-4.263 1.75T12 16q0 .25.013.5t.062.5H6Z"/>`),
		g.Group(children),
	)
}