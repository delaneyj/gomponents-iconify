package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3L2 12h3v8h14v-8h3L12 3m0 6a3 3 0 0 1 3 3v1h1v4H8v-4h1v-1a3 3 0 0 1 3-3m0 2a1 1 0 0 0-1 1v1h2v-1c0-.5-.4-1-1-1Z"/>`),
		g.Group(children),
	)
}