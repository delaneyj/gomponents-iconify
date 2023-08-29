package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 4h8.5A1.5 1.5 0 0 1 17 5.5v6a1.5 1.5 0 0 1-1.5 1.5h-2.77l4.239 3.587a1.5 1.5 0 0 1 .531 1.145V20h-2v-2.036L9.634 13H9v7H7V4Zm2 7h6V6H9v5Z"/>`),
		g.Group(children),
	)
}