package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerTulip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13a9 9 0 0 0 9 9a9 9 0 0 0-9-9m9 9a9 9 0 0 0 9-9a9 9 0 0 0-9 9m6-19v5a6 6 0 0 1-6 6a6 6 0 0 1-6-6V3c.74 0 1.47.12 2.16.39c.55.23 1.04.57 1.45 1L12 2l2.39 2.39c.41-.43.9-.77 1.45-1A5.9 5.9 0 0 1 18 3Z"/>`),
		g.Group(children),
	)
}