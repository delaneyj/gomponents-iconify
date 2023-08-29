package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScaleOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M18 14H5V9l13-5h12l13 5v5H30"/><path d="M18 4h12v40H18z"/><path stroke-linecap="round" d="M18 12h4m-4 18h5m-5-12h5m-5 6h4m-4 12h4m-4-26v28"/></g>`),
		g.Group(children),
	)
}