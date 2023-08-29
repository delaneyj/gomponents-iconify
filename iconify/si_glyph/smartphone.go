package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smartphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.001 0H5C3.897 0 3.003 0 3.003 1v14c0 1.119.895.947 1.997.947h5.001c1.103 0 1.999.172 1.999-.947V1c0-1-.896-1-1.999-1zM8.125 15.188h-1.23v-1.375h1.23v1.375zM11.037 13H4V1h7.037v12z"/>`),
		g.Group(children),
	)
}