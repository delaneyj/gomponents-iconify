package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.293 10.707l7-7a1 1 0 0 1 1.414 0l7 7a1 1 0 0 1 .293.707V21a1 1 0 0 1-1 1h-5v-7h-4v7H5a1 1 0 0 1-1-1v-9.586a1 1 0 0 1 .293-.707Z"/>`),
		g.Group(children),
	)
}