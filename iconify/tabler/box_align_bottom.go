package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxAlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14h16v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-5zm0-5v.01M4 4v.01M9 4v.01M15 4v.01M20 4v.01M20 9v.01"/>`),
		g.Group(children),
	)
}