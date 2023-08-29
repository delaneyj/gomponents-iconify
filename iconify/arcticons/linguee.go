package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linguee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.607 16.201a24.498 24.498 0 0 0 3.662.546c2.465.012 5.082.008 7.325-1.013a6.664 6.664 0 0 0 3.04-2.65a4.023 4.023 0 0 0 .545-3.117a2.215 2.215 0 0 0-1.48-1.325a2.794 2.794 0 0 0-2.338.546c-1.745 1.429-2.155 3.997-2.806 6.156c-1.381 4.583-1.161 9.523-2.26 14.183a17.967 17.967 0 0 1-1.324 4.208a8.821 8.821 0 0 1-4.364 4.597a6.158 6.158 0 0 1-4.286.857a2.557 2.557 0 0 1-1.87-1.558a2.673 2.673 0 0 1 1.013-2.65c1.733-1.193 4.156-.796 6.312 0a79.315 79.315 0 0 0 11.065 3.585a37.28 37.28 0 0 0 8.572.857a77.466 77.466 0 0 0 10.209-1.558"/>`),
		g.Group(children),
	)
}