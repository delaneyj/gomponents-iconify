package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustVerticalAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M3.5 0v8.5m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0V15m8-15v2.5m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0V15"/>`),
		g.Group(children),
	)
}