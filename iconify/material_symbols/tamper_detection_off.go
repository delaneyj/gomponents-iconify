package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TamperDetectionOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.55 23.35l-3.6-3.6q-.2.125-.438.188T16 20h-2v-3.2L6.2 9H2V6q0-.275.063-.513t.187-.437l-1.6-1.6l1.4-1.4l19.9 19.9l-1.4 1.4ZM22 17.5l-4-4v1.7L6.8 4H16q.825 0 1.413.588T18 6v4.5l4-4v11ZM4.45 23q-.425 0-.788-.163t-.637-.437L0 19.4l.35-.35q.2-.2.475-.325T1.4 18.6q.3 0 .575.113t.475.337l.55.55v-7.35q0-.325.225-.537t.525-.213q.325 0 .537.213t.213.537V16h1v-5.25q0-.325.225-.537T6.25 10q.325 0 .537.213T7 10.75V16h1v-4.25q0-.325.225-.537T8.75 11q.325 0 .537.213t.213.537V16h1v-3.25q0-.325.225-.537T11.25 12q.325 0 .537.213t.213.537V21q0 .825-.575 1.413T10 23H4.45Z"/>`),
		g.Group(children),
	)
}