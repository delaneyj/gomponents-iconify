package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.343 15.243a6 6 0 1 0-8.485-8.486a6 6 0 0 0 8.485 8.486Zm1.414-9.9a8.001 8.001 0 0 1 .662 10.565c.016.013.03.027.046.042l4.242 4.242a1 1 0 0 1-1.414 1.415l-4.243-4.243a.99.99 0 0 1-.042-.045A8.001 8.001 0 0 1 5.444 5.343a8 8 0 0 1 11.313 0ZM10.1 7h2v3h3v2h-3v3h-2v-3h-3v-2h3V7Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}