package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubdirectoryLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 13h-6v-3l-5 4l5 4v-3h7a1 1 0 0 0 1-1V5h-2v8z"/>`),
		g.Group(children),
	)
}