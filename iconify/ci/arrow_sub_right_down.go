package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSubRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 13l5 5m0 0l5-5m-5 5v-7.803c0-1.118 0-1.678-.218-2.105a2 2 0 0 0-.874-.874C14.48 7 13.92 7 12.8 7H3"/>`),
		g.Group(children),
	)
}