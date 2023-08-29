package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarbudsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21q-2.075 0-3.538-1.463T3 16V6.2q0-1.35.825-2.275T6 3h3v6H5v7q0 1.25.875 2.125T8 19q1.25 0 2.125-.875T11 16V8q0-2.075 1.463-3.538T16 3q2.075 0 3.538 1.463T21 8v10q0 1.275-.963 2.138T17.8 21H15v-6h4V8q0-1.25-.875-2.125T16 5q-1.25 0-2.125.875T13 8v8q0 2.075-1.463 3.538T8 21ZM5 7h2V5H6q-.425 0-.713.288T5 6v1Zm12 12h1q.425 0 .713-.288T19 18v-1h-2v2ZM6 6Zm12 12Z"/>`),
		g.Group(children),
	)
}