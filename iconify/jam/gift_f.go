package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20v-8h6v6a2 2 0 0 1-2 2H9zM7 6v4H1a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h.535A4 4 0 0 1 8 1.354A4 4 0 0 1 14.465 6H15a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H9V6H7zm0 14H3a2 2 0 0 1-2-2v-6h6v8zM7 6V4a2 2 0 1 0-2 2h2zm2 0h2a2 2 0 1 0-2-2v2z"/>`),
		g.Group(children),
	)
}