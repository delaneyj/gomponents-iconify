package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PuzzleRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="3" cy="10" r="3" fill="currentColor"/><path fill="currentColor" d="M9.42 3A2.94 2.94 0 0 0 9 4.5a3 3 0 0 0 6 0a2.94 2.94 0 0 0-.42-1.5H19v12a2 2 0 0 1-2 2H5V3z"/>`),
		g.Group(children),
	)
}