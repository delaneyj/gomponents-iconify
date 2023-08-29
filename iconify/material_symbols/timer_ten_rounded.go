package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimerTenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 16h3V8h-3v8Zm0 3q-1.25 0-2.125-.875T11 16V8q0-1.25.875-2.125T14 5h3q1.25 0 2.125.875T20 8v8q0 1.25-.875 2.125T17 19h-3ZM6 8h-.5q-.625 0-1.063-.438T4 6.5q0-.625.438-1.063T5.5 5H8q.425 0 .713.288T9 6v11.5q0 .625-.438 1.063T7.5 19q-.625 0-1.063-.438T6 17.5V8Z"/>`),
		g.Group(children),
	)
}