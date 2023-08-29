package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignLayersTwoDesignLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="8" height="6" x="5.5" y="1.5" rx="1"/><path d="M11 10H4a1 1 0 0 1-1-1V4"/><path d="M9 12.5H1.5a1 1 0 0 1-1-1V6"/></g>`),
		g.Group(children),
	)
}