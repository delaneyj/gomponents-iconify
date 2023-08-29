package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<circle cx="9" cy="9" r="8" fill="currentColor" fill-rule="evenodd"/>`),
		g.Group(children),
	)
}