package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path fill-opacity=".5" d="m13.035 7.988l.002 6l4-.002l-.002-6l-4 .002Z"/><path d="M18 4.012L6 4.018v2l12-.006v-2Zm-6.963 15.974l-.005-12l-4 .002l.005 12l4-.002Z"/></g>`),
		g.Group(children),
	)
}