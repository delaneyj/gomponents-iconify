package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaneInflight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11.085h5a2 2 0 1 1 0 4H5l-3-6h3l2 2h3l-2-7h3l4 7zM3 21h18"/>`),
		g.Group(children),
	)
}