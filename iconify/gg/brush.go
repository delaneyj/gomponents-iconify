package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 11h3a1 1 0 0 1 1 1v6a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3v-6a1 1 0 0 1 1-1h3V6a3 3 0 1 1 6 0v5Zm-2-5a1 1 0 1 0-2 0v7H7v5a1 1 0 0 0 1 1h1v-3h2v3h2v-3h2v3h1a1 1 0 0 0 1-1v-5h-4V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}