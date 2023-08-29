package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PTwoPFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 5a3 3 0 1 1-6 0a3 3 0 0 1 6 0ZM7 3a4 4 0 0 0-4 4v2h2V7a2 2 0 0 1 2-2h3V3H7Zm10 18a4 4 0 0 0 4-4v-2h-2v2a2 2 0 0 1-2 2h-3v2h3ZM7 16a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm10-7a4 4 0 0 0-4 4h8a4 4 0 0 0-4-4ZM3 21a4 4 0 0 1 8 0H3Z"/>`),
		g.Group(children),
	)
}