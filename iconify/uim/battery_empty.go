package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 14a1 1 0 0 1-1-1v-2a1 1 0 0 1 2 0v2a1 1 0 0 1-1 1zm-4 3H4a2.002 2.002 0 0 1-2-2V9a2.002 2.002 0 0 1 2-2h13a2.002 2.002 0 0 1 2 2v6a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}