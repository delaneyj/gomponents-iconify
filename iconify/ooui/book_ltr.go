package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 2a7.65 7.65 0 0 0-5 2a7.65 7.65 0 0 0-5-2H1v15h4a7.65 7.65 0 0 1 5 2a7.65 7.65 0 0 1 5-2h4V2zm2.5 13.5H14a4.38 4.38 0 0 0-3 1V5s1-1.5 4-1.5h2.5z"/><path fill="currentColor" d="M9 3.5h2v1H9z"/>`),
		g.Group(children),
	)
}