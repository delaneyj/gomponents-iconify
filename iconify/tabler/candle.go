package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 21h6v-9a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v9zm3-18l1.465 1.638a2 2 0 1 1-3.015.099L12 3z"/>`),
		g.Group(children),
	)
}