package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaBluetoothOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.85 20.75q-.2.2-.525.063T17 20.4v-3.95l-2.325 2.325q-.175.175-.425.175t-.425-.175q-.175-.175-.175-.425t.175-.425L16.75 15l-2.975-2.975q-.175-.175-.175-.425t.175-.425Q13.95 11 14.2 11t.425.175L16.95 13.5V9.6q0-.275.325-.412t.525.062l2.475 2.475q.15.15.225.338t.075.387q0 .2-.075.375t-.225.325L18.4 15l1.875 1.85q.15.15.225.325t.075.375q0 .2-.063.388t-.212.337l-2.45 2.475Zm.3-7.2l1.15-1.1l-1.15-1.15v2.25Zm0 5.15l1.15-1.15l-1.15-1.1v2.25ZM7 21q-1.65 0-2.825-1.175T3 17q0-1.65 1.175-2.825T7 13q.575 0 1.063.138T9 13.55V5q0-.825.588-1.413T11 3h2q.825 0 1.413.588T15 5q0 .825-.588 1.413T13 7h-2v10q0 1.65-1.175 2.825T7 21Z"/>`),
		g.Group(children),
	)
}