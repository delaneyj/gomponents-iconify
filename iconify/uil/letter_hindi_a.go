package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterHindiA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.022 3h-5a1 1 0 0 0 0 2h1.5v6h-4.95a4.951 4.951 0 0 0 1.026-3a5 5 0 0 0-9.33-2.5a1 1 0 1 0 1.731 1A3 3 0 1 1 7.598 11a1 1 0 0 0 0 2a3 3 0 1 1-2.599 4.5a1 1 0 0 0-1.731 1a5 5 0 0 0 9.33-2.5a4.951 4.951 0 0 0-1.026-3h4.95v7a1 1 0 0 0 2 0V5h1.5a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}