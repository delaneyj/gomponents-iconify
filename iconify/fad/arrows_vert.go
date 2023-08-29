package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsVert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M175.863 100.122c0-2.205 1.293-2.747 2.883-1.214l30.096 28.996l-30.11 29.24c-1.585 1.538-2.87 1-2.87-1.209v-19.24l-95.811.637v18.596c0 2.21-1.28 2.746-2.854 1.201l-29.788-29.225l29.774-28.982c1.584-1.542 2.868-1.004 2.868 1.2v19.54h95.812v-19.54z"/>`),
		g.Group(children),
	)
}