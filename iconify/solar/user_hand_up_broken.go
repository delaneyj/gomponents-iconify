package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserHandUpBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M8 13h8c1.71 0 3.15 1.28 3.35 2.98L20 21.5M8 13a7.459 7.459 0 0 1-5.618-5.472L2 6m6 7v5c0 1.886 0 2.828.586 3.414C9.172 22 10.114 22 12 22c1.886 0 2.828 0 3.414-.586C16 20.828 16 19.886 16 18v-1"/><circle cx="12" cy="6" r="4"/></g>`),
		g.Group(children),
	)
}