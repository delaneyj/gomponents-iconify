package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignBackBackDesignLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10.5" height="10.5" x=".5" y=".5" rx="1"/><path d="M13.5 3.5v9a1 1 0 0 1-1 1h-9"/></g>`),
		g.Group(children),
	)
}