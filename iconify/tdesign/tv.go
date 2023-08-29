package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2h22v17h-6v-2h4V4H3l.001 13h4v2H1V2Zm11 13.587l4.242 4.242l2.42 2.415h-2.833l-1-1L12 18.417l-2.827 2.828l-1 1H5.34l2.419-2.414L12 15.587Z"/>`),
		g.Group(children),
	)
}