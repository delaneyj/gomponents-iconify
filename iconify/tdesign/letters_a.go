package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20V6a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v14h2v-7h6v7h2Zm-2-9H9V6h6v5Z"/>`),
		g.Group(children),
	)
}