package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TamperDetectionOnOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17v-2h2V3H4v3H2V3q0-.825.588-1.413T4 1h12q.825 0 1.413.588T18 3v4.5l3.15-3.15q.25-.25.55-.125t.3.475v8.575q0 .35-.3.475t-.55-.125L18 10.5V15q0 .825-.588 1.413T16 17h-2Zm-9.55 3q-.425 0-.788-.163t-.637-.437L.7 17.075q-.275-.275-.275-.687T.7 15.7q.3-.3.713-.3t.687.3l.9.9V9.25q0-.325.225-.537T3.75 8.5q.325 0 .537.213t.213.537V13h1V7.75q0-.325.225-.537T6.25 7q.325 0 .537.213T7 7.75V13h1V8.75q0-.325.225-.537T8.75 8q.325 0 .537.213t.213.537V13h1V9.75q0-.325.225-.537T11.25 9q.325 0 .537.213T12 9.75V18q0 .825-.575 1.413T10 20H4.45ZM16 3v12V3Z"/>`),
		g.Group(children),
	)
}