package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSubDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 11l-5 5m0 0l5 5m-5-5h7.803c1.118 0 1.677 0 2.105-.218a2 2 0 0 0 .874-.874c.218-.428.218-.987.218-2.105V3"/>`),
		g.Group(children),
	)
}