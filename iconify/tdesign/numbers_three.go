package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H7v-2h8v-5H8v-2h7V6H7V4Z"/>`),
		g.Group(children),
	)
}