package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfBall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.031 12l2 1.969V16h.907v-2.031l2-1.969H6.031zm-3-6.547a5.452 5.452 0 1 0 10.906 0a5.452 5.452 0 1 0-10.906 0z"/>`),
		g.Group(children),
	)
}