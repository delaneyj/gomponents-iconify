package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoreOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4h14q.425 0 .713.288T20 5q0 .425-.288.713T19 6H5q-.425 0-.713-.288T4 5q0-.425.288-.713T5 4Zm0 16q-.425 0-.713-.288T4 19v-5h-.175q-.475 0-.775-.363t-.2-.837l1-5q.075-.35.35-.575T4.825 7h14.35q.35 0 .625.225t.35.575l1 5q.1.475-.2.837t-.775.363H20v5q0 .425-.288.713T19 20q-.425 0-.713-.288T18 19v-5h-4v5q0 .425-.288.713T13 20H5Zm1-2h6v-4H6v4Zm-.95-6h13.9h-13.9Zm0 0h13.9l-.6-3H5.65l-.6 3Z"/>`),
		g.Group(children),
	)
}