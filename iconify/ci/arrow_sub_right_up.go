package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSubRightUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 11l5-5m0 0l5 5m-5-5v7.803c0 1.118 0 1.677-.218 2.105a2.002 2.002 0 0 1-.874.874C14.48 17 13.92 17 12.803 17H3"/>`),
		g.Group(children),
	)
}