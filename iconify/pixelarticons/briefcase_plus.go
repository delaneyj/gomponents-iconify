package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcasePlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 3h8v4h6v4h-2V9H4v10h8v2H2V7h6V3zm2 4h4V5h-4v2zm7 14h2v-3h3v-2h-3v-3h-2v3h-3v2h3v3z"/>`),
		g.Group(children),
	)
}