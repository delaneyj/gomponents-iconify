package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 3l18 18m2-11v4M5.5 6H3a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h14.5m2.5-3.5V8a2 2 0 0 0-2-2h-6.5"/>`),
		g.Group(children),
	)
}