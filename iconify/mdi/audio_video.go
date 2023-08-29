package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AudioVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 7H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h1v1c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-1h6v1c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-1h1a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2m-6 5H4v-2h10v2m4 1a2 2 0 0 1-2-2a2 2 0 0 1 2-2a2 2 0 0 1 2 2a2 2 0 0 1-2 2M6 15H4v-1h2v1m4 0H8v-1h2v1m4 0h-2v-1h2v1Z"/>`),
		g.Group(children),
	)
}