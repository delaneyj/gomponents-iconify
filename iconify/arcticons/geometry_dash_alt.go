package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeometryDashAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 24L24 3.5L44.498 24l-20.5 20.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m11.635 24.823l3.879-3.878l3.878 3.878l-3.878 3.878zm8.965-8.965l3.879-3.879l3.878 3.879l-3.878 3.878zm-1.252 17.687L33.2 19.693l3.324 3.324L22.673 36.87z"/>`),
		g.Group(children),
	)
}