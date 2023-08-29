package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TravelMapNavigationMapMapsGps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m4.83 12.5l-4.33 1v-12l4.33-1v12zm0 0l4.34 1v-12L4.83.5v12zm8.67 0l-4.33 1v-12l4.33-1v12z"/>`),
		g.Group(children),
	)
}