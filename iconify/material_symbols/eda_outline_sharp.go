package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdaOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 12V2h2v10H7Zm4 0V1h2v11h-2Zm4 4.075V3h2v11.9l-2 1.175ZM5 21h9.075l4.85-4.85l-7.175 4.175L8.5 16H5v5Zm-2 2v-9h6.5l2.75 3.675L17 14.9l3.4-1.95l2.625 1.95l-8.1 8.1H3Zm2-9H3V4h2v10Zm0-2h10H5Zm0 9h9.075H5Z"/>`),
		g.Group(children),
	)
}