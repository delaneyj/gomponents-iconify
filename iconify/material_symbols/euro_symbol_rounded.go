package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EuroSymbolRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 21q-2.95 0-5.25-1.675T6.5 15H4q-.425 0-.713-.288T3 14q0-.425.288-.713T4 13h2.05Q6 12.725 6 12.5v-1q0-.225.05-.5H4q-.425 0-.713-.288T3 10q0-.425.288-.713T4 9h2.5q.95-2.65 3.25-4.325T15 3q1.375 0 2.613.375t2.312 1.075q.45.3.488.85t-.388.975q-.3.3-.762.35t-.888-.2q-.75-.425-1.6-.675T15 5.5q-1.875 0-3.412.963T9.25 9H14q.425 0 .713.288T15 10q0 .425-.288.713T14 11H8.6q-.05.275-.075.5T8.5 12q0 .275.025.5t.075.5H14q.425 0 .713.288T15 14q0 .425-.288.713T14 15H9.25q.8 1.575 2.337 2.538T15 18.5q.9 0 1.775-.25t1.6-.675q.425-.25.875-.213t.775.363q.425.425.388.975t-.488.85q-1.075.7-2.313 1.075T15 21Z"/>`),
		g.Group(children),
	)
}