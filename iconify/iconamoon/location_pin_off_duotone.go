package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationPinOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6.5 16.5L12 22l5-5L6.022 6.022A7.779 7.779 0 0 0 6.5 16.5Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.022 6.022A7.779 7.779 0 0 0 6.5 16.5L12 22l5-5M9.344 3.687a7.78 7.78 0 0 1 9.969 9.969"/></g>`),
		g.Group(children),
	)
}