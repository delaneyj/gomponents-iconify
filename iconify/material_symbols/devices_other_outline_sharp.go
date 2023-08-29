package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesOtherOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20H2V4h18v2H4v12h3v2Zm4-2.5q.625 0 1.063-.438T12.5 16q0-.625-.438-1.063T11 14.5q-.625 0-1.063.438T9.5 16q0 .625.438 1.063T11 17.5ZM22 20h-7V9h7v11Zm-5-2h3v-7h-3v7Zm-8 2v-1.775q-.475-.425-.738-1T8 16q0-.65.263-1.225t.737-1V12h4v1.775q.475.425.738 1T14 16q0 .65-.263 1.225t-.737 1V20H9Zm8-2h3h-3Z"/>`),
		g.Group(children),
	)
}