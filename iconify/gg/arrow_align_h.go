package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowAlignH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 7h-2v10h2V7Zm-7 .757l1.414 1.415L5.586 11H10v2H5.586l1.828 1.828L6 16.243L1.757 12L6 7.757Zm12 8.486l-1.414-1.414L18.414 13H14v-2h4.414l-1.828-1.828L18 7.757L22.243 12L18 16.243Z"/>`),
		g.Group(children),
	)
}