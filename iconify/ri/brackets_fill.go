package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BracketsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3v2H6v14h3v2H4V3h5Zm6 0h5v18h-5v-2h3V5h-3V3Z"/>`),
		g.Group(children),
	)
}