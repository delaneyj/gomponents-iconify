package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarouselVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 8v8a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1zM7 22v-1a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v1m0-20v1a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V2"/>`),
		g.Group(children),
	)
}