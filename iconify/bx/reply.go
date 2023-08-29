package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 11h6v7h2v-8a1 1 0 0 0-1-1h-7V6l-5 4l5 4v-3z"/>`),
		g.Group(children),
	)
}