package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditSelectBackBehindBoxDesignLayerLayersSelectSendToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 9.5a1 1 0 0 0 1 1m0-10a1 1 0 0 0-1 1m10 0a1 1 0 0 0-1-1m0 10a1 1 0 0 0 1-1"/><path d="M10.5 13.5h-8a2 2 0 0 1-2-2v-8m7-3h2m-2 10h2m4-6v2m-10-2v2"/></g>`),
		g.Group(children),
	)
}