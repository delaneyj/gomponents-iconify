package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveModeratorRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.85 16.05L6.8 3.95l4.5-1.675q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 1.275-.288 2.525t-.862 2.425Zm-2.3 3.3q-.825.875-1.813 1.475t-2.112.975q-.15.05-.3.075T12 21.9q-.1 0-.625-.1Q8 20.675 6 17.637T4 11.1V6.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.55-2.55Z"/>`),
		g.Group(children),
	)
}