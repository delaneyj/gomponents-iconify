package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestTubeEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h10v2h-1v14a4 4 0 0 1-4 4a4 4 0 0 1-4-4V4H7V2m7 2h-4v14a2 2 0 0 0 2 2a2 2 0 0 0 2-2V4Z"/>`),
		g.Group(children),
	)
}