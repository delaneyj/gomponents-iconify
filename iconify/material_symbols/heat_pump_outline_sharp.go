package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeatPumpOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q-2.5 0-4.25-1.75T6 12q0-2.5 1.75-4.25T12 6q2.5 0 4.25 1.75T18 12q0 2.5-1.75 4.25T12 18Zm-.75-2.075V13.8l-1.5 1.5q.35.225.713.388t.787.237Zm1.5 0q.4-.075.775-.238t.725-.387l-1.5-1.5v2.125Zm2.55-1.675q.225-.35.388-.725t.237-.775H13.8l1.5 1.5Zm-1.5-3h2.125q-.075-.4-.238-.775T15.3 9.75l-1.5 1.5Zm-1.05-1.05l1.5-1.5q-.35-.225-.713-.388t-.787-.237V10.2ZM12 13q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11q-.425 0-.713.288T11 12q0 .425.288.713T12 13Zm-.75-2.8V8.075q-.4.075-.775.238T9.75 8.7l1.5 1.5Zm-3.175 1.05H10.2l-1.5-1.5q-.225.35-.388.725t-.237.775Zm.625 3l1.5-1.5H8.075q.075.4.238.775t.387.725ZM3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}