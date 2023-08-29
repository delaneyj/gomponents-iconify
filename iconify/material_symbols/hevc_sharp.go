package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HevcSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15V9h1.5v2h1V9H7v6H5.5v-2.5h-1V15H3Zm14 0V9h4v2h-1.5v-.5h-1v3h1V13H21v2h-4Zm-4 0l-1-6h1.5l.75 4.5L15 9h1.5l-1 6H13Zm-5 0V9h3.5v1.5h-2v.5h2v1.5h-2v1h2V15H8Z"/>`),
		g.Group(children),
	)
}