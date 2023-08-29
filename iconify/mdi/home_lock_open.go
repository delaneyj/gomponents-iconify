package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeLockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L2 12h3v8h14v-8h3L12 3m0 5a3 3 0 0 1 3 3h-2a1 1 0 0 0-1-1a1 1 0 0 0-1 1v2h5v4H8v-4h1v-2a3 3 0 0 1 3-3Z"/>`),
		g.Group(children),
	)
}