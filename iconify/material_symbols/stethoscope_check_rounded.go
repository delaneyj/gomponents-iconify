package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StethoscopeCheckRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 22q-2.7 0-4.6-1.9T7 15.5v-.575q-2.15-.35-3.575-2.013T2 9V4q0-.425.288-.713T3 3h2q0-.425.288-.713T6 2q.425 0 .713.288T7 3v2q0 .425-.288.713T6 6q-.425 0-.713-.288T5 5H4v4q0 1.65 1.175 2.825T8 13q1.65 0 2.825-1.175T12 9V5h-1q0 .425-.288.713T10 6q-.425 0-.713-.288T9 5V3q0-.425.288-.713T10 2q.425 0 .713.288T11 3h2q.425 0 .713.288T14 4v5q0 2.25-1.425 3.913T9 14.925v.575q0 1.875 1.313 3.188T13.5 20v2Zm3.85-3.825l3.525-3.55q.3-.3.713-.3t.712.3q.3.3.3.713t-.3.712l-4.25 4.25q-.3.3-.713.3t-.712-.3L14.5 18.175q-.275-.3-.275-.713t.3-.712q.3-.3.7-.3t.7.3l1.425 1.425Z"/>`),
		g.Group(children),
	)
}