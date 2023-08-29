package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Board(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 2H3a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zM5 10H4V7H3V6h2v4zm4-2H8v1H7V8H6V7h1V6h1v1h1v1zm3 2h-1V7h-1V6h2v4zm3-1h-2V8h2v1zm0-2h-2V6h2v1z"/>`),
		g.Group(children),
	)
}