package heroicons_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9 12.75l3 3m0 0l3-3m-3 3v-7.5M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`),
		g.Group(children),
	)
}