package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageFailed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M25.5 37H21l-10 5v-5H4V7h40v11m-32-3h6m-6 6h12m8 4l12 12m0-12L32 37"/>`),
		g.Group(children),
	)
}