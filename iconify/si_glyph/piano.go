package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Piano(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 3v9h15.986V3H0zm2 1h2v1H2V4zm13.083 7.059H.992V7h1.091v2.938h.909V7h1.091v2.938h.909V7h1.091v2.938h.909V7h2.091v2.938h.909V7h1.091v2.938h.909V7h1.091v2.938h.909V7h1.091v4.059zM15 5H8V4h7v1z"/>`),
		g.Group(children),
	)
}