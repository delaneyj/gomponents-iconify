package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatShapesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.175 14.975l2.85-7.5q.075-.225.275-.35T11.725 7h.55q.225 0 .425.125t.275.35l2.85 7.575q.125.35-.075.65t-.575.3q-.225 0-.425-.138t-.275-.362l-.625-1.8H10.2l-.625 1.825q-.075.225-.275.35T8.875 16q-.4 0-.613-.325t-.087-.7ZM10.65 12.4h2.7l-1.3-3.75h-.1l-1.3 3.75ZM1 22v-4q0-.425.288-.713T2 17h1V7H2q-.425 0-.713-.288T1 6V2q0-.425.288-.713T2 1h4q.425 0 .713.288T7 2v1h10V2q0-.425.288-.713T18 1h4q.425 0 .713.288T23 2v4q0 .425-.288.713T22 7h-1v10h1q.425 0 .713.288T23 18v4q0 .425-.288.713T22 23h-4q-.425 0-.713-.288T17 22v-1H7v1q0 .425-.288.713T6 23H2q-.425 0-.713-.288T1 22Zm6-3h10v-1q0-.425.288-.713T18 17h1V7h-1q-.425 0-.713-.288T17 6V5H7v1q0 .425-.288.713T6 7H5v10h1q.425 0 .713.288T7 18v1Z"/>`),
		g.Group(children),
	)
}