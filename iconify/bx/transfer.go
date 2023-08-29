package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transfer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 12l5-4l-5-4v2.999H2v2h13zm7 3H9v-3l-5 4l5 4v-3h13z"/>`),
		g.Group(children),
	)
}