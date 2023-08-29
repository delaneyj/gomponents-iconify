package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTenSelectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 16h3V8h-3v8Zm0 3q-1.25 0-2.125-.875T7 16V8q0-1.25.875-2.125T10 5h3q1.25 0 2.125.875T16 8v8q0 1.25-.875 2.125T13 19h-3ZM3 8h-.5q-.625 0-1.063-.438T1 6.5q0-.625.438-1.063T2.5 5H6v12.5q0 .625-.438 1.063T4.5 19q-.625 0-1.063-.438T3 17.5V8Zm18.65 11H18q-.425 0-.713-.288T17 18q0-.425.288-.713T18 17h3v-1h-2.65q-.575 0-.963-.388T17 14.65v-2.3q0-.575.388-.963T18.35 11H22q.425 0 .713.288T23 12q0 .425-.288.713T22 13h-3v1h2.65q.575 0 .963.388t.387.962v2.3q0 .575-.388.963T21.65 19Z"/>`),
		g.Group(children),
	)
}