package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallMissed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.59 7L12 14.59L6.41 9H11V7H3v8h2v-4.59l7 7l9-9"/>`),
		g.Group(children),
	)
}