package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WallSconceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.7 6l2.22 5h-5.84l2.22-5h1.4M15 4h-4l-4 9h12l-4-9M4 14v8h2v-3h8v-5h-2v3H6v-3H4Z"/>`),
		g.Group(children),
	)
}