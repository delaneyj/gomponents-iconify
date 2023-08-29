package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRightO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14 13.963h2v-6h-6v2h2.586l-5.33 5.33l1.414 1.414l5.33-5.33v2.586Z"/><path fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm2 0a9 9 0 1 1 18 0a9 9 0 0 1-18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}