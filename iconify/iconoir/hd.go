package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 7v5m0 5v-5m0 0h7m0 0V7m0 5v5m3-5V7c4 0 8 0 8 5s-4 5-8 5v-5Z"/>`),
		g.Group(children),
	)
}