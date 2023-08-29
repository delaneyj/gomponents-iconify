package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 21a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2v-1a7 7 0 1 0-14 0v1h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H3v-9a9 9 0 0 1 18 0v9h-4Zm2-6h-2v4h2v-4ZM7 15H5v4h2v-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}