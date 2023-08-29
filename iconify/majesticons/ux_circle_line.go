package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UxCircleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 10v2c0 .667.4 2 2 2s2-1.333 2-2v-2m3 0l1.5 2m1.5 2l-1.5-2m0 0l1.5-2m-1.5 2L14 14"/><circle cx="12" cy="12" r="10"/></g>`),
		g.Group(children),
	)
}