package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatIndividualSuiteOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 17H3q-.825 0-1.413-.588T1 15V8q0-.425.288-.713T2 7q.425 0 .713.288T3 8v7h8V9q0-.825.588-1.413T13 7h6q1.65 0 2.825 1.175T23 11v4q0 .825-.588 1.413T21 17Zm-8-2h8v-4q0-.825-.588-1.413T19 9h-6v6Zm0-6v6v-6Zm-6 5q1.25 0 2.125-.875T10 11q0-1.25-.875-2.125T7 8q-1.25 0-2.125.875T4 11q0 1.25.875 2.125T7 14Zm0-2q-.425 0-.713-.288T6 11q0-.425.288-.713T7 10q.425 0 .713.288T8 11q0 .425-.288.713T7 12Zm0-1Z"/>`),
		g.Group(children),
	)
}