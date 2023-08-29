package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraveTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 16.17V7a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v9.171M12 7v5m-2-3h4"/><path d="M5 21v-2a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v2H5z"/></g>`),
		g.Group(children),
	)
}