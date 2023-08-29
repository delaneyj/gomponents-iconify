package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableLampRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 21q-.425 0-.713-.288T13 20q0-.425.288-.713T14 19h6q.425 0 .713.288T21 20q0 .425-.288.713T20 21h-6Zm3-3q-.425 0-.713-.288T16 17V8q0-.425-.288-.713T15 7h-4v3q0 .425-.288.713T10 11H4q-.55 0-.85-.45t-.075-.95L5.45 4.2q.25-.55.737-.875T7.275 3H9q.825 0 1.413.588T11 5h4q1.25 0 2.125.875T18 8v9q0 .425-.288.713T17 18Z"/>`),
		g.Group(children),
	)
}