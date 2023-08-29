package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuickReferenceAllOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4v6.025V10v10V4v5v-5Zm2 10h3.5q.225-.575.55-1.075T11.8 12H7v2Zm0 4h3.175q-.125-.5-.163-1t.013-1H7v2Zm-4 4V2h10l6 6v2.5q-.475-.2-.975-.313T17 10.025V9h-5V4H5v16h6.025q.4.6.9 1.113t1.1.887H3Zm13.5-3q1.05 0 1.775-.725T19 16.5q0-1.05-.725-1.775T16.5 14q-1.05 0-1.775.725T14 16.5q0 1.05.725 1.775T16.5 19Zm5.1 4l-2.7-2.7q-.525.35-1.137.525T16.5 21q-1.875 0-3.187-1.313T12 16.5q0-1.875 1.313-3.188T16.5 12q1.875 0 3.188 1.313T21 16.5q0 .65-.175 1.263T20.3 18.9l2.7 2.7l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}