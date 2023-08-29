package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoLibrary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 2h12v2H6V2ZM4 5h16v2H4V5ZM1.839 8H22.16l-2.1 14H3.94L1.84 8Zm2.322 2l1.5 10H18.34l1.5-10H4.16Zm6.34 1.5l4.666 3.5l-4.667 3.5v-7Z"/>`),
		g.Group(children),
	)
}