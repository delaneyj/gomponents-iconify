package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisplayExternalInputSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 22l-1.425-1.425l1.6-1.575H14v-2h4.175L16.6 15.4L18 14l4 4l-4 4ZM3 21v-6h2v4h4v2H3ZM3 9V3h6v2H5v4H3Zm16 0V5h-4V3h6v6h-2Z"/>`),
		g.Group(children),
	)
}