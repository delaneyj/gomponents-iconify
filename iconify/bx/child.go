package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Child(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="12" cy="6" r="2" fill="currentColor"/><path fill="currentColor" d="M14 9h-4a1 1 0 0 0-.8.4l-3 4l1.6 1.2L9 13v7h2v-4h2v4h2v-7l1.2 1.6l1.6-1.2l-3-4A1 1 0 0 0 14 9z"/>`),
		g.Group(children),
	)
}