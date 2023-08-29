package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6v5a1 1 0 0 0 1 1h3v8H9a1 1 0 0 0 0 2h6a1 1 0 0 0 0-2h-2v-8h5a1 1 0 0 0 .78-.37l2-2.5a1 1 0 0 0 0-1.25l-2-2.5A1 1 0 0 0 18 5h-5V3a1 1 0 0 0-2 0v2H8a1 1 0 0 0-1 1Zm2 1h8.52l1.2 1.5l-1.2 1.5H9Z"/>`),
		g.Group(children),
	)
}