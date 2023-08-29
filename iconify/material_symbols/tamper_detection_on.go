package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TamperDetectionOn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.45 20q-.425 0-.788-.163t-.637-.437L0 16.4l.35-.35q.225-.225.5-.338t.55-.112q.3 0 .575.113t.475.337l.55.55V9.25q0-.325.225-.537T3.75 8.5q.325 0 .537.213t.213.537V13h1V7.75q0-.325.225-.537T6.25 7q.325 0 .537.213T7 7.75V13h1V8.75q0-.325.225-.537T8.75 8q.325 0 .537.213t.213.537V13h1V9.75q0-.325.225-.537T11.25 9q.325 0 .537.213T12 9.75V18q0 .825-.575 1.413T10 20H4.45ZM14 17V8q0-.825-.588-1.413T12 6H2V3q0-.825.588-1.413T4 1h12q.825 0 1.413.588T18 3v4.5l4-4v11l-4-4V15q0 .825-.588 1.413T16 17h-2Z"/>`),
		g.Group(children),
	)
}