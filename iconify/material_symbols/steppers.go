package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steppers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15q-1.25 0-2.125-.875T2 12q0-1.25.875-2.125T5 9q1.25 0 2.125.875T8 12q0 1.25-.875 2.125T5 15Zm0-2q.425 0 .713-.288T6 12q0-.425-.288-.713T5 11q-.425 0-.713.288T4 12q0 .425.288.713T5 13Zm7 2q-1.25 0-2.125-.875T9 12q0-1.25.875-2.125T12 9q1.25 0 2.125.875T15 12q0 1.25-.875 2.125T12 15Zm0-2q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm7 2q-1.25 0-2.125-.875T16 12q0-1.25.875-2.125T19 9q1.25 0 2.125.875T22 12q0 1.25-.875 2.125T19 15Z"/>`),
		g.Group(children),
	)
}