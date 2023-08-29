package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 17.5h-8.586l3 3l-1.414 1.414L9.086 16.5l5.414-5.414l1.414 1.414l-3 3H21.5v2Zm-6.586-10L9.5 12.914L8.086 11.5l3-3H2.5v-2h8.586l-3-3L9.5 2.086L14.914 7.5Z"/>`),
		g.Group(children),
	)
}