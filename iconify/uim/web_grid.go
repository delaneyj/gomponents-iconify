package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13h12v9h2V2h-2v9H2z" opacity=".25"/><path fill="currentColor" d="M21 22h-5V2h5a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1z"/><path fill="currentColor" d="M14 22H3a1 1 0 0 1-1-1v-8h12v9zm0-11H2V3a1 1 0 0 1 1-1h11v9z" opacity=".5"/>`),
		g.Group(children),
	)
}