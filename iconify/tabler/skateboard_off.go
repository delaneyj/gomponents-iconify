package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkateboardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15a2 2 0 1 0 4 0a2 2 0 1 0-4 0m10 0a2 2 0 0 0 2 2m2-2a2 2 0 0 0-2-2M3 9c0 .552.895 1 2 1h5m4 0h5c1.105 0 2-.448 2-1M3 3l18 18"/>`),
		g.Group(children),
	)
}