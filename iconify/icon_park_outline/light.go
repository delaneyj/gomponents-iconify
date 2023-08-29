package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Light(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 16v6m14.142-.142L33.9 26.1M44 36h-6M4 36h6m-.142-14.142L14.1 26.1M18 36h12"/>`),
		g.Group(children),
	)
}