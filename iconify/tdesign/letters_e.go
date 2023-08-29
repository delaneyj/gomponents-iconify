package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LettersE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6a2 2 0 0 1 2-2h8v2H9v5h8v2H9v5h8v2H9a2 2 0 0 1-2-2V6Z"/>`),
		g.Group(children),
	)
}