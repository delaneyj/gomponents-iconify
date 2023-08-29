package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 4H8.5A1.5 1.5 0 0 0 7 5.5v13A1.5 1.5 0 0 0 8.5 20H17v-9h-5.2v2H15v5H9V6h7V4Z"/>`),
		g.Group(children),
	)
}