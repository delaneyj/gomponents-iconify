package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 11H10V2H8v20h2v-9h12z" opacity=".25"/><path fill="currentColor" d="M3 2h5v20H3a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M10 2h11a1 1 0 0 1 1 1v8H10V2zm0 11h12v8a1 1 0 0 1-1 1H10v-9z" opacity=".5"/>`),
		g.Group(children),
	)
}