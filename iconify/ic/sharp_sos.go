package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.5 7h-7v10h7V7zm-2 8h-3V9h3v6zM1 15h4v-2H1V7h6v2H3v2h4v6H1v-2zm16 0h4v-2h-4V7h6v2h-4v2h4v6h-6v-2z"/>`),
		g.Group(children),
	)
}