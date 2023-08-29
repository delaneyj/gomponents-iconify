package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropFreeSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h2v4h4v2H3Zm12 0v-2h4v-4h2v6h-6ZM3 9V3h6v2H5v4H3Zm16 0V5h-4V3h6v6h-2Z"/>`),
		g.Group(children),
	)
}