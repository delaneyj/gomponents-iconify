package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarbudsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21q-2.075 0-3.538-1.463T3 16V6.2q0-1.35.825-2.275T6 3q1.35 0 2.175.825T9 6q0 1.275-.863 2.138T6 9H5v7q0 1.25.875 2.125T8 19q1.25 0 2.125-.875T11 16V8q0-2.075 1.463-3.538T16 3q2.075 0 3.538 1.463T21 8v10q0 1.275-.963 2.138T17.8 21q-1.275 0-2.038-.863T15 18q0-1.275.863-2.138T18 15h1V8q0-1.25-.875-2.125T16 5q-1.25 0-2.125.875T13 8v8q0 2.075-1.463 3.538T8 21ZM5 7h1q.425 0 .713-.288T7 6q0-.425-.288-.713T6 5q-.425 0-.713.288T5 6v1Zm13 12q.425 0 .713-.288T19 18v-1h-1q-.425 0-.713.288T17 18q0 .425.288.713T18 19Zm0-1ZM6 6Z"/>`),
		g.Group(children),
	)
}