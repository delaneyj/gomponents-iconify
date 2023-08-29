package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerThreeSelectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 19H5.5q-.625 0-1.063-.438T4 17.5q0-.625.438-1.063T5.5 16H10v-2.5H5.5q-.625 0-1.063-.438T4 12q0-.625.438-1.063T5.5 10.5H10V8H5.5q-.625 0-1.063-.438T4 6.5q0-.625.438-1.063T5.5 5H10q1.25 0 2.125.875T13 8v1.9q0 .875-.613 1.488T10.9 12q.875 0 1.488.613T13 14.1V16q0 1.25-.875 2.125T10 19Zm9.65 0H16q-.425 0-.713-.288T15 18q0-.425.288-.713T16 17h3v-1h-2.65q-.575 0-.963-.388T15 14.65v-2.3q0-.575.388-.963T16.35 11H20q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-3v1h2.65q.575 0 .963.388t.387.962v2.3q0 .575-.388.963T19.65 19Z"/>`),
		g.Group(children),
	)
}