package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafetyCheckOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.55 19.35q-.825.875-1.813 1.475t-2.112.975q-.15.05-.3.075T12 21.9q-.1 0-.625-.1Q8 20.675 6 17.637T4 11.1V6.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.55-2.55ZM12 17q.5 0 .963-.088t.887-.262l-6.5-6.5q-.175.425-.262.888T7 12q0 2.075 1.463 3.538T12 17Zm6.85-.95l-2.2-2.2q.175-.425.263-.887T17 12q0-2.075-1.463-3.538T12 7q-.5 0-.95.088t-.875.262L6.8 3.95l4.5-1.675q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 1.275-.288 2.525t-.862 2.425Z"/>`),
		g.Group(children),
	)
}