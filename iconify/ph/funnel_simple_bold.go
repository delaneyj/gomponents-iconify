package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunnelSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M204 128a12 12 0 0 1-12 12H64a12 12 0 0 1 0-24h128a12 12 0 0 1 12 12Zm28-60H24a12 12 0 0 0 0 24h208a12 12 0 0 0 0-24Zm-80 96h-48a12 12 0 0 0 0 24h48a12 12 0 0 0 0-24Z"/>`),
		g.Group(children),
	)
}