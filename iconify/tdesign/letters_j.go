package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersJ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.667 4H15v14.5a1.5 1.5 0 0 1-1.5 1.5H8v-2h5V6H9.667V4Z"/>`),
		g.Group(children),
	)
}