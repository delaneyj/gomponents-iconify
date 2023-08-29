package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleInfo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.969 0a7.938 7.938 0 1 0 0 15.876A7.938 7.938 0 0 0 8.97 0zM8 3h2v2H8V3zm2 8.765C10 12.447 9.554 13 9 13c-.553 0-1-.552-1-1.235v-4.53C8 6.555 8.447 6 9 6c.554 0 1 .553 1 1.235v4.53z"/>`),
		g.Group(children),
	)
}