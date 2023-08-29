package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capsule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9a6 6 0 0 1 6-6h0a6 6 0 0 1 6 6v6a6 6 0 0 1-6 6h0a6 6 0 0 1-6-6z"/>`),
		g.Group(children),
	)
}