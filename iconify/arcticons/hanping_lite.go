package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HanpingLite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 12.23h37M24 7v5.23M8.7 41c16.455-4.2 24.74-16.315 28.6-28.77"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.3 41c-16.455-4.2-24.74-16.315-28.6-28.77"/>`),
		g.Group(children),
	)
}