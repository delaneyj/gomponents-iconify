package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 6a2 2 0 0 1 2-2h7v2h-7v5h7v2h-7v7H8V6Z"/>`),
		g.Group(children),
	)
}