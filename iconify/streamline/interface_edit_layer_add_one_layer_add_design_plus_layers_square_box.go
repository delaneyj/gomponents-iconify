package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditLayerAddOneLayerAddDesignPlusLayersSquareBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.5 3.19v4.62M3.19 5.5h4.62"/><rect width="10" height="10" x=".5" y=".5" rx="1"/><path d="M3.5 13.5h9a1 1 0 0 0 1-1v-9"/></g>`),
		g.Group(children),
	)
}