package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20H7v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-9H2V8h1V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v3h1v4h-1v9a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1v-1ZM5 5v7h14V5H5Zm2.5 13a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm9 0a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Z"/>`),
		g.Group(children),
	)
}