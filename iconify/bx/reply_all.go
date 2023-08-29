package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 18v-8a1 1 0 0 0-1-1h-6V6l-5 4l5 4v-3h5v7h2z"/><path fill="currentColor" d="M9 12.4L6 10l3-2.4V6l-5 4l5 4z"/>`),
		g.Group(children),
	)
}