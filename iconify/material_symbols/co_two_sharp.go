package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 15V9h5v6h-5Zm1.5-1.5h2v-3h-2v3ZM3 15V9h5v2H6.5v-.5h-2v3h2V13H8v2H3Zm14 3v-3.5h3v-1h-3V12h4.5v3.5h-3v1h3V18H17Z"/>`),
		g.Group(children),
	)
}