package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sbanken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.418 8.839a8.403 8.403 0 0 0-8.366 8.415h0a8.403 8.403 0 0 0 8.366 8.415h2.855m0 .001h2.854a8.403 8.403 0 0 1 8.366 8.415h0a8.403 8.403 0 0 1-8.39 8.415h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.72 11.681C27.401 9.736 24.895 8.84 19.273 8.84h-2.855M8.825 39.657c2.319 1.945 4.825 2.843 10.448 2.843h2.854"/><circle cx="36.857" cy="8.591" r="3.091" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}