package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditSprayCanColorColorsDesignPaintPaintingSpray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="6" height="9" x=".5" y="4.5" rx="1"/><path d="M3.5 2v2.5m5-3l5-1m-5 3l5 1M2.5 2h2"/></g>`),
		g.Group(children),
	)
}