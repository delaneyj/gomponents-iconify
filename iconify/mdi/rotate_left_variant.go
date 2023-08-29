package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RotateLeftVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h3a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2m16 13a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2h-9v-7h9M14 4a8 8 0 0 1 8 8l-.06 1h-2.02l.08-1a6 6 0 0 0-6-6v3l-4-4l4-4v3Z"/>`),
		g.Group(children),
	)
}