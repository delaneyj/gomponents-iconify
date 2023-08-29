package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccinesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 22.5L6 21v-4H3V7.5H2v-2h4V4H4.5V2h5v2H8v1.5h4v2h-1V17H8v5.5ZM5 15h4v-1.5H6.5V12H9v-1.5H6.5V9H9V7.5H5V15Zm8 7v-8.5q0-.725.25-1.2t.525-.825q.275-.35.5-.562t.225-.413V10h-1V8h7v2h-1v.5q0 .2.25.45t.55.6q.275.35.488.825T21 13.5V22h-8Zm2-8h4v-.5q0-.375-.225-.65t-.5-.6q-.275-.325-.525-.725T17.5 10.5V10h-1v.5q0 .6-.238 1t-.512.725q-.275.325-.513.613T15 13.5v.5Zm0 3h4v-1.5h-4V17Zm0 3h4v-1.5h-4V20Zm0-3h4h-4Z"/>`),
		g.Group(children),
	)
}