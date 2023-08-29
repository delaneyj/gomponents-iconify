package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActivityZone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="18.595" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.405V2.5m0 43v-2.905M42.595 24H45.5m-43 0h2.905m37.19 0l-9.298-9.297l-16.451 16.451m7.503-14.307L7.897 33.298"/>`),
		g.Group(children),
	)
}