package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 1.5h2v13H5v-13Zm4 3h2v18H9v-18Zm6 0h-2v18h2v-18Zm2 0h2v10h-2v-10Z"/>`),
		g.Group(children),
	)
}