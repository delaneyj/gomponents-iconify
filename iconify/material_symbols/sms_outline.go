package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11q.425 0 .713-.288T9 10q0-.425-.288-.713T8 9q-.425 0-.713.288T7 10q0 .425.288.713T8 11Zm4 0q.425 0 .713-.288T13 10q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10q0 .425.288.713T12 11Zm4 0q.425 0 .713-.288T17 10q0-.425-.288-.713T16 9q-.425 0-.713.288T15 10q0 .425.288.713T16 11ZM2 22V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6l-4 4Zm3.15-6H20V4H4v13.125L5.15 16ZM4 16V4v12Z"/>`),
		g.Group(children),
	)
}