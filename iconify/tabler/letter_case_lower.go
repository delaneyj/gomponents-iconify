package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterCaseLower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 1 0-7 0m7-3.5v7m4-3.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 1 0-7 0m7-3.5v7"/>`),
		g.Group(children),
	)
}