package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustHorizontalAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15 3.5H6.5m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H0m15 8h-2.5m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H0"/>`),
		g.Group(children),
	)
}