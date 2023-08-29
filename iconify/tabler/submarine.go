package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Submarine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11v6h2l1-1.5L9 17h10a3 3 0 0 0 0-6H9h0l-3 1.5L5 11H3zm14 0l-1-3h-5l-1 3m3-3V6a1 1 0 0 1 1-1h1"/>`),
		g.Group(children),
	)
}