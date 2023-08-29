package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoveModeratorOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.85 16.05l-1.5-1.55q.3-.8.475-1.663T18 11.1V6.375l-6-2.25L8.35 5.5L6.8 3.95l4.5-1.675q.35-.125.7-.125t.7.125l6 2.25q.575.225.938.725T20 6.375V11.1q0 1.275-.288 2.525t-.862 2.425ZM12 21.9q-.1 0-.625-.1Q8 20.675 6 17.637T4 11.1V6.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.55-2.55q-.825.875-1.813 1.475t-2.112.975q-.15.05-.3.075T12 21.9Zm-1.425-8.525ZM12.85 10ZM12 19.9q.875-.275 1.675-.775t1.475-1.175L6 8.8v2.3q0 3.025 1.7 5.5t4.3 3.3Z"/>`),
		g.Group(children),
	)
}