package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 10v6M2 10l10-5l10 5l-10 5z"/><path d="M6 12v5c3 3 9 3 12 0v-5"/></g>`),
		g.Group(children),
	)
}