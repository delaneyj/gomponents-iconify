package memory

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassSouthWest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 22 22"),
		g.Raw(`<path fill="currentColor" d="M3 6h5v2H4v2h3v1h1v4H7v1H2v-2h4v-2H3v-1H2V7h1m7-1h2v6h1v1h1V8h2v5h1v-1h1V6h2v7h-1v2h-1v1h-2v-1h-2v1h-2v-1h-1v-2h-1"/>`),
		g.Group(children),
	)
}