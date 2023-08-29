package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 22q-2.7 0-4.6-1.9T7 15.5v-.575q-2.15-.35-3.575-2.013T2 9V4q0-.425.288-.713T3 3h2q0-.425.288-.713T6 2q.425 0 .713.288T7 3v2q0 .425-.288.713T6 6q-.425 0-.713-.288T5 5H4v4q0 1.65 1.175 2.825T8 13q1.65 0 2.825-1.175T12 9V5h-1q0 .425-.288.713T10 6q-.425 0-.713-.288T9 5V3q0-.425.288-.713T10 2q.425 0 .713.288T11 3h2q.425 0 .713.288T14 4v5q0 2.25-1.425 3.913T9 14.925v.575q0 1.875 1.313 3.188T13.5 20q1.875 0 3.188-1.313T18 15.5v-1.675q-.875-.325-1.438-1.088T16 11q0-1.25.875-2.125T19 8q1.25 0 2.125.875T22 11q0 .975-.563 1.738T20 13.825V15.5q0 2.7-1.9 4.6T13.5 22Z"/>`),
		g.Group(children),
	)
}