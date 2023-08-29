package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmPanel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3H4c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2M8 19H5v-2h3v2m0-3H5v-2h3v2m0-3H5v-2h3v2m5.5 6h-3v-2h3v2m0-3h-3v-2h3v2m0-3h-3v-2h3v2m5.5 6h-3v-2h3v2m0-3h-3v-2h3v2m0-3h-3v-2h3v2m0-4H5V5h14v4Z"/>`),
		g.Group(children),
	)
}