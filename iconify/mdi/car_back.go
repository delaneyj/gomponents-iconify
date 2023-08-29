package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 11l1-4h10l1 4m.92-5c-.21-.6-.78-1-1.42-1h-11c-.64 0-1.21.4-1.42 1L3 12v8a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-2h12v2a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-8l-2.08-6M7 16H5v-2h2v2m12 0h-2v-2h2v2m-5 0h-4v-2h4v2Z"/>`),
		g.Group(children),
	)
}