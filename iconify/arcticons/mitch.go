package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16 23.6l8 8l8-8l9.95 19.9H32l-4-4l-4 4l-4-4l-4 4H6.09Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 26.78l-9.28-9.28h5.57v-13h7.42v13h5.57Z"/>`),
		g.Group(children),
	)
}