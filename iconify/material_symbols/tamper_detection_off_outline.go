package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TamperDetectionOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 17.5l-4-4v1.7l-2-2V6H8.8l-2-2H16q.825 0 1.413.588T18 6v4.5l4-4v11ZM2 9V6q0-.825.588-1.413T4 4l2 2H4v3H2Zm12 11v-2h2v-2l2 2q0 .825-.588 1.413T16 20h-2Zm6.55 3.35L.65 3.45l1.4-1.4l19.9 19.9l-1.4 1.4ZM11.5 11.5Zm.9-1.9ZM4.45 23q-.425 0-.788-.163t-.637-.437L0 19.4l.35-.35q.2-.2.475-.325T1.4 18.6q.3 0 .575.113t.475.337l.55.55v-7.35q0-.325.225-.537t.525-.213q.325 0 .537.213t.213.537V16h1v-5.25q0-.325.225-.537T6.25 10q.325 0 .537.213T7 10.75V16h1v-4.25q0-.325.225-.537T8.75 11q.325 0 .537.213t.213.537V16h1v-3.25q0-.325.225-.537T11.25 12q.325 0 .537.213t.213.537V21q0 .825-.575 1.413T10 23H4.45Z"/>`),
		g.Group(children),
	)
}