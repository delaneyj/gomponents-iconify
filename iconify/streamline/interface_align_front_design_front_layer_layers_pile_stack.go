package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceAlignFrontDesignFrontLayerLayersPileStack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10.5" height="10.5" x="3" y="3" rx="1" transform="rotate(180 8.25 8.25)"/><path d="M.5 10.5v-9a1 1 0 0 1 1-1h9"/></g>`),
		g.Group(children),
	)
}