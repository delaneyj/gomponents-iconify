package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNib(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m10.723.04l-.785 1.843l4.25 4.25l1.83-.773L10.723.04zM3.357 6.132s.94 4.213-3.294 8.447l.302.302l6.764-6.764a1.492 1.492 0 0 1-.145-.633a1.5 1.5 0 1 1 1.5 1.5c-.235 0-.455-.059-.654-.156l-6.758 6.76l.389.389c4.249-4.248 8.463-3.292 8.463-3.292l2.758-6.054l-3.295-3.295l-6.03 2.796z"/>`),
		g.Group(children),
	)
}