package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumbersOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 18h-2V6a2 2 0 0 0-2-2H9v2h2v12H9v2h6v-2Z"/>`),
		g.Group(children),
	)
}