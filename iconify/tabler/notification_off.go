package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.154 6.187A2 2 0 0 0 5 8v9a2 2 0 0 0 2 2h9a2 2 0 0 0 1.811-1.151M14 7a3 3 0 1 0 6 0a3 3 0 1 0-6 0M3 3l18 18"/>`),
		g.Group(children),
	)
}