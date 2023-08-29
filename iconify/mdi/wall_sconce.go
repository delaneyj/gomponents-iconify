package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallSconce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 4l-4 9h12l-4-9h-4M4 14v8h2v-3h8v-5h-2v3H6v-3H4Z"/>`),
		g.Group(children),
	)
}