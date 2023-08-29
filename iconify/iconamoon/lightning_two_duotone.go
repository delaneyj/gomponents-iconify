package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M11 14H6L9.5 2H16l-3 8h5l-8 12l1-8Z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M11 14H6L9.5 2H16l-3 8h5l-8 12l1-8Z"/></g>`),
		g.Group(children),
	)
}