package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Swapcard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.32 20.648l-1.601-1.911h0a8.613 8.613 0 1 0 0 10.527h0l8.816-10.527h0a8.615 8.615 0 0 1 14.754 1.91"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m26.935 27.352l1.6 1.911h0a8.615 8.615 0 0 0 14.735-1.868"/>`),
		g.Group(children),
	)
}