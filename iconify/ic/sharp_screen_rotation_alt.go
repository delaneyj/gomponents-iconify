package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpScreenRotationAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 7.59l6.41-6.41L20.24 11h-2.83L10.4 4L5.41 9H8v2H2V5h2v2.59zM20 19h2v-6h-6v2h2.59l-4.99 5l-7.01-7H3.76l9.83 9.83L20 16.41V19z"/>`),
		g.Group(children),
	)
}