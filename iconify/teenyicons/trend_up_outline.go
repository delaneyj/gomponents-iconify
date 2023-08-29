package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendUpOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 11.5L5 7l3 3l6.5-6.5m0 0V10m0-6.5H8"/>`),
		g.Group(children),
	)
}