package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpiritLevel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 8H2v8h20V8m-4 6v-4h2v4h-2m-7-2h2a2 2 0 0 0 2-2h2v4H7v-4h2a2 2 0 0 0 2 2m-7 2v-4h2v4H4Z"/>`),
		g.Group(children),
	)
}