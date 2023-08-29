package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditLayerAddTwoLayerAddDesignPlusLayersSquareBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="10" height="10" x=".5" y="3.5" rx="1"/><path d="M3.5.5h9a1 1 0 0 1 1 1v9M5.5 6v5M8 8.5H3"/></g>`),
		g.Group(children),
	)
}