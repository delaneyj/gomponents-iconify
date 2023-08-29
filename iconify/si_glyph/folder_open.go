package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.03 4.042l-.802-1H1.015v1H.009V13h1.017l.005.984l13.619-.021l1.303-9.922H8.03v.001zm5.79 8.999H1.711L3.1 4.953h11.932l-1.212 8.088z"/>`),
		g.Group(children),
	)
}