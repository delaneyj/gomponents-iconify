package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vibb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.868 8.558L24 25.28L17.132 8.558H5.5l12.684 30.884h11.632L42.5 8.558H30.868z"/>`),
		g.Group(children),
	)
}