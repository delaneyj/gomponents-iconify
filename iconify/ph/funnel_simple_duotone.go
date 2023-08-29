package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 56v144a16 16 0 0 1-16 16H40a16 16 0 0 1-16-16V56a16 16 0 0 1 16-16h176a16 16 0 0 1 16 16Z" opacity=".2"/><path d="M200 128a8 8 0 0 1-8 8H64a8 8 0 0 1 0-16h128a8 8 0 0 1 8 8Zm32-56H24a8 8 0 0 0 0 16h208a8 8 0 0 0 0-16Zm-80 96h-48a8 8 0 0 0 0 16h48a8 8 0 0 0 0-16Z"/></g>`),
		g.Group(children),
	)
}