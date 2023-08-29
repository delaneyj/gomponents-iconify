package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Parabola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 25h2c6 0 10-16 18-16s12 16 18 16h2M4 33v6m20-6v6m20-6v6M4 36h40"/>`),
		g.Group(children),
	)
}