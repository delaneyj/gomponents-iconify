package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstAidKitFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 1a1 1 0 0 1 1 1v3h4a1 1 0 0 1 1 1v14a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h4V2a1 1 0 0 1 1-1h8Zm-3 8h-2v3H8v2h2.999L11 17h2l-.001-3H16v-2h-3V9Zm2-6H9v2h6V3Z"/>`),
		g.Group(children),
	)
}