package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Contrast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.502.115a7.42 7.42 0 1 0 .003 14.839A7.42 7.42 0 0 0 8.502.115zm-.584 13.127V1.827c3.477 0 6.292 2.556 6.292 5.707s-2.815 5.708-6.292 5.708z"/>`),
		g.Group(children),
	)
}