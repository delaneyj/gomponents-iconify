package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BicyclingRoundLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="14" cy="4" r="2"/><circle cx="6" cy="18" r="3"/><circle cx="18" cy="18" r="3"/><path stroke-linecap="round" d="M18.5 10h-3.918c-.377 0-.742-.128-1.037-.363L11.386 7.92A2.638 2.638 0 1 0 8.69 12.4l3.09 1.349a1.89 1.89 0 0 1 .982 2.477L12 18"/></g>`),
		g.Group(children),
	)
}