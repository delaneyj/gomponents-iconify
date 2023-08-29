package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRightO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M14.037 10.037h2v6h-6v-2h2.585l-5.329-5.33l1.414-1.414l5.33 5.33v-2.586Z"/><path fill-rule="evenodd" d="M23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11Zm-2 0a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}