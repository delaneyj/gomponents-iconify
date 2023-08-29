package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ServiceVictoria(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.197 25.9c.298.388.672.532 1.193.532h.72c.67 0 1.213-.543 1.213-1.213v-.005c0-.67-.543-1.213-1.213-1.213h-.795c-.67 0-1.214-.544-1.214-1.215h0c0-.672.545-1.217 1.217-1.217h.716c.52 0 .894.145 1.192.533M13.178 24h1.585m.423 2.432h-2.431v-4.863h2.431m2.432 4.863V21.57h1.592c.9 0 1.63.731 1.63 1.633s-.73 1.633-1.63 1.633h-1.592m1.592 0l1.592 1.595m5.691-4.862l-1.611 4.863l-1.611-4.863m5.654 0v4.863m5.653-1.61a1.61 1.61 0 1 1-3.222 0V23.18c0-.89.721-1.61 1.611-1.61h0c.89 0 1.611.72 1.611 1.61m2.432.82h1.585m.846 2.432H37.01v-4.863h2.431M28.874 42.5l-4.44-13.662m-3.142-9.666L16.848 5.5m12.026 37l4.43-13.632m3.174-9.767l4.421-13.6M16.848 5.5H40.9"/>`),
		g.Group(children),
	)
}