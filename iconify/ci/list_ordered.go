package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListOrdered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 17h10M4 15.685V15.5A1.5 1.5 0 0 1 5.5 14h.04c.807 0 1.46.653 1.46 1.46c0 .35-.114.692-.324.973L4 20h3m3-8h10M10 7h10M4 5l2-1v6"/>`),
		g.Group(children),
	)
}