package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VipCrownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.005 19h20v2h-20v-2Zm0-14l5 3.5l5-6.5l5 6.5l5-3.5v12h-20V5Zm2 3.841V15h16V8.841l-3.42 2.394l-4.58-5.955l-4.58 5.955l-3.42-2.394Z"/>`),
		g.Group(children),
	)
}