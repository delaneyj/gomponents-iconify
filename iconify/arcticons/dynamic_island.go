package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DynamicIsland(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="41" height="17.872" x="3.5" y="15.629" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="8.936" ry="8.936"/>`),
		g.Group(children),
	)
}