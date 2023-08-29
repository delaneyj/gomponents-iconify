package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignVerticalDistributeStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M2 14h20M2 4h20"/></g>`),
		g.Group(children),
	)
}