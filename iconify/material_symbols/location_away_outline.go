package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationAwayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.575 12.25ZM2 21V9l8-6l5.375 4.05q-.625.075-1.175.288t-1.05.562L10 5.5L4 10v9h4v2H2Zm8 0v-1.9q0-.525.263-.988t.712-.737q1.15-.675 2.413-1.025T16 16q1.35 0 2.613.35t2.412 1.025q.45.275.713.738T22 19.1V21H10Zm2.15-2h7.7q-.875-.5-1.85-.75T16 18q-1.025 0-2 .25t-1.85.75ZM16 15q-1.25 0-2.125-.875T13 12q0-1.25.875-2.125T16 9q1.25 0 2.125.875T19 12q0 1.25-.875 2.125T16 15Zm0-2q.425 0 .713-.288T17 12q0-.425-.288-.713T16 11q-.425 0-.713.288T15 12q0 .425.288.713T16 13Zm0 6Z"/>`),
		g.Group(children),
	)
}