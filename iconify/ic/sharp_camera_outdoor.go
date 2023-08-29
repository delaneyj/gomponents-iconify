package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpCameraOutdoor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 14v-2h-6v6h6v-2l2 1.06v-4.12L18 14zM12 3L4 9v12h16v-2H6v-9l6-4.5l6 4.5v1h2V9l-8-6z"/>`),
		g.Group(children),
	)
}