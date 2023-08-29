package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.938 13A8.004 8.004 0 0 1 13 19.938V17h-2v2.938A8.004 8.004 0 0 1 4.062 13H7v-2H4.062A8.004 8.004 0 0 1 11 4.062V7h2V4.062A8.004 8.004 0 0 1 19.938 11H17v2h2.938ZM2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2S2 6.477 2 12Zm10 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/>`),
		g.Group(children),
	)
}