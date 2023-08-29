package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubdirectoryRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 13H8V5H6v9a1 1 0 0 0 1 1h7v3l5-4l-5-4v3z"/>`),
		g.Group(children),
	)
}