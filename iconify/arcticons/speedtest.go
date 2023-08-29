package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speedtest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.29 2a22 22 0 0 1 15.27 37.57l-4.11-4.11a16.2 16.2 0 1 0-22.86-.05l-4.15 4.16A22 22 0 0 1 24.29 2Zm8.52 13.21L26.94 27h-5.88v-5.91Z"/>`),
		g.Group(children),
	)
}