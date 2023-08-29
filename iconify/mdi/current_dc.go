package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurrentDc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 9v2h20V9H2m0 4v2h5v-2H2m7 0v2h6v-2H9m8 0v2h5v-2h-5Z"/>`),
		g.Group(children),
	)
}