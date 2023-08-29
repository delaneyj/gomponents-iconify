package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BeakerPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 14h2v3h3v2h-3v3h-2v-3h-3v-2h3v-3M3 3h18v2c-1.1 0-2 .9-2 2v5c-.7 0-1.37.12-2 .34V5H7v2h5v1H7v1h3v1H7v1h3v1H7v1h5v1H7v1h3v1H7v3h6.08c.12.72.37 1.39.72 2H7a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2V3Z"/>`),
		g.Group(children),
	)
}