package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 3H8v4H2v14h10v-2H4V9h16v2h2V7h-6V3zm-2 4h-4V5h4v2zm6 6h-6v6h6v2h2v-2h-2v-6zm-4 4v-2h2v2h-2z"/>`),
		g.Group(children),
	)
}