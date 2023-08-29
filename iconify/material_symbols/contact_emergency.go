package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactEmergency(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.25 12h1.5v-1.7l1.475.85l.75-1.3L19.5 9l1.475-.85l-.75-1.3l-1.475.85V6h-1.5v1.7l-1.475-.85l-.75 1.3L16.5 9l-1.475.85l.75 1.3l1.475-.85V12ZM2 21q-.825 0-1.413-.588T0 19V5q0-.825.588-1.413T2 3h20q.825 0 1.413.588T24 5v14q0 .825-.588 1.413T22 21H2Zm7-7q1.25 0 2.125-.875T12 11q0-1.25-.875-2.125T9 8q-1.25 0-2.125.875T6 11q0 1.25.875 2.125T9 14Zm-6.9 5h13.8q-1.05-1.875-2.9-2.938T9 15q-2.15 0-4 1.063T2.1 19Z"/>`),
		g.Group(children),
	)
}