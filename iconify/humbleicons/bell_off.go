package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16M5 17h2m0 0h10M7 17v-6c0-1.126.372-2.164 1-3m3.5-2.5V4a.5.5 0 0 1 1 0v1.5M17 13v-2a5 5 0 0 0-6.5-4.771M13 20a1 1 0 1 1-2 0h2Z"/>`),
		g.Group(children),
	)
}