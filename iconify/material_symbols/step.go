package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Step(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15q-1.275 0-2.138-.863T16 12q0-1.275.863-2.138T19 9q1.275 0 2.138.863T22 12q0 1.275-.863 2.138T19 15ZM9 17l-1.4-1.425L10.175 13H2v-2h8.175L7.6 8.4L9 7l5 5l-5 5Z"/>`),
		g.Group(children),
	)
}