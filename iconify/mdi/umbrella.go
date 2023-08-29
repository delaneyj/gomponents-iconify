package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a9 9 0 0 1 9 9h-8v8a3 3 0 0 1-3 3a3 3 0 0 1-3-3v-1h2v1a1 1 0 0 0 1 1a1 1 0 0 0 1-1v-8H3a9 9 0 0 1 9-9Z"/>`),
		g.Group(children),
	)
}