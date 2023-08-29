package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunsetTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 13h1m16 0h1M5.6 6.6l.7.7m12.1-.7l-.7.7M8 13a4 4 0 1 1 8 0M3 17h18M7 20h5m4 0h1M12 5V4"/>`),
		g.Group(children),
	)
}