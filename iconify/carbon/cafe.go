package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cafe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M2 28h28v2H2zm22.5-17H8a2.002 2.002 0 0 0-2 2v8a5.006 5.006 0 0 0 5 5h8a5.006 5.006 0 0 0 5-5v-1h.5a4.5 4.5 0 0 0 0-9zM22 21a3.003 3.003 0 0 1-3 3h-8a3.003 3.003 0 0 1-3-3v-8h14zm2.5-3H24v-5h.5a2.5 2.5 0 0 1 0 5zM19 9h-2v-.146a1.988 1.988 0 0 0-1.105-1.789L13.21 5.724A3.979 3.979 0 0 1 11 2.146V1h2v1.146a1.99 1.99 0 0 0 1.105 1.789l2.684 1.341A3.98 3.98 0 0 1 19 8.854z"/>`),
		g.Group(children),
	)
}