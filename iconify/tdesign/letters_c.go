package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersC(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4H9a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8v-2H9V6h8V4Z"/>`),
		g.Group(children),
	)
}