package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityAncientTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.998.802l6.054 3.993L17 6.39V9.5l3.334 2.5h1.667v2h-1v8h-18v-8h-1v-2h1.667L7 9.5V6.39L5.948 4.794l6.05-3.993ZM9 5.177V9h6V5.177l-3.001-1.98l-3 1.98ZM15.666 11H8.333L7 12h10l-1.334-1Zm3.335 3h-14v6h5.998v-3h2v3H19v-6Zm-8-9.002h2.004v2.004H11V4.998Z"/>`),
		g.Group(children),
	)
}