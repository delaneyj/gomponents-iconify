package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Group(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 1v4h1v14H1v4h4v-1h14v1h4v-4h-1V5h1V1h-4v1H5V1m0 3h14v1h1v14h-1v1H5v-1H4V5h1m1 1v8h3v4h9V9h-4V6M8 8h4v4H8m6-1h2v5h-5v-2h3"/>`),
		g.Group(children),
	)
}