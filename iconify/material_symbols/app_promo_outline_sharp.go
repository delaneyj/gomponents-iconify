package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppPromoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 11l-4 4l-4-4l1.4-1.4l1.6 1.55V7h2v4.15l1.6-1.55L16 11Zm-6 9h4v-1h-4v1Zm-5 3V1h14v22H5Zm2-5v3h10v-3H7Zm0-2h10V6H7v10ZM7 4h10V3H7v1Zm0 14v3v-3ZM7 4V3v1Z"/>`),
		g.Group(children),
	)
}