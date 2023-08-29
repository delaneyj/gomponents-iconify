package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15h2v4h14v-4h2v4c0 1.1-.9 2-2 2H5c-1.1 0-2-.9-2-2v-4Z"/>`),
		g.Group(children),
	)
}