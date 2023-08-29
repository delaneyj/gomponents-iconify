package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleHangouts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5A16.24 16.24 0 0 0 7.76 20.74c0 8.51 7.74 15.58 16.24 16.18v6.58c8-2.23 16.24-12.47 16.24-22.76S33 4.5 24 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.73 15.64v6.83h3.41v3.42a3.42 3.42 0 0 0 3.42-3.42h0v-6.83Zm9.71 0v6.83h3.42v3.42a3.42 3.42 0 0 0 3.41-3.42h0v-6.83Z"/>`),
		g.Group(children),
	)
}