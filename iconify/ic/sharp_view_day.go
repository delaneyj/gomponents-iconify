package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpViewDay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21h19v-3H2v3zM21 8H2v8h19V8zM2 3v3h19V3H2z"/>`),
		g.Group(children),
	)
}