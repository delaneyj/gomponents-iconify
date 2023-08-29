package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationHomeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 3l8 6v12H4V9l8-6Zm0 12q1.25 0 2.125-.875T15 12q0-1.25-.875-2.125T12 9q-1.25 0-2.125.875T9 12q0 1.25.875 2.125T12 15Zm0-2q-.425 0-.713-.288T11 12q0-.425.288-.713T12 11q.425 0 .713.288T13 12q0 .425-.288.713T12 13Zm0 5q-1.025 0-2 .25T8.15 19h7.7q-.875-.5-1.85-.75T12 18Zm-6-8v8q1.3-.975 2.825-1.488T12 16q1.65 0 3.175.513T18 18v-8l-6-4.5L6 10Zm6 2Z"/>`),
		g.Group(children),
	)
}