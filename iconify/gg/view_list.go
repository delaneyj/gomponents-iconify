package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3H5Zm2 2H5a1 1 0 0 0-1 1v1h3V7Zm2 0v2h11V8a1 1 0 0 0-1-1H9Zm-2 4H4v2h3v-2Zm2 2v-2h11v2H9Zm-2 2H4v1a1 1 0 0 0 1 1h2v-2Zm2 2v-2h11v1a1 1 0 0 1-1 1H9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}