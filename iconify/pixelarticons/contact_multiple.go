package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h20v14H4V3zm18 12V5H6v10h16zm-2 4H2V7H0v14h20v-2zM9 7h2v2H9V7zm3 4H8v2h4v-2zm2-4h6v2h-6V7zm6 4h-6v2h6v-2z"/>`),
		g.Group(children),
	)
}