package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.57 3h18.86l-6.93 9.817V21h-5v-8.183L2.57 3Zm3.86 2l5.07 7.183V19h1v-6.817L17.57 5H6.43Z"/>`),
		g.Group(children),
	)
}