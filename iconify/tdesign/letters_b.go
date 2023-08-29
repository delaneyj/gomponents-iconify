package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersB(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H7V4Zm2 2v5h6V6H9Zm6 7H9v5h6v-5Z"/>`),
		g.Group(children),
	)
}