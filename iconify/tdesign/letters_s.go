package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersS(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 6a2 2 0 0 1 2-2H16v2H9.5v5h5a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H8v-2h6.5v-5h-5a2 2 0 0 1-2-2V6Z"/>`),
		g.Group(children),
	)
}