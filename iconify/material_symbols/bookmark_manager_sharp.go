package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkManagerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 22v-2.125l5.15-5.175l2.15 2.1l-5.175 5.2H14Zm8.025-5.9L19.9 13.975l1.4-1.4l2.125 2.125l-1.4 1.4ZM2 20V4h8l2 2h10v2.9L11.075 20H2Z"/>`),
		g.Group(children),
	)
}