package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChopsticksFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14 4v40M8 5v10c0 5 6 5 6 5s6 0 6-5V5m17-1l3 40M31 4l-3 40"/>`),
		g.Group(children),
	)
}