package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SourceCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12a5 5 0 0 1-4 4.9V21h-2v-4.1a5 5 0 0 1 0-9.8V3h2v4.1a5 5 0 0 1 4 4.9m-5-3a3 3 0 0 0-3 3a3 3 0 0 0 3 3a3 3 0 0 0 3-3a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}