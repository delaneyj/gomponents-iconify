package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeometryDashSubzero(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 3.5L44.5 24L24 44.498l-20.5-20.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.222 23.528a3.561 3.561 0 1 0 6.52-1.448l-6.52 1.448Zm11.306-11.306a3.561 3.561 0 1 1-1.448 6.52l1.448-6.52Zm9.188 16.976v-4.16h2.676v-2.437l-.77-3.971h-1.796v2.23h-3.775v4.068h-4.123v4.123H20.86v3.775h-2.23v1.796l3.971.77h2.437v-2.676h4.16V29.18l3.518.018z"/>`),
		g.Group(children),
	)
}