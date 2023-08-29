package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3h8v4h6v6h-2V9H4v10h10v2H2V7h6V3zm6 2h-4v2h4V5zm2 12h6v2h-6v-2z"/>`),
		g.Group(children),
	)
}