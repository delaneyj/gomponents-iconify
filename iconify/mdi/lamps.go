package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 2l2 7H2l2-7m2 8h2v10h3v2H3v-2h3V10m14-2l2 7H12l2-7m2 8h2v4h3v2h-8v-2h3v-4Z"/>`),
		g.Group(children),
	)
}