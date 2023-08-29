package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.289.023L6.925 0L2.984 8H8l-4.334 7.916L14.924 4.941H10.35L14.289.023z"/>`),
		g.Group(children),
	)
}