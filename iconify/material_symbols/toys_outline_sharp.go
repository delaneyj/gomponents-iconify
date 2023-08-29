package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToysOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20q-1.125 0-1.963-.713T4.05 17.45H2V10.2h2.8L3 8.4l-1 1L.6 8L4 4.6L5.4 6l-1 1l1.4 1.4L7.3 4h9.425l2.025 6.1H22v7.35h-2.05q-.15 1.125-.988 1.838T17 20q-.95 0-1.713-.55T14.2 18H9.8q-.325.9-1.088 1.45T7 20Zm.4-10H11V6H8.725L7.4 10Zm5.6 0h3.6l-1.325-4H13v4Zm-3.2 6h4.4q.325-.9 1.088-1.45T17 14q.75 0 1.4.35t1.1.95h.5V12H4v3.3h.5q.45-.6 1.1-.95T7 14q.95 0 1.713.55T9.8 16ZM7 18q.425 0 .713-.288T8 17q0-.425-.288-.713T7 16q-.425 0-.713.288T6 17q0 .425.288.713T7 18Zm10 0q.425 0 .713-.288T18 17q0-.425-.288-.713T17 16q-.425 0-.713.288T16 17q0 .425.288.713T17 18Zm-5-4Z"/>`),
		g.Group(children),
	)
}